package dbmodels

import "time"

type MerchantAggregatorDetailTemp struct {
	ID                    int64       `json:"id" gorm:"column:id"`
	ActionType            int         `json:"actionType" gorm:"column:action_type"`
	ApprovalStatus        int         `json:"approvalStatus" gorm:"column:approval_status"`
	LatestApproval        time.Time      `json:"latestApproval" gorm:"column:latest_approval"`
	LatestApprover        string      `json:"latestApprover" gorm:"column:latest_approver"`
	LatestSuggestion      time.Time      `json:"latestSuggestion" gorm:"column:latest_suggestion"`
	LatestSuggestor       string      `json:"latestSuggestor" gorm:"column:latest_suggestor"`
	Status                int64       `json:"status" gorm:"column:status"`
	MerchantAggregatorDetailId	int64	`json:"merchantAggregatorDetailId" gorm:"column:merchant_aggregator_detail_id"`
	MidAggregator		  string	  `json:"midAggregator" gorm:"column:mid_aggregator"`
	MidMerchant 		  string 	  `json:"midMerchant" gorm:"column:mid_merchant"`
	PartnerName 		  string 	  `json:"partnerName" gorm:"-"`
	MerchantName 		  string 	  `json:"merchantName" gorm:"-"`
	MerchantPan		  	  string      `json:"merchantPan" gorm:"-"`
	MerchantNmid		  string      `json:"merchantNmid" gorm:"-"`
}

func (q *MerchantAggregatorDetailTemp) TableName() string {
	return "public.merchant_aggregator_detail_temp"
}