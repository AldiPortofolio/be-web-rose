package dto

import "rose-be-go/models/dbmodels"

type MasterLimitation struct {
	dbmodels.LimitationMerchant
	MerchantGroup []dbmodels.Merchant `json:"merchantGroup"`
}

type MasterLimitationReq struct {
	MasterLimitationId int64 `json:"masterLimitationId"`
	ByGroup     string `json:"byGroup"`
	ByGroupFilter string `json:"byGroupFilter"`
	ProductType string `json:"productType"`
	ProductName string `json:"productName"`
	Status      int64 `json:"status"`
	LimitFreq   int `json:"limitFreq"`
	LimitAmt    int `json:"limitAmt"`
	LimitFreqMin   int `json:"limitFreqMin"`
	LimitAmtMin    int `json:"limitAmtMin"`
	ByTime      string `json:"byTime"`
	ActionType      int `json:"actionType"`
	Limit       int `json:"limit"`
	Page        int `json:"page"`
}
