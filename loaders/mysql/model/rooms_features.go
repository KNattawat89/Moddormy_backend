package model

type RoomFeature struct {
	RoomID      *uint64 `json:"room_id" gorm:"not null"`
	Room        *Room   `json:"room" gorm:"foreignKey:RoomID"`
	DormID      *uint64 `json:"dorm_id" gorm:"not null"`
	Dorm        *Dorm   `json:"dorm" gorm:"foreignKey:DormID"`
	Airc        *bool   `json:"airc" gorm:"not null"`
	Furniture   *bool   `json:"furniture" gorm:"not null"`
	WaterHeater *bool   `json:"water_heater" gorm:"not null"`
	Fan         *bool   `json:"fan" gorm:"not null"`
	Fridge      *bool   `json:"fridge" gorm:"not null"`
	Bathroom    *bool   `json:"bathroom" gorm:"not null"`
	TV          *bool   `json:"tv" gorm:"not null"`
}
