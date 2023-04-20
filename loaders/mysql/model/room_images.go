package model

type RoomImage struct {
	DormId  *uint64 `json:"dorm_id" gorm:"not null;primaryKey"`
	Dorm    *Dorm   `json:"dorm" gorm:"foreignKey:DormId;references:Id;not null"`
	RoomId  *uint64 `json:"room_id" gorm:"not null;primaryKey"`
	Room    *Room   `json:"room" gorm:"foreignKey:RoomId;references:Id;not null"`
	ImageId *uint64 `json:"image_id" gorm:"not null;primaryKey"`
	Image   *File   `json:"image" gorm:"foreignKey:ImageId;references:Id;not null"`
}
