package payload

type DormFeature struct {
	Pet           *bool `json:"pet"`
	SmokeFree     *bool `json:"smokeFree"`
	Parking       *bool `json:"parking"`
	Lift          *bool `json:"lift"`
	Pool          *bool `json:"pool"`
	Fitness       *bool `json:"fitness"`
	Wifi          *bool `json:"wifi"`
	KeyCard       *bool `json:"keyCard"`
	CCTV          *bool `json:"cctv"`
	SecurityGuard *bool `json:"securityGuard"`
}
