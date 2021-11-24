package dbmodels

type AkuisisiSfaFailed struct {
	ID                   int64  `json:"id"`
	StoreName            string `json:"store_name"`
	StorePhoneNumber     string `json:"store_phone_number"`
	MerchantOutletId     string `json:"merchant_outlet_id"`
	MerchantPan          string `json:"merchant_pan"`
	Request              string `json:"request"`
	Response             string `json:"response"`
	UpdatedAt            string `json:"updated_at"`
}

func (t *AkuisisiSfaFailed) TableName() string {
	return "akuisisi_sfa_failed"
}
