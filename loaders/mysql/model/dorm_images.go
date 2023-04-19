package model

type DormImage struct {
	DormID  *Dorm `json:"dorm_id" gorm:"foreignKey:DormID"`
	ImageID *File `json:"image_id" gorm:"foreignKey:ImageID"`
}
