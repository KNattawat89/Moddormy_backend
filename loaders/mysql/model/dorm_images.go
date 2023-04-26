package model

import "time"

type DormImage struct {
	DormId    *uint64    `json:"dorm_id" gorm:"not null;primaryKey"`
	Dorm      *Dorm      `json:"dorm" gorm:"foreignKey:DormId;references:Id;not null"`
	FileName  *string    `json:"file_name" gorm:"not null;primaryKey"`
	UpdatedAt *time.Time `json:"update_at" gorm:"not null"`
}
