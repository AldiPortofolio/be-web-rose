package dbmodels

import "time"

type FeeCicilanSetting struct {
	ID					int 		`json:"id" gorm:"id"`
	AdminFeeDoku        float64   `json:"adminFeeDoku"`
	AdminFeeInfinitium  float64   `json:"adminFeeInfinitium"`
	VaBcaFee         	float64  `json:"vaBcaFee"`
	VaMandiriFee        float64  `json:"vaMandiriFee"`
	VaLainnyaFee        float64  `json:"vaLainnyaFee"`
	User 				string `json:"user"`
	CreatedAt		time.Time 	`json:"createdAt" grom:"created_at"`
	UpdatedAt		time.Time 	`json:"updatedAt" grom:"updated_at"`
}

// TableName ...
func (o *FeeCicilanSetting) TableName() string {
	return "public.fee_cicilan_setting"
}


