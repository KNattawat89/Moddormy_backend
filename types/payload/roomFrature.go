package payload

type RoomFeature struct {
	Airc        *bool `json:"airc"`
	Furniture   *bool `json:"furniture"`
	WaterHeater *bool `json:"waterHeater"`
	Fan         *bool `json:"fan"`
	Fridge      *bool `json:"fridge"`
	Bathroom    *bool `json:"bathroom"`
	TV          *bool `json:"tv"`
}
