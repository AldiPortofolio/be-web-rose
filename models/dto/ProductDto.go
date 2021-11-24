package dto

type ReqProductDto struct {
	ID 					int64  `json:"id" gorm:"id"`
	Code				string `json:"code"`
	Name 				string `json:"name"`
	Title 				string `json:"title"`
	Desc 				string `json:"desc"`
	Notes 				string `json:"notes"`
	Seq	 				int    `json:"seq"`
	Limit       		  int `json:"limit"`
	Page        		  int `json:"page"`
}

