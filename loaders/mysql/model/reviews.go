package model

import "time"

type Review struct {
	ID       *uint64    `json:"review_id" gorm:"primaryKey"`
	DormID   *uint64    `json:"dorm_id" gorm:"not null"`
	Dorm     *Dorm      `json:"dorm" gorm:"foreignKey:DormID"`
	UserID   *uint64    `json:"user_id" gorm:"not null"`
	User     *User      `json:"user" gorm:"foreignKey:UserID"`
	Review   *string    `json:"review" gorm:"not null"`
	CreateAt *time.Time `json:"create_date" gorm:"not null"`
}
