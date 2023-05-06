package payload

type DormSearch struct {
	DormId      *uint64  `json:"dorm_id"`
	DormName    *string  `json:"dorm_name"`
	CoverImage  *string  `json:"cover_image"`
	OverallRate *float64 `json:"overall_rate"`
	MinPrice    *float64 `json:"min_price"`
	MaxPrice    *float64 `json:"max_price"`
	FavStatus   *bool    `json:"fav_status"`
}
