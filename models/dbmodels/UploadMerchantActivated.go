package dbmodels

import "time"

type UploadMerchantActivated struct {
	ID         int64  `json:"id"`
	Date       time.Time `json:"date"`
	FilePath	string `json:"filePath"`
	FilePathErr	string `json:"filePathErr"`
	Status 		string `json:"status"`
	Notes		string `json:"notes"`
	User 		string `json:"user"`
}


func (t *UploadMerchantActivated) TableName() string {
	return "upload_merchant_activated"
}