package payload

type Filter struct {
	UserId     *string   `json:"userId"`
	MinPrice   *uint64   `json:"min_price"`
	MaxPrice   *uint64   `json:"max_price"`
	Distant    *float32  `json:"distant"`
	Rate       *string   `json:"rate"`
	Facilities []*string `json:"facilities"`
}
