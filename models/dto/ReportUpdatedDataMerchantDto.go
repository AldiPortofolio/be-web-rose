package dto

type ReqReportUpdatedDataMerchantDto struct {
	Topic           string `json:"topic"`
	StartDate 		string 	`json:"startDate"`
	EndDate			string	`json:"endDate"`
	User			string	`json:"user"`
}


type ReqGetReportUpdatedDataMerchantDto struct {
	Page int `json:"page"`
	Limit int `json:"limit"`
}
