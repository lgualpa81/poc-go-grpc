# Builder para protobuf (stage 1)
FROM golang:1.24.4-alpine3.22 AS protobuf-builder

WORKDIR /app

# Instalar dependencias (protoc, plugins de Go)
RUN apk add --no-cache protobuf-dev make && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Copiar solo lo necesario para generar protos
COPY pkg/auth/pb/ ./pkg/auth/pb/
COPY pkg/order/pb/ ./pkg/order/pb/
COPY pkg/product/pb/ ./pkg/product/pb/
COPY Makefile .

# Generar stubs de gRPC
RUN make proto

#Builder para aplicacion (stage 2)
FROM golang:1.24.4-alpine AS app-builder
WORKDIR /app
# Copiar los archivos generados en la etapa 1
COPY --from=protobuf-builder /app/pkg/auth/pb/gen/ ./pkg/pb/auth/gen/
COPY --from=protobuf-builder /app/pkg/order/pb/gen/ ./pkg/pb/order/gen/
COPY --from=protobuf-builder /app/pkg/product/pb/gen/ ./pkg/pb/product/gen/

# Copiar todo el código (incluyendo cmd/, pkg/, etc.)
COPY . .

# Compilar el servicio (asumiendo que `make server` depende de un build previo)
RUN go build -o /app/server ./cmd/main.go

# Imagen final mínima (etapa 3)
FROM alpine:3.22
WORKDIR /app
COPY --from=app-builder /app/server .

EXPOSE $PORT
# Comando para iniciar el servicio (equivalente a `make server`)
CMD ["./server"]