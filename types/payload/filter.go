package payload

type Filter struct{
	UserId			*uint64		`form:"userId"`
	MinPrice		*uint64 	`form:"min_price"`
	MaxPrice		*uint64 	`form:"max_price"`
	Distant			*float32 	`form:"distant"`
	Rate			*string		`form:"rate"`
	Facilities		[]*string	`form:"facilities"`
}