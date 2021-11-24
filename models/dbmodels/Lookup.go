package dbmodels

import "time"

type Lookup struct {
	Id               int64     `json:"id" gorm:"column:id"`
	ActionType       int       `json:"actionType" gorm:"column:action_type"`
	ApprovalStatus   int       `json:"approvalStatus" gorm:"column:approval_status"`
	LatestApproval   time.Time `json:"latestApproval" gorm:"column:latest_approval"`
	LatestApprover   string    `json:"latestApprover" gorm:"column:latest_approver"`
	LatestSuggestion time.Time `json:"latestSuggestion" gorm:"column:latest_suggestion"`
	LatestSuggestor  string    `json:"latestSuggestor" gorm:"column:latest_suggestor"`
	Status           int64     `json:"status" gorm:"column:status"`
	Version          int       `json:"version" gorm:"column:version"`
	Code             string    `json:"code" gorm:"column:code"`
	Descr            string    `json:"descr" gorm:"column:descr"`
	IsAlternateEntry bool      `json:"isAlternateEntry" gorm:"column:is_alternate_entry"`
	IsHighRisk       bool      `json:"isHighRisk" gorm:"column:is_high_risk"`
	LookupGroup      string    `json:"lookupGroup" gorm:"column:lookup_group"`
	Name             string    `json:"name" gorm:"column:name"`
	OrderNo          int64     `json:"orderNo" gorm:"column:order_no"`
	CreatedAt        time.Time `json:"createdAt"`
	CreatedBy        string    `json:"createdBy"`
	UpdatedAt        time.Time `json:"updatedAt"`
	UpdatedBy        string    `json:"updatedBy"`
}

func (q *Lookup) TableName() string {
	return "public.lookup"
}
