package model

import "time"

type RoomImage struct {
	RoomId    *uint64    `json:"room_id" gorm:"not null;primaryKey"`
	Room      *Room      `json:"room" gorm:"foreignKey:RoomId;references:Id;not null"`
	FileName  *string    `json:"file_name" gorm:"not null;primaryKey"`
	UpdatedAt *time.Time `json:"update_at" gorm:"not null"`
}
