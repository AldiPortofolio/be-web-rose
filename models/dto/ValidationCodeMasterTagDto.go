package dto

import "time"

type ReqValidationCodeMasterTagDto struct {
	ValidationCodeID int64  `json:"validationCodeId" example:"1"`
	MasterTagID      int64  `json:"masterTagId" example:"1"`
	MasterTagCode    string `json:"masterTagCode" example:"001"`
	Limit            int    `json:"limit" example:"10"`
	Page             int    `json:"page" example:"1"`
}

type ResValidationCodeMasterTagDto struct {
	ID               int64     `json:"id"`
	ValidationCodeID int64     `json:"validationCodeId"`
	ValidationCode   string    `json:"validationCode"`
	MasterTagID      int64     `json:"masterTagId"`
	MasterTagCode    string    `json:"masterTagCode"`
	CreatedAt        time.Time `json:"createdAt"`
	CreatedBy        string    `json:"createdBy"`
	UpdatedAt        time.Time `json:"updatedAt"`
	UpdatedBy        string    `json:"updatedBy"`
}

type ReqSaveValidationCodeMasterTagDto struct {
	ID               int64 `json:"id"`
	ValidationCodeID int64 `json:"validationCodeId" example:"1"`
	MasterTagID      int64 `json:"masterTagId" example:"1"`
}