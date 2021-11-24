package dto

import "time"

type ReqLimitTransactionDto struct {
	ID 					int64 		`json:"id"`
	UserCategory		string 		`json:"userCategoryId"`
	LevelMerchant 		string 		`json:"levelMerchant"`
	LimitFreq 			int64 		`json:"limitFreq"`
	MinLimitAmount 		int64 		`json:"minLimitAmount"`
	LimitAmount 		int64 		`json:"limitAmount"`
	TimeFrame 			string 		`json:"timeFrame"`
	FeatureProduct 		string 		`json:"featureProduct"`
	Limit       		int 		`json:"limit"`
	Page        		int 		`json:"page"`
	CreatedAt			time.Time	`json:"createdAt" gorm:"column:created_at;"`
	CreatedBy       	string    	`json:"createdBy" gorm:"column:created_by;"`	
	UpdatedAt       	time.Time 	`json:"updatedAt" gorm:"column:updated_at;"`
	UpdatedBy       	string    	`json:"updatedBy" gorm:"column:updated_by;"`
}