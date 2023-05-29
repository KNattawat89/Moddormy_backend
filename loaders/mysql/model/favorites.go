package model

import "time"

type Favorite struct {
	DormId    *uint64    `json:"dorm_id" gorm:"not null;primaryKey;"`
	Dorm      *Dorm      `json:"dorm" gorm:"foreignKey:DormId;references:Id;not null"`
	UserId    *string    `json:"user_id" gorm:"not null;primaryKey"`
	User      *User      `json:"user" gorm:"foreignKey:UserId;references:Id;not null"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"not null"`
}
