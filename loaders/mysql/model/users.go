package model

type userType int

const (
	Customer userType = iota
	DormOwner
)

func (u userType) String() string {
	return [...]string{"Customer", "DormOwner"}[u]
}

type User struct {
	UserId   *uint64   `gorm:"primaryKey" json:"userId"`
	UserName *string   `gorm:"not null" json:"userName"`
	Fname    *string   `gorm:"not null" json:"fname"`
	Lname    *string   `gorm:"not null" json:"lname"`
	Email    *string   `gorm:"not null" json:"email"`
	Tel      *string   `gorm:"not null" json:"tel"`
	UserType *userType `gorm:"type:enum('Customer', 'DormOwner')" gorm:"not null" json:"userType"`
}
