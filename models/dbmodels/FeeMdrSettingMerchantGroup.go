package dbmodels

import "time"

// FeeMdrSettingMerchantGroup ..
type FeeMdrSettingMerchantGroup struct {
	ID              int64     `json:"id" gorm:"column:id;"`
	IdMerchantGroup int64     `json:"idMerchantGroup" gorm:"column:id_merchant_group;"`
	MidBank         string    `json:"midBank" gorm:"column:mid_bank;"`
	SecretID        string    `json:"secretId" gorm:"column:secret_id;"`
	Bank            string    `json:"bank" gorm:"column:bank;"`
	Tenor           string    `json:"tenor" gorm:"column:tenor;"`
	BankMdr         float64   `json:"bankMdr" gorm:"column:bank_mdr;"`
	MerchantMdr     float64   `json:"merchantMdr" gorm:"column:merchant_mdr;"`
	MerchantFeeType string    `json:"merchantFeeType" gorm:"column:merchant_fee_type;"`
	MerchantFee     float64   `json:"merchantFee" gorm:"column:merchant_fee;"`
	CustomerFeeType string    `json:"customerFeeType" gorm:"column:customer_fee_type;"`
	CustomerFee     float64   `json:"customerFee" gorm:"column:customer_fee;"`
	UpdatedAt       time.Time `json:"updatedAt" gorm:"column:updated_at;"`
	UpdatedBy       string    `json:"updatedBy" gorm:"column:updated_by;"`
}

// TableName ..
func (q *FeeMdrSettingMerchantGroup) TableName() string {
	return "public.fee_mdr_setting_merchant_group"
}
