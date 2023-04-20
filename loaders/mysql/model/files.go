package model

type File struct {
	Id       *uint64 `json:"file_id" gorm:"not null;primaryKey"`
	FileName *string `json:"name" gorm:"not null"`
	DormId   *uint64 `json:"dorm_id"`
	Dorm     *Dorm   `json:"dorm" gorm:"foreignKey:DormId;references:Id;not null"`
	RoomId   *uint64 `json:"room_id"`
	Room     *Room   `json:"room" gorm:"foreignKey:RoomId;references:Id;not null"`
	UserId   *uint64 `json:"user_id" gorm:"not null;primaryKey"`
	User     *User   `json:"user" gorm:"foreignKey:UserId;references:Id;not null'"`
}
