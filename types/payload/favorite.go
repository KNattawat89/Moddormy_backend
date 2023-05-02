package payload

type Favorite struct {
	DormId *uint64 `form:"dormId"`
	UserId *uint64 `form:"userId"`
}