package dbmodels

// UploadNmidData ...
type UploadNmidData struct {
	ID              int       	`json:"id" gorm:"PRIMARY KEY; column:id"`
	Date          	string   	`json:"date" gorm:"column:date"`
	FilePath 		string 		`json:"filePath" gorm:"column:file_path"`
	FilePathErr 	string 		`json:"filePathErr" gorm:"column:file_path_err"`
	Status 			string 		`json:"status" gorm:"column:status"`
	User 			string 		`json:"user" gorm:"column:user"`
	TotalUpload int32 `json:"totalUpload"`
	TotalSuccess int32 `json:"totalSuccess"`
	TotalError int32 `json:"totalError"`
	StartDate       string    	`json:"startDate,omitempty" gorm:"-"`
	EndDate         string    	`json:"endDate,omitempty" gorm:"-"`
	Page            int       	`json:"page,omitempty" gorm:"-"`
	Limit           int       	`json:"limit,omitempty" gorm:"-"`
}
// TableName ...
func (o *UploadNmidData) TableName() string {
	return "public.upload_nmid"
}



