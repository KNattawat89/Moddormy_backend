package model

import "time"

type Dorm struct {
	Id             *uint64    `json:"id" gorm:"primaryKey;not null;"`
	DormName       *string    `json:"dorm_name" gorm:"not null"`
	OwnerId        *uint64    `gorm:"primaryKey;not null"`
	Owner          *User      `gorm:"foreignKey:OwnerId;references:Id;not null"`
	HouseNumber    *string    `json:"house_number" gorm:"not null"`
	Street         *string    `json:"street"`
	Soi            *string    `json:"soi"`
	SubDistrict    *string    `json:"sub_district" gorm:"not null"`
	District       *string    `json:"district" gorm:"not null"`
	City           *string    `json:"city" gorm:"not null"`
	Zipcode        *int16     `json:"zipcode" gorm:"not null"`
	Desc           *string    `json:"desc" gorm:"not null"`
	AdvancePayment *int       `json:"advance_payment" gorm:"not null"`
	ElectricPrice  *float32   `json:"electric_price" gorm:"not null"`
	WaterPrice     *float32   `json:"water_price" gorm:"not null"`
	Other          *string    `json:"other" gorm:"not null"`
	LastUpdate     *time.Time `json:"last_update" gorm:"not null"`
	Distant        *float32   `json:"distant" gorm:"not null"` // km
	Pet            *bool      `json:"pet" gorm:"not null"`
	SmokeFree      *bool      `json:"smoke_free" gorm:"not null"`
	Parking        *bool      `json:"parking" gorm:"not null"`
	Lift           *bool      `json:"lift" gorm:"not null"`
	Pool           *bool      `json:"pool" gorm:"not null"`
	Fitness        *bool      `json:"fitness" gorm:"not null"`
	Wifi           *bool      `json:"wifi" gorm:"not null"`
	KeyCard        *bool      `json:"key_card" gorm:"not null"`
	CCTV           *bool      `json:"cctv" gorm:"not null"`
	SecurityGuard  *bool      `json:"security_guard" gorm:"not null"`
	UpdatedAt      *time.Time `json:"updated_at" gorm:"not null"`
}
