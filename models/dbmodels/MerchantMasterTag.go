package dbmodels

import "time"

type MerchantMasterTag struct {
	ID int64 `json:"id"`
	Mid string `json:"mid"`
	MasterTagCode string `json:"masterTagCode"`
	CreatedAt        time.Time `json:"createdAt"`
	CreatedBy        string    `json:"createdBy"`
	UpdatedAt        time.Time `json:"updatedAt"`
	UpdatedBy        string    `json:"updatedBy"`
}

func (t *MerchantMasterTag)TableName() string {
	return "public.merchant_master_tag"
}
