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
	UserId   *uint64   `gorm:"primaryKey;not null" json:"user_id"`
	UserName *string   `gorm:"not null" json:"username"`
	Password *string   `gorm:"not null" json:"password"`
	Fname    *string   `gorm:"not null" json:"fname"`
	Lname    *string   `gorm:"not null" json:"lname"`
	Email    *string   `gorm:"not null" json:"email"`
	Tel      *string   `gorm:"not null" json:"tel"`
	LineID   *string   `json:"line_id"`
	UserType *userType `gorm:"type:enum('Customer', 'DormOwner');not null" json:"user_type"`
}
