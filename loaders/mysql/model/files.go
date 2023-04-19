package model

type File struct {
	FileID *uint64 `json:"file_id" gorm:"primaryKey" gorm:"AUTO_INCREMENT"`
	Name   *string `json:"name" gorm:"not null"`
	Type   *string `json:"type" gorm:"not null"`
	Data   *byte   `json:"data" gorm:"not null"`
}
