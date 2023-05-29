package payload

import "Moddormy_backend/types/enum"

type RegisterAccount struct {
	UserId    *string    `json:"userId"`
	UserName  *string    `json:"username"`
	FirstName *string    `json:"fname"`
	LastName  *string    `json:"lname"`
	Email     *string    `json:"email"`
	Account   *enum.User `json:"account"`
}
