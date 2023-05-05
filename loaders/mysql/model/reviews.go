package model

import "time"

type Review struct {
	Id             *uint64    `json:"id" gorm:"not null;primaryKey"`
	DormId         *uint64    `json:"dorm_id" gorm:"not null;primaryKey;autoIncrement:false;"`
	Dorm           *Dorm      `json:"dorm" gorm:"foreignKey:DormId;references:Id;not null"`
	UserId         *string    `json:"user_id" gorm:"not null;primaryKey;autoIncrement:false;"`
	User           *User      `json:"user" gorm:"foreignKey:UserId;references:Id;not null"`
	Review         *string    `json:"review" gorm:"not null"`
	RatingOverall  *float32   `json:"rating_overall" gorm:"not null"`
	RatingPrice    *int       `json:"price" gorm:"not null"`
	RatingLocation *int       `json:"location" gorm:"not null"`
	RatingFacility *int       `json:"facility" gorm:"not null"`
	RatingSanitary *int       `json:"sanitary" gorm:"not null"`
	RatingSecurity *int       `json:"security" gorm:"not null"`
	CreateAt       *time.Time `json:"create_date" gorm:"not null"`
}
