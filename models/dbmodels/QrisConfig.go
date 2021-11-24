package dbmodels

import "time"

type QrisConfig struct {
	ID 					int64  `json:"id" gorm:"id"`
	InstitutionID		string `json:"institutionId"`
	IssuerName			string `json:"issuerName"`
	TransactionType 	string `json:"transactionType"`
	Status 				int32  `json:"status"`
	CreatedAt 			time.Time `json:"createdAt"`
	CreatedBy			string `json:"createdBy"`
	UpdatedAt 			time.Time `json:"updatedAt"`
	UpdatedBy			string `json:"updatedBy"`
}

func (q *QrisConfig) TableName() string {
	return "public.qris_config"
}