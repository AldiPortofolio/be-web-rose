package dbmodels

type MerchantGroupInternalContactPerson struct {
	Id  			int64 		`json:"id" gorm:"column:id"`
	BusinessPic   	string 		`json:"businessPic" gorm:"column:business_pic"`
	Notes  			string 		`json:"notes" gorm:"column:notes"`
	SettleOperationPic string 	`json:"settleOperationPic" gorm:"column:settle_operation_pic"`
	TechnicalPic   	string 		`json:"technicalPic" gorm:"column:technical_pic"`
}

// TableName ..
func (q *MerchantGroupInternalContactPerson) TableName() string {
	return "public.merchant_group_internal_contact_person"
}
