package payload

type Room struct {
	RoomId      *uint64  `form:"roomId"`
	DormId      *uint64  `form:"dormId"`
	RoomName    *string  `form:"roomName"`
	CoverImage  *string  `form:"coverImage"`
	Price       *float64 `form:"price"`
	Desc        *string  `form:"desc"`
	Size        *string  `form:"size"`
	Airc        *bool    `form:"airc"`
	Furniture   *bool    `form:"furniture"`
	WaterHeater *bool    `form:"waterHeater"`
	Fan         *bool    `form:"fan"`
	Fridge      *bool    `form:"fridge"`
	Bathroom    *bool    `form:"bathroom"`
	TV          *bool    `form:"tv"`
}
