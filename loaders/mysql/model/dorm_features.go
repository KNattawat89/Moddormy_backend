package model

type DormFeature struct {
	DormID         *Dorm    `json:"dorm_id" gorm:"foreignKey:DormID"`
	Distant        *float32 `json:"distant" gorm:"not null"`
	Pet            *bool    `json:"pet" gorm:"not null"`
	Smoking        *bool    `json:"smoking" gorm:"not null"`
	Parking        *bool    `json:"parking" gorm:"not null"`
	Lift           *bool    `json:"lift" gorm:"not null"`
	Pool           *bool    `json:"pool" gorm:"not null"`
	Fitness        *bool    `json:"fitness" gorm:"not null"`
	Wifi           *bool    `json:"wifi" gorm:"not null"`
	Keycard        *bool    `json:"keycard" gorm:"not null"`
	CCTV           *bool    `json:"cctv" gorm:"not null"`
	Security_guard *bool    `json:"security_guard" gorm:"not null"`
}
