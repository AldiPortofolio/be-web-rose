package dbmodels

import "time"

type UploadFeeMdrSetting struct {
	ID         int64  `json:"id"`
	Date       time.Time `json:"date"`
	FilePath	string `json:"filePath"`
	FilePathErr	string `json:"filePathErr"`
	Status 		string `json:"status"`
	Notes		string `json:"notes"`
	User 		string `json:"user"`
}


func (t *UploadFeeMdrSetting) TableName() string {
	return "upload_fee_mdr_setting"
}
