package model

type DormImage struct {
	DormID  *uint64 `json:"dorm_id" gorm:"not null;primaryKey"`
	Dorm    *Dorm   `json:"dorm" gorm:"foreignKey:DormID;not null"`
	ImageID *uint64 `json:"image_id" gorm:"not null;primaryKey"`
	Image   *File   `json:"image" gorm:"foreignKey:ImageID;not null"`
}
