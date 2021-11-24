package dbmodels

import "time"

type MdrAggregatorTemp struct {
	ID 					  int64       `json:"id" gorm:"id"`
	MdrAggregatorID		  int64		  `json:"mdrAggregatorId" gorm:"column:mdr_aggregator_id"`
	ActionType            int         `json:"actionType" gorm:"column:action_type"`
	ApprovalStatus        int         `json:"approvalStatus" gorm:"column:approval_status"`
	LatestApproval        time.Time   `json:"latestApproval" gorm:"column:latest_approval"`
	LatestApprover        string      `json:"latestApprover" gorm:"column:latest_approver"`
	LatestSuggestion      time.Time   `json:"latestSuggestion" gorm:"column:latest_suggestion"`
	LatestSuggestor       string      `json:"latestSuggestor" gorm:"column:latest_suggestor"`
	Status                int64       `json:"status" gorm:"column:status"`
	GroupPartner		  string	  `json:"groupPartner" gorm:"column:group_partner"`
	MerchantCategory	  string 	  `json:"merchantCategory" gorm:"column:merchant_category"`
	TransactionType 	  string 	  `json:"transactionType" gorm:"column:transaction_type"`
	MdrType				  string 	  `json:"mdrType" gorm:"column:mdr_type"`
	Mdr 			   	  float64	  `json:"mdr" gorm:"column:mdr"`
	Notes 				  string 	  `json:"notes" gorm:"column:notes"`
	MidPartner 			  string 	  `json:"midPartner" gorm:"column:mid_partner"`
	MidMerchant 		  string	  `json:"midMerchant" gorm:"column:mid_merchant"`
	ActionTypeDesc        string         `json:"actionTypeDesc" gorm:"-"`

}


func (q *MdrAggregatorTemp) TableName() string {
	return "public.mdr_aggregator_temp"
}