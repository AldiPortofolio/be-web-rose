package dto

type ReqLoanProductMaintenanceDto struct {
	BankCode string `json:"bankCode"`
	Status string `json:"status"`
	LoanProductCode string `json:"loanProductCode"`
	LoanProductName string `json:"loanProductName"`
	Limit    int    `json:"limit" example:"10"`
	Page     int    `json:"page" example:"1"`
}

type ReqSaveLoanProductMaintenanceDto struct {
	ID            int64     `json:"id"`
	LoanProductCode string `json:"loanProductCode"`
	LoanProductName string `json:"loanProductName"`
	BankCode      string    `json:"bankCode"`
	AdminFeeType  string    `json:"adminFeeType"`
	AdminFeeValue float64   `json:"adminFeeValue"`
	Status        string    `json:"status"`
	Description string `json:"description"`

}