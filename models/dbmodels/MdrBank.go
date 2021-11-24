package dbmodels

import "time"

type MdrBank struct {
	ID 			int64 `json:"id"`
	BankCode 	string `json:"bankCode"`
	BankName	string `json:"bankName"`
	DokuBankCode	string `json:"dokuBankCode"`
	Status		string 	`json:"status"`
	Seq			int32	`json:"seq"`
	AcquiringStatus 	string `json:"acquiringStatus"`
	UpdatedAt 			time.Time `json:"updatedAt"`
	UpdatedBy			string `json:"updatedBy"`

}

func (q *MdrBank) TableName() string {
	return "public.mdr_bank"
}
