package dto

type ReqImageManagementDto struct {
	ID 					int64  `json:"id" gorm:"id"`
	Name	 			string  `json:"name"`
	URL       		  	string 	`json:"url"`
	Notes       		string  `json:"notes"`
	Images				string	`json:"images"`
	Limit       		  int `json:"limit"`
	Page        		  int `json:"page"`
}

