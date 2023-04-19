package model

type RoomImage struct {
	DormID  *Dorm `json:"dorm_id" gorm:"foreignKey:DormID"`
	RoomID  *Room `json:"room_id" gorm:"foreignKey:RoomID"`
	ImageID *File `json:"image_id" gorm:"foreignKey:ImageID"`
}
