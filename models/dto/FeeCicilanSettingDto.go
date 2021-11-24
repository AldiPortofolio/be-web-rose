package dto

type ReqFeeCicilanSettingDto struct {
	AdminFeeDoku        float64   `json:"adminFeeDoku"`
	AdminFeeInfinitium  float64   `json:"adminFeeInfinitium"`
	VaBcaFee         	float64  `json:"vaBcaFee"`
	VaMandiriFee        float64  `json:"vaMandiriFee"`
	VaLainnyaFee        float64  `json:"vaLainnyaFee"`
}
