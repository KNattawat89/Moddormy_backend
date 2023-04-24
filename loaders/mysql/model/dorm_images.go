package model

type DormImage struct {
	DormId   *uint64 `json:"dorm_id" gorm:"not null;primaryKey"`
	Dorm     *Dorm   `json:"dorm" gorm:"foreignKey:DormId;references:Id;not null"`
	FileId   *uint64 `json:"file_id_id" gorm:"not null;primaryKey"`
	FileName *string `json:"file_name" gorm:"not null"`
}
