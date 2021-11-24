package dbmodels

import "time"

type Acquititions struct {
	ID               int64     `json:"id"`
	MerchantType     string    `json:"merchantType"`
	MerchantGroupId  int64     `json:"merchantGroupId"`
	MerchantCategory string    `json:"merchantCategory"`
	Name             string    `json:"name"`
	LogoUrl          string    `json:"logoUrl"`
	RegisterUsingId  bool      `json:"registerUsingId"`
	Sequence         int       `json:"sequence"`
	ShowInApp        string    `json:"showInApp"`
	SalesRetails     string    `json:"salesRetails"`
	BusinessType     string    `json:"businessType"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	UpdatedBy        string    `json:"updatedBy"`
}

func (q *Acquititions) TableName() string {
	return "public.acquisitions"
}
