package model

type Room struct {
	Id          *uint64  `json:"room_id" gorm:"primaryKey;not null"`
	DormId      *uint64  `json:"dorm_id" gorm:"primaryKey;not null"`
	Dorm        *Dorm    `json:"dorm" gorm:"foreignKey:DormId;references:Id;not null"`
	RoomName    *string  `json:"room_name" gorm:"not null"`
	Price       *float64 `json:"price" gorm:"not null"`
	Desc        *string  `json:"desc" gorm:"not null"`
	Size        *string  `json:"size" gorm:"not null"`
	Airc        *bool    `json:"airc" gorm:"not null"`
	Furniture   *bool    `json:"furniture" gorm:"not null"`
	WaterHeater *bool    `json:"water_heater" gorm:"not null"`
	Fan         *bool    `json:"fan" gorm:"not null"`
	Fridge      *bool    `json:"fridge" gorm:"not null"`
	Bathroom    *bool    `json:"bathroom" gorm:"not null"`
	TV          *bool    `json:"tv" gorm:"not null"`
}

//autoIncrement:false;
