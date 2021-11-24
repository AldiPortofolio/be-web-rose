package dto

import "time"

type ReqMerchantBankAccountDto struct {
	ID int64 `json:"id"`
	Mid string `json:"mid"`
	BankCode string `json:"bankCode"`
	AccountNumber string `json:"accountNumber"`
	AccountName string `json:"accountName"`
	Notes string `json:"notes"`
	Status string `json:"status"`
	Limit       		  int `json:"limit"`
	Page        		  int `json:"page"`
}

type ReqApprovalMerchantBankAccountDto struct {
	ID int64 `json:"id"`
	Notes string `json:"notes"`
}

type ResMerchantBankAccount struct {
	ID int64 `json:"id"`
	Mid string `json:"mid"`
	BankCode string `json:"bankCode"`
	FullName string `json:"fullName"`
	ShortName string `json:"shortName"`
	AccountNumber string `json:"accountNumber"`
	AccountName string `json:"accountName"`
	Notes string `json:"notes"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy string `json:"updatedBy"`
	PushNotifStatus string `json:"pushNotifStatus"`
}

type ReqValidationBankAccount struct {
	AccountNo string `json:"accountNo"`
	BankCode string `json:"bankCode"`
	Mid string `json:"mid"`
}

type ResValidationBankAccountDto struct {
	AccountNo string `json:"accountNo"`
	AccountName string `json:"accountName"`
	AccountBankName string `json:"accountBankName"`
	BankCode string `json:"bankCode"`
	OwnerFirstName string `json:"ownerFirstName"`
	OwnerLastName string `json:"ownerLastName"`
}