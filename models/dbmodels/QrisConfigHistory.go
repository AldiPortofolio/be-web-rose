package dbmodels

import "time"

type QrisConfigHistory struct {
	ID 					int64  `json:"id" gorm:"id"`
	QrisConfigID 		int64 `json:"qrisConfigId"`
	InstitutionID		string `json:"institutionId"`
	IssuerName			string `json:"issuerName"`
	TransactionType 	string `json:"transactionType"`
	Status 				int32  `json:"status"`
	CreatedAt 			time.Time `json:"createdAt"`
	CreatedBy			string `json:"createdBy"`
}

func (q *QrisConfigHistory) TableName() string {
	return "public.qris_config_history"
}