package model

import "time"

type DormImage struct {
	DormId    *uint64    `json:"dorm_id" gorm:"not null"`
	Dorm      *Dorm      `json:"dorm" gorm:"foreignKey:DormId;references:Id;not null"`
	FileName  *string    `json:"file_name" gorm:"not null"`
	UpdatedAt *time.Time `json:"update_at" gorm:"not null"`
}
