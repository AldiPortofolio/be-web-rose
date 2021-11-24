package dto

import "time"

type ReqMdrDto struct {
	ID 					int64  `json:"id" gorm:"id"`
	BankCode			string `json:"bankCode"`
	BankName			string `json:"bankName"`
	DokuBankCode		string `json:"dokuBankCode"`
	Status 				string `json:"status"`
	Seq					int32 `json:"seq"`
	UpdatedAt 			time.Time `json:"updatedAt"`
	UpdatedBy			string `json:"updatedBy"`
	Limit       		  int `json:"limit"`
	Page        		  int `json:"page"`
	AcquiringStatus 	string `json:"acquiringStatus"`
}
