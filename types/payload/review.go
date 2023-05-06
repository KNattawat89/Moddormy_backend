package payload

type DormReview struct {
	UserId         *string `json:"userId"`
	DormId         *uint64 `json:"dormId"`
	Review         *string `json:"review"`
	RatingPrice    *int    `json:"ratingPrice"`
	RatingLocation *int    `json:"ratingLocation"`
	RatingFacility *int    `json:"ratingFacility"`
	RatingSanitary *int    `json:"ratingSanitary"`
	RatingSecurity *int    `json:"ratingSecurity"`
	RatingOverall  *int    `json:"ratingOverall"`
}
