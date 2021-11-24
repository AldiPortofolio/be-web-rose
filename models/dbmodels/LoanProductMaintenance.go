package dbmodels

import "time"

type LoanProductMaintenance struct {
	ID            int64     `json:"id"`
	BankCode      string    `json:"bankCode"`
	BankName      string    `json:"bankName"`
	LoanProductCode string `json:"loanProductCode"`
	LoanProductName string `json:"loanProductName"`
	AdminFeeType  string    `json:"adminFeeType"`
	AdminFeeValue float64   `json:"adminFeeValue"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
	CreatedBy     string    `json:"createdBy"`
	UpdatedAt     time.Time `json:"updatedAt"`
	UpdatedBy     string    `json:"updatedBy"`
	Description string `json:"description"`
}

func (t *LoanProductMaintenance) TableName() string {
	return "public.loan_product_maintenance"
}

