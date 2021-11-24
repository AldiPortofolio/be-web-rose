package models

import "time"

type PushNotifBankAccount struct {
	StorePhoneNumber string `json:"storePhoneNumber"`
	BankName string `json:"bankName"`
	AccountNumber string `json:"accountNumber"`
	StatusApprove string `json:"statusApprove"`
	Notes string `json:"notes"`
	TransactionTime time.Time `json:"transactionTime"`
}
