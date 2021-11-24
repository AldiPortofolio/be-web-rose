package dbmodels

type LevelMerchant struct {
	ID 						int64  `json:"id" gorm:"id"`
	CodeId 					string  `json:"code_id"`
	Code 					string  `json:"code"`
	Name 					string  `json:"name"`
	Notes       			string 	`json:"notes"`
	Seq       				string 	`json:"seq"`
}

func (q *LevelMerchant) TableName() string {
	return "public.level_merchant"
}
