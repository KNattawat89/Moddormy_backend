package payload

type Dorm struct {
	DormId         *uint64      `json:"dormId"`
	DormName       *string      `json:"dormName"`
	UserId         *string      `json:"userId"`
	CoverImage     *string      `json:"coverImage"`
	HouseNumber    *string      `json:"houseNumber"`
	Street         *string      `json:"street"`
	Soi            *string      `json:"soi"`
	SubDistrict    *string      `json:"subDistrict"`
	District       *string      `json:"district"`
	City           *string      `json:"city"`
	Zipcode        *int32       `json:"zipcode"`
	Desc           *string      `json:"desc"`
	AdvancePayment *int         `json:"advancePayment"`
	ElectricPrice  *float64     `json:"electricPrice"`
	WaterPrice     *float64     `json:"waterPrice"`
	Other          *string      `json:"other"`
	Distant        *float64     `json:"distant"`
	DormFeatures   *DormFeature `json:"dormFeatures"`
}
