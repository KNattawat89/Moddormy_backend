package model

type Favorite struct {
	DormID *uint64 `json:"dorm_id" gorm:"not null"`
	Dorm   *Dorm   `json:"dorm" gorm:"foreignKey:DormID"`
	UserID *uint64 `json:"user_id" gorm:"not null"`
	User   *User   `json:"user" gorm:"foreignKey:UserID"`
}
