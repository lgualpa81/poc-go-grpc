package models

import "time"

type StockDecreaseLog struct {
	Id           int64     `json:"id" gorm:"primaryKey"`
	OrderId      int64     `json:"order_id"`
	ProductRefer int64     `json:"product_id"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime;type:timestamp with time zone;default:(now() AT TIME ZONE 'UTC')"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime;type:timestamp with time zone;default:(now() AT TIME ZONE 'UTC')"`
}
