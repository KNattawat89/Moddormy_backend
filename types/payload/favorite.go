package payload

type Favorite struct {
	DormId *uint64 `form:"dormId"`
	UserId *string `form:"userId"`
}
