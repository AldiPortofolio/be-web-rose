package dbmodels

// ReportQr ...
type ReportQr struct {
	Id              	int       `gorm:"column:id" json:"id"`
	FilePath        	string    `gorm:"column:file_path" json:"filePath"`
	FilePathSuccess     string    `gorm:"column:file_path_success" json:"filePathSuccess"`
	FilePathErr     	string    `gorm:"column:file_path_err" json:"filePathErr"`
	User            	string    `gorm:"column:user" json:"user"`
	Status          	string    `gorm:"column:status" json:"status"`
	TransactionDate 	string 	  `gorm:"column:transaction_date" json:"transactionDate"`
	Page            	int       `json:"page,omitempty" gorm:"-"`
	Limit           	int       `json:"limit,omitempty" gorm:"-"`
}

// TableName ...
func (q *ReportQr) TableName() string {
	return "public.report_qr"
}
