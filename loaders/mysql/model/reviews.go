package model

import "time"

type Review struct {
	Id             *uint64    `json:"id" gorm:"not null;primaryKey"`
	DormId         *uint64    `json:"dormId" gorm:"not null;primaryKey;autoIncrement:false;"`
	Dorm           *Dorm      `json:"dorm" gorm:"foreignKey:DormId;references:Id;not null"`
	UserId         *string    `json:"userId" gorm:"not null;primaryKey;autoIncrement:false;"`
	User           *User      `json:"user" gorm:"foreignKey:UserId;references:Id;not null"`
	Review         *string    `json:"review" gorm:"not null"`
	RatingOverall  *float64   `json:"ratingOverall" gorm:"not null"`
	RatingPrice    *int       `json:"ratingPrice" gorm:"not null"`
	RatingLocation *int       `json:"ratingLocation" gorm:"not null"`
	RatingFacility *int       `json:"ratingFacility" gorm:"not null"`
	RatingSanitary *int       `json:"ratingSanitary" gorm:"not null"`
	RatingSecurity *int       `json:"ratingSecurity" gorm:"not null"`
	CreatedAt      *time.Time `json:"createdAt" gorm:"not null"`
}
