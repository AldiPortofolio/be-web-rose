package dbmodels

type WipQueue struct {
	Id 			int64 `json:"id" gorm:"id"`
	TransactionTime		string `json:"transactionTime" gorm:"transaction_time"`
	Key 				string `json:"key" gorm:"key"`
	StoreName			string `json:"storeName" gorm:"store_name"`
	WipId 				string `json:"wipId" gorm:"wip_id"`
	StatusRegistration 	string `json:"statusRegistration" gorm:"status_registration"`
	UserName 			string `json:"userName" gorm:"user_name"`
	Value 				string `json:"value" gorm:"value"`
}

func (t *WipQueue) TableName() string {
	return "public.wip_queue"
}

