package dbmodels

import "time"

type MerchantBankLoan struct {
	ID            int64     `json:"id"`
	Mid           string    `json:"mid"`
	BankCode      string    `json:"bankCode"`
	LoanProductCode string `json:"loanProductCode"`
	AccountNumber string    `json:"accountNumber"`
	StoreName     string    `json:"storeName"`
	PhoneNumber   string    `json:"phoneNumber"`
	Tenor         int32     `json:"tenor"`
	Limit         int64     `json:"limit"`
	ExpireDate    time.Time `json:"expireDate"`
	UploadedBy    string    `json:"uploadedBy"`
	UploadedAt    time.Time `json:"uploadedAt"`
	Status        string    `json:"status"`
	LoanProductName string `json:"loanProductName"`
	MasterBankLoanId	string `json:"masterBankLoanId"`


}

func (t *MerchantBankLoan)TableName() string {
	return "public.merchant_bank_loan"
}
