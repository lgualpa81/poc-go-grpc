package models

import "time"

type User struct {
	Id        int64      `json:"id" gorm:"primaryKey"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime;type:timestamp with time zone;default:(now() AT TIME ZONE 'UTC')"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime;type:timestamp with time zone;default:(now() AT TIME ZONE 'UTC')"`
	LastLogin *time.Time `json:"last_login" gorm:"type:timestamp with time zone"`
}
