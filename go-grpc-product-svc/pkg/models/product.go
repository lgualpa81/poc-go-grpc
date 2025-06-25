package models

import "time"

type Product struct {
	Id                int64            `json:"id" gorm:"primaryKey"`
	Name              string           `json:"name"`
	Stock             int64            `json:"stock"`
	Price             float64          `json:"price"`
	CreatedAt         time.Time        `json:"created_at" gorm:"autoCreateTime;type:timestamp with time zone;default:(now() AT TIME ZONE 'UTC')"`
	UpdatedAt         time.Time        `json:"updated_at" gorm:"autoUpdateTime;type:timestamp with time zone;default:(now() AT TIME ZONE 'UTC')"`
	StockDecreaseLogs StockDecreaseLog `gorm:"foreignKey:ProductRefer"`
}
