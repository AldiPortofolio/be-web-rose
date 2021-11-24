package dbmodels

import "time"

type SettlementConfig struct {
	ID             int64     `json:"id" gorm:"column:id"`
	BankCode       string    `json:"bankCode" gorm:"column:bank_code"`
	SettlementType string    `json:"settlementType" gorm:"column:settlement_type"`
	Status         string    `json:"status" gorm:"column:status"`
	CreatedAt      time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt      time.Time `json:"updatedAt" gorm:"column:updated_at"`
	UpdatedBy      string    `json:"updatedBy" gorm:"column:updated_by"`
	
}

func (t *SettlementConfig) TableName() string {
	return "public.host_settlement_fee_config"
}
