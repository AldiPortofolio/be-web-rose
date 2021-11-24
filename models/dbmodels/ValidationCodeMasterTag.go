package dbmodels

import "time"

type ValidationCodeMasterTag struct {
	ID               int64     `json:"id"`
	ValidationCodeID int64     `json:"validationCodeId" example:"1"`
	MasterTagID      int64     `json:"masterTagId" example:"1"`
	CreatedAt        time.Time `json:"createdAt"`
	CreatedBy        string    `json:"createdBy"`
	UpdatedAt        time.Time `json:"updatedAt"`
	UpdatedBy        string    `json:"updatedBy"`
}

func (t *ValidationCodeMasterTag) TableName() string {
	return "public.validation_code_master_tag"
}
