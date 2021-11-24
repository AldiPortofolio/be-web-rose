package dto

import (
	"time"
)

type ReqBankListDto struct {
	ID                  int64  `json:"id"`
	Code                string `json:"code"`
	FullName            string `json:"fullName"`
	ShortName           string `json:"shortName"`
	SettlementFeeConfig string `json:"settlementFeeConfig"`
	UrlImage            string `json:"urlImage"`
	Status              string `json:"status"`
	Seq                 int32  `json:"seq"`
	Limit               int    `json:"limit"`
	Page                int    `json:"page"`
}

type ResBankList struct {
	ID                  int64     `json:"id"`
	Code                string    `json:"code"`
	FullName            string    `json:"fullName"`
	ShortName           string    `json:"shortName"`
	SettlementFeeConfig string    `json:"settlementFeeConfig" gorm:"column:settlement_type"`
	UrlImage            string    `json:"urlImage"`
	Status              string    `json:"status"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
	CreatedBy           string    `json:"createdBy"`
	UpdatedBy           string    `json:"updatedBy"`
	Seq                 int32     `json:"seq"`
}
