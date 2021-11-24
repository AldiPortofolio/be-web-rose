package dbmodels

import "time"

// MerchantQrisStatus
type MerchantQrisStatus struct {
	ID                   int64     `json:"id"`
	Mid     string    `json:"mid" gorm:"column:merchant_outlet_id"`
	StoreName             string      `json:"storeName" gorm:"store_name"`
	AgentID              string    `json:"agentId" gorm:"column:agent_id"`
	QrisStatus 			 int32		`json:"qrisStatus" gorm:"column:qris_status"`
	QrisStatusDesc 	  	 string		`json:"qrisStatusDesc" gorm:"column:qris_status_desc"`
	QrisRequestDate    	 time.Time `json:"qrisRequestDate" gorm:"column:qris_request_date"`
	QrisInstallDate     time.Time `json:"qrisInstallDate" gorm:"column:qris_install_date"`

}

func (t *MerchantQrisStatus) TableName() string {
	return "merchant"
}
