package dto

type ReqReportSendDto struct {
	Topic           string `json:"topic"`
	StartDate 		string 	`json:"startDate"`
	EndDate			string	`json:"endDate"`
	User			string	`json:"user"`
}
