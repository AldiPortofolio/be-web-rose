package dto

type ReqDepositBank struct {
	Mid				string `json:"mid"`
	PhoneNumber		string `json:"phone_number"`
	BankName		string `json:"bank_name"`
	AccountNumber	string `json:"account_number"`
	AccountName		string `json:"account_name"`
	Notes 			string `json:"notes"`
	Page			int  `json:"page"`
	Limit			int  `json:"limit"`
}

// TableName ...
func (o *ReqDepositBank) TableName() string {
	return "public.deposit_bank"
}


