package model

type Favorite struct {
	DormID *uint64 `json:"dorm_id" gorm:"not null;primaryKey"`
	Dorm   *Dorm   `json:"dorm" gorm:"foreignKey:DormID;not null"`
	UserID *uint64 `json:"user_id" gorm:"not null;primaryKey"`
	User   *User   `json:"user" gorm:"foreignKey:UserID;not null"`
}
