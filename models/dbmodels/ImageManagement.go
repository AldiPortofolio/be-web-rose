package dbmodels

type ImageManagement struct {
	ID 					int64  `json:"id" gorm:"id"`
	Name	 			string  `json:"name"`
	URL       		  	string 	`json:"url"`
	Notes       		string  `json:"notes"`
}

func (q *ImageManagement) TableName() string {
	return "public.image_management"
}