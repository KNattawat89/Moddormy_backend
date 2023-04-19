package model

import "time"

type Review struct {
	ReviewID *uint64    `json:"review_id" gorm:"primaryKey" gorm:"AUTO_INCREMENT"`
	DormID   *Dorm      `json:"dorm_id" gorm:"foreignKey:DormID"`
	UserID   *User      `json:"user_id" gorm:"foreignKey:UserID"`
	Review   *string    `json:"review" gorm:"not null"`
	CreateAt *time.Time `json:"create_date" gorm:"not null"`
}
