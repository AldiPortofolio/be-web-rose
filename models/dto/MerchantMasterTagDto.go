package dto

type ReqMerchantMasterTagDto struct {
	Mid           string `json:"mid"`
	MasterTagCode string `json:"masterTagCode"`
	Limit         int    `json:"limit" example:"10"`
	Page          int    `json:"page" example:"1"`
}

type ReqMerchantMasterTagByMidDto struct {
	Mid string `json:"mid"`
}

type ResMerchantMasterTag struct {
	Mid           string `json:"mid"`
	MasterTagCode string `json:"masterTagCode"`
	Name          string `json:"name"`
}
