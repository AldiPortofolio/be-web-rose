package dto

type ReqFiturProductDto struct {
	ID 					int64  `json:"id" gorm:"id"`
	ProductID			int64  `json:"productId"`
	Code				string `json:"code"`
	Icon			 	string `json:"icon"`
	Name 				string `json:"name"`
	Notes 				string `json:"notes"`
	Seq 				int    `json:"seq"`
	Url 				string `json:"url"`
	Limit       		  int `json:"limit"`
	Page        		  int `json:"page"`
}

type ResFiturProductDto struct {
	ID 					int64  `json:"id" gorm:"id"`
	ProductID			int64  `json:"productId"`
	ProductName			string `json:"productName" gorm:"product_name"`
	Code				string `json:"code"`
	Icon			 	string `json:"icon"`
	Name 				string `json:"name"`
	Notes 				string `json:"notes"`
	Seq 				int    `json:"seq"`
	Url 				string `json:"url"`
}
