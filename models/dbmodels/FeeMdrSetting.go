package dbmodels

import "time"

// FeeMdrSetting ..
type FeeMdrSetting struct {
	ID                int64     `json:"id" gorm:"column:id;"`
	MidMerchant       string    `json:"midMerchant" gorm:"column:mid_merchant;"`
	MidBank           string    `json:"midBank" gorm:"column:mid_bank;"`
	SecretID          string    `json:"secretId" gorm:"column:secret_id;"`
	Bank              string    `json:"bank" gorm:"column:bank;"`
	Tenor             string    `json:"tenor" gorm:"column:tenor;"`
	BankMdr           float64   `json:"bankMdr" gorm:"column:bank_mdr;"`
	BankMdrCredit     float64   `json:"bankMdrCredit" gorm:"column:bank_mdr_credit;"`
	PlanId            string    `json:"planId"`
	MerchantMdr       float64   `json:"merchantMdr" gorm:"column:merchant_mdr;"`
	MerchantMdrCredit float64   `json:"merchantMdrCredit" gorm:"column:merchant_mdr_credit;"`
	MerchantFeeType   string    `json:"merchantFeeType" gorm:"column:merchant_fee_type;"`
	MerchantFee       float64   `json:"merchantFee" gorm:"column:merchant_fee;"`
	CustomerFeeType   string    `json:"customerFeeType" gorm:"column:customer_fee_type;"`
	CustomerFee       float64   `json:"customerFee" gorm:"column:customer_fee;"`
	Status            string    `json:"status" gorm:"column:status;"`
	UpdatedAt         time.Time `json:"updatedAt" gorm:"column:updated_at;"`
	UpdatedBy         string    `json:"updatedBy" gorm:"column:updated_by;"`
}

// TableName ..
func (q *FeeMdrSetting) TableName() string {
	return "public.fee_mdr_setting"
}
