package payload

type Room struct {
	RoomId      *uint64      `json:"roomId"`
	DormId      *uint64      `json:"dormId"`
	RoomName    *string      `json:"roomName"`
	CoverImage  *string      `json:"coverImage"`
	Price       *float64     `json:"price"`
	Desc        *string      `json:"desc"`
	Size        *string      `json:"size"`
	RoomFeature *RoomFeature `json:"roomFeature"`
}
