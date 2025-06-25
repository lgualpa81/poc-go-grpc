#!/bin/bash

# Variables de conexión
PG_USER="${POSTGRES_USER:-postgres}"  
PG_PASSWORD="${POSTGRES_PASSWORD}"
PG_HOST="localhost"
PG_PORT="${POSTGRES_PORT:-5432}"
MAX_ATTEMPTS=30  # Número máximo de intentos de conexión
ATTEMPT=0
WAIT_SECONDS=2

# Lista de bases de datos a crear
DATABASES=("auth_svc" "order_svc" "product_svc")

# Función mejorada para esperar PostgreSQL con timeout
wait_for_postgres() {
  echo "Esperando a que PostgreSQL esté listo (máximo $((MAX_ATTEMPTS * WAIT_SECONDS)) segundos)..."
  
  while [ $ATTEMPT -lt $MAX_ATTEMPTS ]; do
    if PGPASSWORD=$PG_PASSWORD psql -p "$PG_PORT" -U "$PG_USER" -d "postgres" -c '\q' >/dev/null 2>&1; then
    #if pg_isready  -p "$PG_PORT" -U "$PG_USER" >/dev/null 2>&1; then
      echo "¡PostgreSQL está listo!"
      return 0
    fi
    
    ATTEMPT=$((ATTEMPT + 1))
    echo "Intento $ATTEMPT/$MAX_ATTEMPTS: PostgreSQL no está disponible aún, esperando $WAIT_SECONDS segundos..."
    sleep $WAIT_SECONDS
  done
  
  echo "ERROR: No se pudo conectar a PostgreSQL después de $MAX_ATTEMPTS intentos"
  echo "Verifica:"
  echo "1. Que el contenedor esté corriendo (docker ps)"
  echo "2. Que las credenciales sean correctas"
  echo "3. Que el puerto $PG_PORT esté accesible"
  return 1
}

# Función para verificar si una base de datos existe
database_exists() {
  local dbname="$1"
  PGPASSWORD=$PG_PASSWORD psql -p "$PG_PORT" -U "$PG_USER" -d "postgres" \
    -tc "SELECT 1 FROM pg_database WHERE datname = '$dbname'" | grep -q 1
}

# Crear las bases de datos si no existen
create_databases_if_not_exist() {
  for db in "${DATABASES[@]}"; do
    if database_exists "$db"; then
      echo "La base de datos '$db' ya existe, omitiendo..."
    else
      echo "Creando base de datos: '$db'"
      if PGPASSWORD=$PG_PASSWORD psql  -p "$PG_PORT" -U "$PG_USER" -d "postgres" \
        -c "CREATE DATABASE $db;"; then
        echo "Base de datos '$db' creada exitosamente"
      else
        echo "ERROR al crear la base de datos '$db'"
        return 1
      fi
    fi
  done
  return 0
}

echo "Probando conexión: psql -h $PG_HOST -p $PG_PORT -U $PG_USER -W $PG_PASSWORD"
# Ejecución principal
if wait_for_postgres; then
  if create_databases_if_not_exist; then
    echo "Proceso completado exitosamente!"
    #exit 0
  else
    echo "Error al crear las bases de datos"
    #exit 1
  fi
else
  exit 1
fi