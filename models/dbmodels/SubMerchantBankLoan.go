package dbmodels

import "time"

type SubMerchantBankLoan struct {
	ID            int64     `json:"id"`
	Mid           string    `json:"mid"`
	BankCode      string    `json:"bankCode"`
	LoanProductCode string `json:"loanProductCode"`
	MasterBankLoanId string `json:"masterBankLoanId"`
	LoanId string `json:"loanId"`
	LoanAmount float64 `json:"loanAmount"`
	LoanPaidAmount float64 `json:"loanPaidAmount"`
	LoanTxnDate time.Time `json:"loanTxnDate"`
	LoanEffectiveDate time.Time `json:"loanEffectiveDate"`
	LoanMaturityDate time.Time `json:"loanMaturityDate"`
	LastPayDate time.Time `json:"lastPayDate"`
	LoanStatus string `json:"loanStatus"`
	InvoiceNo string `json:"invoiceNo"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (t *SubMerchantBankLoan)TableName() string {
	return "public.sub_merchant_bank_loan"
}
