package dbmodels

import "time"

type DepositBank struct {
	ID				int 		`json:"id" gorm:"id"`
	Mid 			string		`json:"mid" gorm:"mid"`
	PhoneNumber		string 		`json:"phone_number" gorm:"phone_number"`
	BankName		string 		`json:"bank_name" gorm:"bank_name"`
	BankCode		string		`json:"bank_code" gorm:"bank_code"`
	AccountNumber	string 		`json:"account_number" gorm:"account_number"`
	AccountName		string 		`json:"account_name" gorm:"account_name"`
	Notes 			string 		`json:"notes" gorm:"notes"`
	CreatedAt		time.Time 	`json:"created_at" gorm:"created_at"`
	UpdatedAt		time.Time 	`json:"updated_at" gorm:"updated_at"`
}

// TableName ...
func (o *DepositBank) TableName() string {
	return "public.deposit_bank"
}
