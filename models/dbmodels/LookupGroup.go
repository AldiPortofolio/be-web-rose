package dbmodels

import "time"

type LookupGroup struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	CreatedAt 			time.Time `json:"createdAt"`
	CreatedBy			string `json:"createdBy"`
	UpdatedAt 			time.Time `json:"updatedAt"`
	UpdatedBy			string `json:"updatedBy"`
}

func (t *LookupGroup) TableName() string {
	return "public.lookup_group"
}