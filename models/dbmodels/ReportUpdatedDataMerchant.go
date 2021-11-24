package dbmodels

import "time"

type ReportUpdatedDataMerchant struct {
	ID 				int64 `json:"id"`
	StartDate		string		`json:"startDate"`
	EndDate			string		`json:"endDate"`
	FilePath 		string		`json:"filePath"`
	FilePathErr		string		`json:"filePathErr"`
	User 			string 		`json:"user"`
	Status			string		`json:"status"`
	TransactionDate 	time.Time	`json:"transactionDate"`
}

func (t *ReportUpdatedDataMerchant) TableName() string  {
	return "report_updated_data_merchant"
}