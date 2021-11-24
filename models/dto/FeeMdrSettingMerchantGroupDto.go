package dto

// ReqFeeMdrSettingMerchantGroupDto ..
type ReqFeeMdrSettingMerchantGroupDto struct {
	ID              int64   `json:"id"`
	IdMerchantGroup int64   `json:"idMerchantGroup"`
	MidBank         string  `json:"midBank"`
	SecretID        string  `json:"secretId"`
	Bank            string  `json:"bank"`
	Tenor           string  `json:"tenor"`
	BankMdr         float64 `json:"bankMdr"`
	MerchantMdr     float64 `json:"merchantMdr"`
	MerchantFeeType string  `json:"merchantFeeType"`
	MerchantFee     float64 `json:"merchantFee"`
	CustomerFeeType string  `json:"customerFeeType"`
	CustomerFee     float64 `json:"customerFee"`
	Limit           int     `json:"limit"`
	Page            int     `json:"page"`
}
