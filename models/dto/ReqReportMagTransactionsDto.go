package dto

type ReqReportMagTransactionsDto struct {
	Keyword		string	`json:"keyword"`
	FilterBy	string 	`json:"filter_by"`
	Page		int		`json:"page"`
	Limit		int 	`json:"limit"`
}
