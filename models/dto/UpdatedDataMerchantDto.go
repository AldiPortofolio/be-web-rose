package dto

import "time"

type ReqUpdatedDataMerchantDto struct {
	Mid string `json:"mid"`
	StoreName string `json:"storeName"`
	LoanBankCode string `json:"loanBankCode"`
	Status string `json:"status"`
	UpdatedBy string `json:"updatedBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	Limit       		  int `json:"limit"`
	Page        		  int `json:"page"`
}


type ResUpdatedDataMerchantDto struct {
	Mid          string    `json:"mid"`
	StoreName    string    `json:"storeName"`
	LoanBankCode string    `json:"loanBankCode"`
	Status       string    `json:"status"`
	UpdatedData  interface{}    `json:"updatedData"`
	UpdatedBy    string    `json:"updatedBy"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type ReqUpdateDataMerchantDto struct {
	ID int64 `json:"id"`
} 