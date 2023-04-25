package model

import "time"

type Review struct {
	DormId         *uint64    `json:"dorm_id" gorm:"not null;primaryKey"`
	Dorm           *Dorm      `json:"dorm" gorm:"foreignKey:DormId;references:Id;not null"`
	UserId         *uint64    `json:"user_id" gorm:"not null;primaryKey"`
	User           *User      `json:"user" gorm:"foreignKey:UserId;references:Id;not null"`
	Review         *string    `json:"review" gorm:"not null"`
	RatingOverall  *int       `json:"rating_overall" gorm:"not null"`
	RatingPrice    *int       `json:"price" gorm:"not null"`
	RatingLocation *int       `json:"location" gorm:"not null"`
	RatingFacility *int       `json:"facility" gorm:"not null"`
	RatingSanitary *int       `json:"sanitary" gorm:"not null"`
	RatingSecurity *int       `json:"security" gorm:"not null"`
	CreateAt       *time.Time `json:"create_date" gorm:"not null"`
}
