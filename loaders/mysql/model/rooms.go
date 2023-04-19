package model

type Room struct {
	RoomID   *uint64  `json:"room_id" gorm:"primaryKey" gorm:"AUTO_INCREMENT"`
	DormID   *Dorm    `json:"dorm_id" gorm:"foreignKey:DormID"`
	RoomName *string  `json:"room_name" gorm:"not null"`
	Price    *float64 `json:"price" gorm:"not null"`
	Desc     *string  `json:"desc" gorm:"not null"`
	Size     *string  `json:"size" gorm:"not null"`
}
