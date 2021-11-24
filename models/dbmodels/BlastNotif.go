package dbmodels

import "time"

type BlastNotif struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Target string `json:"target"`
	Desc string `json:"desc"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
}

func (t *BlastNotif)TableName() string {
	return "public.blast_notification"
}