package payload

type Home struct {
	DormId         	*uint64  `form:"dormId"`
	DormName 		*string  `form:"dorm_name`
	CoverImage    	*string  `form:"coverImage"`
	MinPrice		*float64 `form:"min_price"`
	MaxPrice		*float64 `form:"max_price"`
	OverallRate		*float64 `form:"overall_rate"`
	FavStatus 		*bool    `form:"fav_status"`
}
