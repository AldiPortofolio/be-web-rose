package dto

type ReqMonitoringActivationFDSDto struct {
	Key       string `json:"key" example:"Toko Syandi"`
	StartDate string `json:"startDate" example:"2021-07-01"`
	EndDate   string `json:"endDate" example:"2021-07-15"`
	Limit     int    `json:"limit" example:"10"`
	Page      int    `json:"page" example:"1"`
}