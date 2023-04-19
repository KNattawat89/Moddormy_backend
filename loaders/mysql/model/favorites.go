package model

type Favorite struct {
	DormID *Dorm `json:"dorm_id" gorm:"foreignKey:DormID"`
	UserID *User `json:"user_id" gorm:"foreignKey:UserID"`
}
