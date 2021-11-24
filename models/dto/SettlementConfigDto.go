package dto

import "time"

type ReqSettlementConfigDto struct {
	// ID             int64  `json:"id"`
	Code           string `json:"bankCode"`
	SettlementType string `json:"settlementType"`
	Status         string `json:"status"`
	User           string `json:"user"`
	Limit          int    `json:"limit"`
	Page           int    `json:"page"`
}

type ResSettlementConfigDto struct {
	ID             int64     `json:"id"`
	BankCode       string    `json:"bank_code"`
	BankName       string    `json:"bank_name"`
	SettlementType string    `json:"settlement_type"`
	Status         string    `json:"status"`
	UpdatedBy      string    `json:"updated_by"`
	UpdatedAt      time.Time `json:"updated_at"`
}
