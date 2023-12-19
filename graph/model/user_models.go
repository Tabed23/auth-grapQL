package model

import "time"

type User struct {
	ID       string `json:"id" gorm:"type:varchar(255);primaryKey"`
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" gorm:"type:varchar(100);not null;unique"`
	Password string `json:"password" gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
