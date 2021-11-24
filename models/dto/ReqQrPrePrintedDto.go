package dto

type ReqQrPrePrintedDto struct {
	StartDate       string    	`json:"startDate"`
	EndDate         string    	`json:"endDate"`
	Page            int       	`json:"page"`
	Limit           int       	`json:"limit"`
} 