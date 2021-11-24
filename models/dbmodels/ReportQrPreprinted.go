package dbmodels

import "time"

type ReportQrPreprinted struct {
	ID 				int64 `json:"id"`
	StartDate		string		`json:"startDate"`
	EndDate			string		`json:"endDate"`
	FilePath 		string		`json:"filePath"`
	FilePathErr		string		`json:"filePathErr"`
	User 			string 		`json:"user"`
	Status			string		`json:"status"`
	TransactionDate 	time.Time	`json:"transactionDate"`
}

func (t *ReportQrPreprinted) TableName() string  {
	return "report_qr_preprinted"
}