package dbmodels

import "time"

type MonitoringActivationFDS struct {
	ID 					int64 		`json:"id"`
	StoreName			string 		`json:"storeName" gorm:"column:store_name"`
	StorePhoneNumber 	string 		`json:"storePhoneNumber" gorm:"column:store_phone_number"`
	MID					string 		`json:"mid" gorm:"column:merchant_id;"`
	MPAN				string 		`json:"mpan" gorm:"column:merchant_pan;"`
	StatusActivationFds string 		`json:"statusActivationFds" gorm:"column:status;"`
	CreatedAt			time.Time	`json:"createdAt" gorm:"column:created_at;"`
	UpdatedAt       	time.Time 	`json:"updatedAt" gorm:"column:updated_at;"`
}

func (q *MonitoringActivationFDS) TableName() string {
	return "public.monitoring_activation_fds"
}
