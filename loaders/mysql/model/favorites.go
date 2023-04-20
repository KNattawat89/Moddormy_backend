package model

type Favorite struct {
	DormId *uint64 `json:"dorm_id" gorm:"not null;primaryKey;"`
	Dorm   *Dorm   `json:"dorm" gorm:"foreignKey:DormId;references:Id;not null"`
	UserId *uint64 `json:"user_id" gorm:"not null;primaryKey"`
	User   *User   `json:"user" gorm:"foreignKey:UserId;references:Id;not null"`
}
