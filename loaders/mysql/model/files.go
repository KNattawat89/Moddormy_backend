package model

type File struct {
	FileID   *uint64 `json:"file_id" gorm:"primaryKey" gorm:"AUTO_INCREMENT"`
	FileName *string `json:"name" gorm:"not null"`
}
