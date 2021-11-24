package dto

import "time"

// ReqFeeMdrSettingDto ..
type ReqFeeMdrSettingDto struct {
	ID              int64   `json:"id"`
	StoreName       string  `json:"storeName"`
	MidMerchant     string  `json:"midMerchant"`
	MidBank         string  `json:"midBank"`
	SecretID        string  `json:"secretId"`
	Bank            string  `json:"bank"`
	Tenor           string  `json:"tenor"`
	BankMdr         float64 `json:"bankMdr"`
	PlanId          string  `json:"planId"`
	MerchantMdr     float64 `json:"merchantMdr"`
	MerchantFeeType string  `json:"merchantFeeType"`
	MerchantFee     float64 `json:"merchantFee"`
	CustomerFeeType string  `json:"customerFeeType"`
	CustomerFee     float64 `json:"customerFee"`
	BankMdrCredit     float64   `json:"bankMdrCredit"`
	MerchantMdrCredit float64   `json:"merchantMdrCredit"`
	Status          string  `json:"status"`
	Limit           int     `json:"limit"`
	Page            int     `json:"page"`
}

type ResFeeMdrSettingDto struct {
	ID                int64     `json:"id" gorm:"column:id;"`
	StoreName         string    `json:"storeName"`
	MidMerchant       string    `json:"midMerchant" gorm:"column:mid_merchant;"`
	MidBank           string    `json:"midBank" gorm:"column:mid_bank;"`
	SecretID          string    `json:"secretId" gorm:"column:secret_id;"`
	Bank              string    `json:"bank" gorm:"column:bank;"`
	Tenor             string    `json:"tenor" gorm:"column:tenor;"`
	BankMdr           float64   `json:"bankMdr" gorm:"column:bank_mdr;"`
	PlanId            string    `json:"planId"`
	MerchantMdr       float64   `json:"merchantMdr" gorm:"column:merchant_mdr;"`
	MerchantFeeType   string    `json:"merchantFeeType" gorm:"column:merchant_fee_type;"`
	MerchantFee       float64   `json:"merchantFee" gorm:"column:merchant_fee;"`
	CustomerFeeType   string    `json:"customerFeeType" gorm:"column:customer_fee_type;"`
	CustomerFee       float64   `json:"customerFee" gorm:"column:customer_fee;"`
	Status            string    `json:"status" gorm:"column:status;"`
	UpdatedAt         time.Time `json:"updatedAt" gorm:"column:updated_at;"`
	UpdatedBy         string    `json:"updatedBy" gorm:"column:updated_by;"`
	BankMdrCredit     float64   `json:"bankMdrCredit" gorm:"column:bank_mdr_credit;"`
	MerchantMdrCredit float64   `json:"merchantMdrCredit" gorm:"column:merchant_mdr_credit;"`
}
