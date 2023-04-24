package payload

type UploadDorm struct {
	DormId *uint64 `form:"dormId"`
}

type UploadRoom struct {
	RoomId *uint64 `form:"roomId"`
}
