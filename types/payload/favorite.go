package payload

type Favorite struct {
	DormId *uint64 `json:"dorm_id"`
	UserId *string `json:"user_id"`
}
