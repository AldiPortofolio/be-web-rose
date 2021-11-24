package dbmodels

import "time"

// ReportExportAkuisisiSfa ...
type ReportExportAkuisisiSfa struct {
	Id              int       `gorm:"column:id" json:"id"`
	Filter       string    `gorm:"column:filter" json:"filter"`
	FilePath        string    `gorm:"column:file_path" json:"filePath"`
	FilePathErr     string    `gorm:"column:file_path_err" json:"filePathErr"`
	User            string    `gorm:"column:user" json:"user"`
	Status          string    `gorm:"column:status" json:"status"`
	TransactionDate time.Time `gorm:"column:transaction_date" json:"transactionDate"`
	StartDate       string    `gorm:"column:start_date" json:"startDate,omitempty"`
	EndDate         string    `gorm:"column:end_date" json:"endDate,omitempty"`
	Page            int       `json:"page,omitempty"`
	Limit           int       `json:"limit,omitempty"`
}


// TableName ...
func (f *ReportExportAkuisisiSfa) TableName() string {
	return "public.report_export_akuisisi_sfa"
}
