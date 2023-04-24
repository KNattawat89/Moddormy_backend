package model

type RoomImage struct {
	RoomId   *uint64 `json:"room_id" gorm:"not null;primaryKey"`
	Room     *Room   `json:"room" gorm:"foreignKey:RoomId;references:Id;not null"`
	FileId   *uint64 `json:"image_id" gorm:"not null;primaryKey"`
	FileName *string `json:"file_name" gorm:"not null"`
}
