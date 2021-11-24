package dto

import "time"

type ReqValidationCodeDto struct {
	ID int64 `json:"id"`
	AppID 	string `json:"appId"`
	UserCategoryCode string `json:"userCategoryCode"`
	ValidationCode string `json:"validationCode"`
	ValidFrom string `json:"validFrom"`
	ValidTo string `json:"validTo"`
	Limit       		  int `json:"limit"`
	Page        		  int `json:"page"`
}

type ResValidationCodeDto struct {
	ID                   int64     `json:"id"`
	AppID 	string `json:"appId"`
	AppName				string `json:"appName"`
	UserCategoryCode string `json:"userCategoryCode"`
	ValidationCode string `json:"validationCode"`
	ValidFrom time.Time `json:"validFrom"`
	ValidTo time.Time `json:"validTo"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy string `json:"updatedBy"`
}