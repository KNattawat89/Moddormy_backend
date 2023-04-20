package model

import "Moddormy_backend/types/enum"

type User struct {
	Id          *uint64    `gorm:"primaryKey;not null;index:idx_id" json:"id"`
	UserName    *string    `gorm:"type:VARCHAR(255);not null" json:"username"`
	Password    *string    `gorm:"type:VARCHAR(255); not null" json:"password"`
	Fname       *string    `gorm:"type:VARCHAR(255);not null" json:"fname"`
	Lname       *string    `gorm:"type:VARCHAR(255);not null" json:"lname"`
	Email       *string    `gorm:"type:VARCHAR(255); index:email,unique; not null" json:"email"`
	Tel         *string    `gorm:"not null" json:"tel"`
	LineID      *string    `json:"line_id" gorm:"type:VARCHAR(255);"`
	UserType    *enum.User `gorm:"type:ENUM('Customer', 'DormOwner');not null" json:"user_type"`
	UnusedField string     `gorm:"-" json:"-"`
}
