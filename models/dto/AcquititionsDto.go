package dto

import "time"

type ReqAcquititionsDto struct {
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
	Limit            int       `json:"limit"`
	Page             int       `json:"page"`
}

type ResAcquititionsDto struct {
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

type ReqFilterAcquititionsDto struct {
	Limit            int    `json:"limit" example:"10"`
	MerchantCategory string `json:"merchantCategory" example:"Sales"`
	MerchantGroup    int    `json:"merchantGroup" example:"55"`
	BusinessType     string `json:"businessType" example:"UMKM"`
	Name             string `json:"name" example:"test"`
	Page             int    `json:"page" example:"1"`
}
