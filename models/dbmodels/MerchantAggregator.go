package dbmodels

import "time"

type MerchantAggregator struct {
	ID                    int64       `json:"id" gorm:"column:id"`
	ActionType            int         `json:"actionType" gorm:"column:action_type"`
	ApprovalStatus        int         `json:"approvalStatus" gorm:"column:approval_status"`
	LatestApproval        time.Time      `json:"latestApproval" gorm:"column:latest_approval"`
	LatestApprover        string      `json:"latestApprover" gorm:"column:latest_approver"`
	LatestSuggestion      time.Time      `json:"latestSuggestion" gorm:"column:latest_suggestion"`
	LatestSuggestor       string      `json:"latestSuggestor" gorm:"column:latest_suggestor"`
	Status                int64       `json:"status" gorm:"column:status"`
	Version 			  int 			`json:"version" gorm:"column:version"`
	Name 			      string 			`json:"name" gorm:"column:name"`
	Mpan					string 	`json:"mpan" gorm:"column:mpan"`
	Description					string 	`json:"description" gorm:"column:description"`
	Mid					string 	`json:"mid" gorm:"column:mid"`
	Nmid					string 	`json:"nmid" gorm:"column:nmid"`
	Mcc					string 	`json:"mcc" gorm:"column:mcc"`
	Alamat					string 	`json:"alamat" gorm:"column:alamat"`
	KodePos					string 	`json:"kodePos" gorm:"column:kode_pos"`
	Kota					string 	`json:"kota" gorm:"column:kota"`
	MerchantCriteria					string 	`json:"merchantCriteria" gorm:"column:merchant_criteria"`
	MerchantType					string 	`json:"merchantType" gorm:"column:merchant_type"`
	Npwp					string 	`json:"npwp" gorm:"column:npwp"`
	Ktp					string 	`json:"ktp" gorm:"column:ktp"`

}

func (q *MerchantAggregator) TableName() string {
	return "public.merchant_aggregator"
}