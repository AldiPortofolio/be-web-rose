package dbmodels

import "time"

type HistoryAppVersion struct {
	ID 			int64 	`json:"id"`
	AppName 	string `json:"appName"`
	Key			string `json:"key"`
	Version 	string `json:"version"`
	CreatedBy 	string `json:"createdBy"`
	CreatedAt 	time.Time `json:"createdAt"`
}

func (q *HistoryAppVersion) TableName() string {
	return "public.history_app_version"
}
