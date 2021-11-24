package dbmodels

import (
	"time"
)


type MerchantBankAccount struct {
	ID int64 `json:"id"`
	Mid string `json:"mid"`
	BankCode string `json:"bankCode"`
	AccountNumber string `json:"accountNumber"`
	AccountName string `json:"accountName"`
	Notes string `json:"notes"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy string `json:"updatedBy"`
	PushNotifStatus string `json:"pushNotifStatus"`
	PushNotifData string `json:"pushNotifData"`
}

func (t *MerchantBankAccount) TableName() string {
	return "public.merchant_bank_account"
}
