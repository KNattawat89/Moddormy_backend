package payload

type Home struct {
	DormId      *uint64  `json:"dorm_id"`
	DormName    *string  `json:"dorm_name"`
	CoverImage  *string  `json:"cover_image"`
	MinPrice    *float64 `json:"min_price"`
	MaxPrice    *float64 `json:"max_price"`
	OverallRate *float64 `json:"overall_rate"`
	FavStatus   *bool    `json:"fav_status"`
}
