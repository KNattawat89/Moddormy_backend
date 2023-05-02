package payload

type DormSearch struct {
	DormId     *uint64  `json:"dorm_id"`
	DormName   *string  `json:"dorm_name"`
	CoverImage *string  `json:"cover_image"`
	Rating     *float32 `json:"rating"`
	MinPrice   *float64 `json:"min_price"`
	MaxPrice   *float64 `json:"max_price"`
	//IsFav      *bool   `json:"is_fav"`
}
