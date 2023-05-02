package payload

import "Moddormy_backend/types/enum"

type RegisterAccount struct {
	UserName  *string    `json:"username"`
	PassWord  *string    `json:"password"`
	FirstName *string    `json:"fname"`
	LastName  *string    `json:"lname"`
	Email     *string    `json:"email"`
	Account   *enum.User `json:"account"`
	Tel       *string    `json:"tel"`
}
