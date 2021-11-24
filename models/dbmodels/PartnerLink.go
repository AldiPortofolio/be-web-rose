package dbmodels

import "time"

type PartnerLink struct {
	ID         int64  `json:"id" gorm:"id"`
	Code       string `json:"code" gorm:"code"`
	PartnerId  string `json:"partner_id" gorm:"partner_id"`
	MerchantId string `json:"merchant_id" gorm:"merchant_id"`
	CreatedBy  string `json:"created_by" gorm:"created_by"`
	CreatedAt  time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"updated_at"`
}

func (q *PartnerLink) TableName() string {
	return "public.partner_link"
}
