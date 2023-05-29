package payload

import "time"

type DormProfile struct {
	DormId     *uint64    `json:"dormId"`
	DormName   *string    `json:"dormName"`
	CoverImage *string    `json:"coverImage"`
	CreatedAt  *time.Time `json:"createdAt"`
}
