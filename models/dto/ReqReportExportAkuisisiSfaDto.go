package dto

type ReqReportExportAkuisisiSfaDto struct {
	Topic     string `json:"topic"`
	Key       string `json:"key"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	User      string `json:"user"`
}
