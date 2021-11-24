package dto

import "time"

type ResPartnerLinkDto struct {
	ID         int64  `json:"id"`
	Code       string `json:"code"`
	PartnerId  string `json:"partner_id"`
	MerchantId string `json:"merchant_id"`
	CreatedBy  string `json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ReqPartnerLinkDto struct {
	ID         int64  `json:"id"`
	Code       string `json:"code"`
	PartnerId  string `json:"partner_id"`
	MerchantId string `json:"merchant_id"`
	CreatedBy  string `json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
}
