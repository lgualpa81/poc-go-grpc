package db

import (
	"log"
	"time"

	"github.com/lgualpa81/go-grpc-auth-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h Handler) Where(user *models.User) {
	panic("unimplemented")
}

func Init(url string) (h Handler, err error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.User{})
	return Handler{db}, err
}
