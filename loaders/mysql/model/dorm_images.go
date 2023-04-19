package model

type DormImage struct {
	DormID  *uint64 `json:"dorm_id" gorm:"not null"`
	Dorm    *Dorm   `json:"dorm" gorm:"foreignKey:DormID"`
	ImageID *uint64 `json:"image_id" gorm:"not null"`
	Image   *File   `json:"image" gorm:"foreignKey:ImageID"`
}
