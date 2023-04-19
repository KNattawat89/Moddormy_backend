package model

type RoomImage struct {
	DormID  *uint64 `json:"dorm_id" gorm:"not null"`
	Dorm    *Dorm   `json:"dorm" gorm:"foreignKey:DormID"`
	RoomID  *uint64 `json:"room_id" gorm:"not null"`
	Room    *Room   `json:"room" gorm:"foreignKey:RoomID"`
	ImageID *uint64 `json:"image_id" gorm:"not null"`
	Image   *File   `json:"image" gorm:"foreignKey:ImageID"`
}
