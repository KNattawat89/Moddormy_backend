package payload

type Favorite struct {
	DormId *uint64 `json:"dormId"`
	UserId *string `json:"userId"`
}
