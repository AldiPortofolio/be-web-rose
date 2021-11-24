package dto

type ReqReportQrPreprintedDto struct {
	Topic           string `json:"topic"`
	StartDate 		string 	`json:"startDate"`
	EndDate			string	`json:"endDate"`
	Mid 			string `json:"mid"`
	StoreName     	string `json:"storeName"`
	User			string	`json:"user"`
}


type ReqGetReportQrPreprintedDto struct {
	Page int `json:"page"`
	Limit int `json:"limit"`
}
