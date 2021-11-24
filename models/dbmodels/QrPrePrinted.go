package dbmodels

import "time"

type QrPrePrinted struct {
	ID 					int64  `json:"id" gorm:"id"`
	TotalReq 			int32  `json:"totalReq"`
	City				string `json:"city"`
	PostalCode 			string `json:"postalCode"`
	Mcc 				string `json:"mcc"`
	User 				string `json:"user"`
	Key 				string `json:"key"`
	Status 				string `json:"status"`
	Notes 				string `json:"notes"`
	Path 				string `json:"path"`
	CreatedAt			time.Time `json:"createdAt"`
	UpdatedAt 			time.Time `json:"updatedAt"`
}

func (q *QrPrePrinted) TableName() string {
	return "public.qr_preprinted"
}
