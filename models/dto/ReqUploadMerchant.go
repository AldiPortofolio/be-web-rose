package dto

type ReqUploadMerchant struct {
	StartDate       string    	`json:"startDate"`
	EndDate         string    	`json:"endDate"`
	Page            int       	`json:"page"`
	Limit           int       	`json:"limit"`
} 