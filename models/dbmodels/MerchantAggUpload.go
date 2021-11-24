package dbmodels

//MerchantAggUpload ...
type MerchantAggUpload struct {
	ID              int       	`json:"id" gorm:"PRIMARY KEY; column:id"`
	Date          	string   	`json:"date" gorm:"column:date"`
	FilePath 		string 		`json:"filePath" gorm:"column:file_path"`
	FilePathSuccess string		`json:"filePathSuccess" gorm:"column:file_path_success"`
	FilePathErr 	string 		`json:"filePathErr" gorm:"column:file_path_err"`
	Status 			string 		`json:"status" gorm:"column:status"`
	User 			string 		`json:"user" gorm:"column:user"`
	MidAggregator   string 		`json:"midAggregator" gorm:"column:mid_aggregator"`
	StartDate       string    	`json:"startDate,omitempty" gorm:"-"`
	EndDate         string    	`json:"endDate,omitempty" gorm:"-"`
	Page            int       	`json:"page,omitempty" gorm:"-"`
	Limit           int       	`json:"limit,omitempty" gorm:"-"`
}

//TableName ...
func (o *MerchantAggUpload) TableName() string {
	return "public.merchant_aggregator_upload_temp"
}
