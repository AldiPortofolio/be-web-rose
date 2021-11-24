package dbmodels

import "time"

type HistoryClearSession struct {
	ID 			int64 	`json:"id"`
	CreatedBy 	string `json:"createdBy"`
	CreatedAt 	time.Time `json:"createdAt"`
}

func (q *HistoryClearSession) TableName() string {
	return "public.history_clear_session"
}



