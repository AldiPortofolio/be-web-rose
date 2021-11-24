package dbmodels

type FiturProduct struct {
	ID 					int64  `json:"id" gorm:"id"`
	ProductID			int64  `json:"productId"`
	Code				string `json:"code"`
	Icon			 	string `json:"icon"`
	Name 				string `json:"name"`
	Notes 				string `json:"notes"`
	Seq 				int    `json:"seq"`
	Url 				string `json:"url"`
}

func (q *FiturProduct) TableName() string {
	return "public.fitur_product"
}