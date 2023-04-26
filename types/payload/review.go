package payload

type DormReview struct {
	UserId         *uint64 `form:"userId"`
	DormId         *uint64 `form:"dormId"`
	Review         *string `form:"review"`
	RatingPrice    *int    `form:"ratingPrice"`
	RatingLocation *int    `form:"ratingLocation"`
	RatingFacility *int    `form:"ratingFacility"`
	RatingSanitary *int    `form:"ratingSanitary"`
	RatingSecurity *int    `form:"ratingSecurity"`
	RatingOverall  *int    `form:"ratingOverall"`
}
