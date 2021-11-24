package dto

type ReqReportCrmPenugasanDto struct {
	Topic           string `json:"topic"`
	StartDate 		string 	`json:"startDate"`
	EndDate			string	`json:"endDate"`
	User			string	`json:"user"`
}


type ReqGetReporCrmPenugasanDto struct {
	Page int `json:"page"`
	Limit int `json:"limit"`
}
