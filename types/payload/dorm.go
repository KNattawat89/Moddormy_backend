package payload

type Dorm struct {
	DormId         *uint64  `form:"dormId"`
	DormName       *string  `form:"dormName"`
	UserId         *string  `form:"userId"`
	CoverImage     *string  `form:"coverImage"`
	HouseNumber    *string  `form:"houseNumber"`
	Street         *string  `form:"street"`
	Soi            *string  `form:"soi"`
	SubDistrict    *string  `form:"subDistrict"`
	District       *string  `form:"district"`
	City           *string  `form:"city"`
	Zipcode        *int16   `form:"zipcode"`
	Desc           *string  `form:"desc"`
	AdvancePayment *int     `form:"advancePayment"`
	ElectricPrice  *float64 `form:"electricPrice"`
	WaterPrice     *float64 `form:"waterPrice"`
	Other          *string  `form:"other"`
	Distant        *float64 `form:"distant"`
	Pet            *bool    `form:"pet"`
	SmokeFree      *bool    `form:"smokeFree"`
	Parking        *bool    `form:"parking"`
	Lift           *bool    `form:"lift"`
	Pool           *bool    `form:"pool"`
	Fitness        *bool    `form:"fitness"`
	Wifi           *bool    `form:"wifi"`
	KeyCard        *bool    `form:"keyCard"`
	CCTV           *bool    `form:"cctv"`
	SecurityGuard  *bool    `form:"securityGuard"`
}
