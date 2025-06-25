package models

import "time"

type Order struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	Price     float64   `json:"price"`
	Quantity  int64     `json:"quantity"`
	ProductId int64     `json:"product_id"`
	UserId    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;type:timestamp with time zone;default:(now() AT TIME ZONE 'UTC')"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;type:timestamp with time zone;default:(now() AT TIME ZONE 'UTC')"`
}
