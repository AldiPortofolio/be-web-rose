package dbmodels

type MerchantWipStatusList struct {
	Id 				int64	`json:"id" gorm:"id"`
	MerchantWipId 	int64 	`json:"merchantWipId" gorm:"merchant_wip_id"`
	TransactionDate string	`json:"transactionDate" gorm:"transaction_date"`
	RegistrationStatus	string `json:"registrationStatus" gorm:"registration_status"`
	Username			string `json:"username" gorm:"username"`
	Code 				string `json:"code" gorm:"code"`
	Reason 				string `json:"reason" gorm:"reason"`
	Rolename 			string `json:"rolename" gorm:"rolename"`
	Totaltimelasttransinsec	int64 `json:"totaltimelasttransinsec" gorm:"totaltimelasttransinsec"`
	Lastidtrans				int64 `json:"lastidtrans"`
}

// TableName ...
func (T *MerchantWipStatusList) TableName() string {
	return "public.merchant_wip_status_list2"
}
