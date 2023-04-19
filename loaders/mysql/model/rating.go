package model

type Rating struct {
	ReviewID *Review `json:"review_id" gorm:"foreignKey:ReviewID"`
	Price    *int    `json:"price" gorm:"not null"`
	Location *int    `json:"location" gorm:"not null"`
	Facility *int    `json:"facility" gorm:"not null"`
	Sanitary *int    `json:"sanitary" gorm:"not null"`
	Security *int    `json:"security" gorm:"not null"`
}
