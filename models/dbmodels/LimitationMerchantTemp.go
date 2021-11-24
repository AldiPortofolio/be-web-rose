package dbmodels

import "time"

type LimitationMerchantTemp struct {
	ID 					  int64       `json:"id" gorm:"id"`
	ActionType            int         `json:"actionType" gorm:"column:action_type"`
	ApprovalStatus        int         `json:"approvalStatus" gorm:"column:approval_status"`
	LatestApproval        time.Time   `json:"latestApproval" gorm:"column:latest_approval"`
	LatestApprover        string      `json:"latestApprover" gorm:"column:latest_approver"`
	LatestSuggestion      time.Time   `json:"latestSuggestion" gorm:"column:latest_suggestion"`
	LatestSuggestor       string      `json:"latestSuggestor" gorm:"column:latest_suggestor"`
	Status                int64       `json:"status" gorm:"column:status"`
	MasterLimitationId    int64 	  `json:"masterLimitationId" gorm:"column:master_limitation_id"`
	ProductType           string      `json:"productType" gorm:"column:product_type"`
	ProductName           string	  `json:"productName" gorm:"column:product_name"`
	LimitFreq             int         `json:"limitFreq" gorm:"column:limit_freq"`
	LimitAmt              int         `json:"limitAmt" gorm:"column:limit_amt"`
	LimitFreqMin          int         `json:"limitFreqMin" gorm:"column:limit_freq_min"`
	LimitAmtMin           int         `json:"limitAmtMin" gorm:"column:limit_amt_min"`
	ByTime                string      `json:"byTime" gorm:"column:by_time"`
	ByGroup				  string `json:"byGroup" gorm:"column:by_group"`
	ActionTypeDesc        string         `json:"actionTypeDesc" gorm:"-"`

}

func (q *LimitationMerchantTemp) TableName() string {
	return "public.master_limitation_temp"
}
