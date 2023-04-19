package model

import "time"

type Dorm struct {
	ID         *uint64    `json:"dorm_id" gorm:"primaryKey"`
	DormName   *string    `json:"dorm_name" gorm:"not null"`
	OwnerId    *uint64    `gorm:"not null"`
	Owner      *User      `gorm:"foreignKey:OwnerId"`
	Address    *string    `json:"address" gorm:"not null"`
	Desc       *string    `json:"desc" gorm:"not null"`
	Contract   *string    `json:"contract" gorm:"not null"`
	LastUpdate *time.Time `json:"last_update" gorm:"not null"`
}
