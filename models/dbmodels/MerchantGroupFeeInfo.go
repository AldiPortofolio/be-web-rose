package dbmodels

type MerchantGroupFeeInfo struct {
	Id 						int64 		`json:"id" gorm:"column:id"`
	MdrCreditOffUsValue		float64 	`json:"mdrCreditOffUsValue" gorm:"column:mdr_credit_off_us_value"`
	MdrCreditOnUsValue		float64 	`json:"mdrCreditOnUsValue" gorm:"column:mdr_credit_on_us_value"`
	MdrDebitOffUsValue		float64 	`json:"mdrDebitOffUsValue" gorm:"column:mdr_debit_off_us_value"`
	MdrDebitOnUsValue		float64 	`json:"mdrDebitOnUsValue" gorm:"column:mdr_debit_on_us_value"`
	MdrEmoneyOffUsValue		float64 	`json:"mdrEmoneyOffUsValue" gorm:"column:mdr_emoney_off_us_value"`
	MdrEmoneyOnUsValue		float64 	`json:"mdrEmoneyOnUsValue" gorm:"column:mdr_emoney_on_us_value"`
	FkLookupMdr 			string 		`json:"fkLookupMdr" gorm:"column:fk_lookup_mdr"`
	OtherFee 				float64 	`json:"otherFee" gorm:"column:other_fee"`
	FkLookupProcessingFee 	string 		`json:"fkLookupProcessingFee" gorm:"column:fk_lookup_processing_fee"`
	ProcessingFeeValue 		float64 	`json:"processingFeeValue" gorm:"column:processing_fee_value"`
	RentalEdcFee 			float64 	`json:"rentalEdcFee" gorm:"column:rental_edc_fee"`
	FmsFee 					float64 	`json:"fmsFee" gorm:"column:fms_fee"`
}

// TableName ..
func (q *MerchantGroupFeeInfo) TableName() string {
	return "public.merchant_group_fee_info"
}
