package dbmodels

import "time"

type UploadMerchantMasterTag struct {
	ID         int64  `json:"id"`
	Date       time.Time `json:"date"`
	FilePath	string `json:"filePath"`
	FilePathErr	string `json:"filePathErr"`
	Status 		string `json:"status"`
	Notes		string `json:"notes"`
	User 		string `json:"user"`
	TotalUpload int32 `json:"totalUpload"`
	TotalSuccess int32 `json:"totalSuccess"`
	TotalError int32 `json:"totalError"`
}


func (t *UploadMerchantMasterTag) TableName() string {
	return "upload_merchant_master_tag"
}
