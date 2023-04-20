package model

type File struct {
	FileID   *uint64 `json:"file_id" gorm:"not null;primaryKey"`
	FileName *string `json:"name" gorm:"not null"`
	DormID   *uint64 `json:"dorm_id"`
	Dorm     *Dorm   `json:"dorm"`
	RoomID   *uint64 `json:"room_id"`
	Room     *Room   `json:"room"`
	UserID   *uint64 `json:"user_id" gorm:"not null;primaryKey"`
	User     *User   `json:"user" gorm:"foreignKey:UserID;not null'"`
}
