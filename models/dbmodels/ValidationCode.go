package dbmodels

import "time"

type ValidationCode struct {
	ID                   int64     `json:"id"`
	AppID 	string `json:"appId"`
	UserCategoryCode string `json:"userCategoryCode"`
	ValidationCode string `json:"validationCode"`
	ValidFrom time.Time `json:"validFrom"`
	ValidTo time.Time `json:"validTo"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy string `json:"updatedBy"`
}

func (t *ValidationCode) TableName() string {
	return "validation_code_self_register"
}