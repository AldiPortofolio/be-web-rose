package dbmodels

type UserCategory struct {
	ID 					int64  `json:"id" gorm:"id"`
	Code				string `json:"code"`
	Name 				string `json:"name"`
	Logo				string `json:"logo"`
	AppID				string `json:"appId"`
	Notes 				string `json:"notes"`
	Seq	 				int    `json:"seq"`
}

func (q *UserCategory) TableName() string {
	return "public.user_category"
}