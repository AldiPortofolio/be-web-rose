package dbmodels

type ConfigVaMerchantGroup struct {
	Id                                           int64  `json:"id"`
	MerchantGroupId                              int64  `json:"merchantGroupId"`
	VaBca                                        string `json:"vaBca"`
	VaMandiri                                    string `json:"vaMandiri"`
	VaBri                                        string `json:"vaBri"`
	VaLain                                       string `json:"vaLain"`
	VaBcaCompanyCode                             string `json:"vaBcaCompanyCode"`
	VaMandiriCompanyCode                         string `json:"vaMandiriCompanyCode"`
	VaBriCompanyCode                             string `json:"vaBriCompanyCode"`
	VaLainCompanyCode                            string `json:"vaLainCompanyCode"`
	VaBcaSubCompanyCode                          string `json:"vaBcaSubCompanyCode"`
	VaBriSubCompanyCode                          string `json:"vaBriSubCompanyCode"`
	VaMandiriSubCompanyCode                      string `json:"vaMandiriSubCompanyCode"`
	VaLainSubCompanyCode                         string `json:"vaLainSubCompanyCode"`
	VaTransactionType                            string `json:"vaTransactionType"`
	InquiryUrl                                   string `json:"inquiryUrl"`
	PaymentUrl                                   string `json:"paymentUrl"`
	VaTokenUrl                                   string `json:"vaTokenUrl"`
	VaTokenUser                                  string `json:"vaTokenUser"`
	VaTokenPassword                              string `json:"vaTokenPassword"`
	VaAuthKey                                    string `json:"vaAuthKey"`
	VaAuthType                                   string `json:"vaAuthType"`
	HostType                                     string `json:"hostType"`
	VaBcaUpperLimitAmount                        int64  `json:"vaBcaUpperLimitAmount"`
	VaBcaLowerLimitAmount                        int64  `json:"vaBcaLowerLimitAmount"`
	VaBcaMerchantFeeAboveUpperLimitAmount        int64  `json:"vaBcaMerchantFeeAboveUpperLimitAmount"`
	VaBcaFeeTransactionAboveUpperLimitAmount     int64  `json:"vaBcaFeeTransactionAboveUpperLimitAmount"`
	VaBcaMerchantFeeBelowLowerLimitAmount        int64  `json:"vaBcaMerchantFeeBelowLowerLimitAmount"`
	VaBcaFeeTransactionBelowLowerLimitAmount     int64  `json:"vaBcaFeeTransactionBelowLowerLimitAmount"`
	VaBcaMerchantFeeInBetween                    int64  `json:"vaBcaMerchantFeeInBetween"`
	VaBcaFeeTransactionInBetween                 int64  `json:"vaBcaFeeTransactionInBetween"`
	VaBriUpperLimitAmount                        int64  `json:"vaBriUpperLimitAmount"`
	VaBriLowerLimitAmount                        int64  `json:"vaBriLowerLimitAmount"`
	VaBriMerchantFeeAboveUpperLimitAmount        int64  `json:"vaBriMerchantFeeAboveUpperLimitAmount"`
	VaBriFeeTransactionAboveUpperLimitAmount     int64  `json:"vaBriFeeTransactionAboveUpperLimitAmount"`
	VaBriMerchantFeeBelowLowerLimitAmount        int64  `json:"vaBriMerchantFeeBelowLowerLimitAmount"`
	VaBriFeeTransactionBelowLowerLimitAmount     int64  `json:"vaBriFeeTransactionBelowLowerLimitAmount"`
	VaBriMerchantFeeInBetween                    int64  `json:"vaBriMerchantFeeInBetween"`
	VaBriFeeTransactionInBetween                 int64  `json:"vaBriFeeTransactionInBetween"`
	VaMandiriUpperLimitAmount                    int64  `json:"vaMandiriUpperLimitAmount"`
	VaMandiriLowerLimitAmount                    int64  `json:"vaMandiriLowerLimitAmount"`
	VaMandiriMerchantFeeAboveUpperLimitAmount    int64  `json:"vaMandiriMerchantFeeAboveUpperLimitAmount"`
	VaMandiriFeeTransactionAboveUpperLimitAmount int64  `json:"vaMandiriFeeTransactionAboveUpperLimitAmount"`
	VaMandiriMerchantFeeBelowLowerLimitAmount    int64  `json:"vaMandiriMerchantFeeBelowLowerLimitAmount"`
	VaMandiriFeeTransactionBelowLowerLimitAmount int64  `json:"vaMandiriFeeTransactionBelowLowerLimitAmount"`
	VaMandiriMerchantFeeInBetween                int64  `json:"vaMandiriMerchantFeeInBetween"`
	VaMandiriFeeTransactionInBetween             int64  `json:"vaMandiriFeeTransactionInBetween"`
	VaLainUpperLimitAmount                       int64  `json:"vaLainUpperLimitAmount"`
	VaLainLowerLimitAmount                       int64  `json:"vaLainLowerLimitAmount"`
	VaLainMerchantFeeAboveUpperLimitAmount       int64  `json:"vaLainMerchantFeeAboveUpperLimitAmount"`
	VaLainFeeTransactionAboveUpperLimitAmount    int64  `json:"vaLainFeeTransactionAboveUpperLimitAmount"`
	VaLainMerchantFeeBelowLowerLimitAmount       int64  `json:"vaLainMerchantFeeBelowLowerLimitAmount"`
	VaLainFeeTransactionBelowLowerLimitAmount    int64  `json:"vaLainFeeTransactionBelowLowerLimitAmount"`
	VaLainMerchantFeeInBetween                   int64  `json:"vaLainMerchantFeeInBetween"`
	VaLainFeeTransactionInBetween                int64  `json:"vaLainFeeTransactionInBetween"`
}

func (q *ConfigVaMerchantGroup) TableName() string {
	return "public.config_va_merchant_group"
}
