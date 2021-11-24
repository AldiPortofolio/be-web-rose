package dbmodels

import "time"

// ReportFinished ...
type ReportReject struct {
	Id              int       `gorm:"column:id" json:"id"`
	StartDate       string    `gorm:"column:start_date" json:"startDate,omitempty"`
	EndDate         string    `gorm:"column:end_date" json:"endDate,omitempty"`
	FilePath        string    `gorm:"column:file_path" json:"filePath"`
	FilePathErr     string    `gorm:"column:file_path_err" json:"filePathErr"`
	User            string    `gorm:"column:user" json:"user"`
	Status          string    `gorm:"column:status" json:"status"`
	TransactionDate time.Time `gorm:"column:transaction_date" json:"transactionDate"`
	Page            int       `json:"page,omitempty"`
	Limit           int       `json:"limit,omitempty"`
}

type ReportRejected struct {
	StoreName string `json:"storeName" gorm:"column:store_name"`
	OwnerName string `json:"ownerName" gorm:"column:owner_name"`
	TransactionDate string `json:"transactionDate" gorm:"column:transaction_date"`
	Username string `json:"username" gorm:"column:username"`
	Reason string `json:"reason" gorm:"column:reason"`
	Status string `json:"status" gorm:"column:status"`
	IdMerchant string `json:"idMerchant" gorm:"column:id_merchant"`
}

type ResultRejected struct {
	OwnerWip OwnerWip `json:"ownerWip"`
	Reason string `json:"reason"`
	StatusRegistration string `json:"statusRegistration"`
}

type OwnerWip struct {
	OwnerFirstName string `json:"ownerFirstName"`
}

// TableName ...
func (f *ReportReject) TableName() string {
	return "public.report_rejected"
}
