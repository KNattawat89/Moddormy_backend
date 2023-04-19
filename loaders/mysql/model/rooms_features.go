package model

type RoomFeature struct {
	RoomID      *Room `json:"room_id" gorm:"foreignKey:RoomID"`
	DormID      *Dorm `json:"dorm_id" gorm:"foreignKey:DormID"`
	Airc        *bool `json:"airc" gorm:"not null"`
	Furniture   *bool `json:"furniture" gorm:"not null"`
	WaterHeater *bool `json:"water_heater" gorm:"not null"`
	Fan         *bool `json:"fan" gorm:"not null"`
	Fridge      *bool `json:"fridge" gorm:"not null"`
	Bathroom    *bool `json:"bathroom" gorm:"not null"`
	TV          *bool `json:"tv" gorm:"not null"`
}
