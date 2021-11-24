package dbmodels

type Product struct {
	ID 					int64  `json:"id" gorm:"id"`
	Code				string `json:"code"`
	Name 				string `json:"name"`
	Title 				string `json:"title"`
	Desc 				string `json:"desc"`
	Notes 				string `json:"notes"`
	Seq	 				int    `json:"seq"`
}

func (q *Product) TableName() string {
	return "public.product"
}