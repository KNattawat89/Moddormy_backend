package model

type Room struct {
	RoomID   *uint64  `json:"room_id" gorm:"primaryKey"`
	DormID   *uint64  `json:"dorm_id" gorm:"not null"`
	Dorm     *Dorm    `json:"dorm" gorm:"foreignKey:DormID"`
	RoomName *string  `json:"room_name" gorm:"not null"`
	Price    *float64 `json:"price" gorm:"not null"`
	Desc     *string  `json:"desc" gorm:"not null"`
	Size     *string  `json:"size" gorm:"not null"`
}
