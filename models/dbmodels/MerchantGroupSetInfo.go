package dbmodels

type MerchantGroupSetInfo struct {
	Id 							int64 		`json:"id" gorm:"column:id"`
	NamaBankTujuanSettlement	string 		`json:"namaBankTujuanSettlement" gorm:"column:nama_bank_tujuan_settlement"`
	NamaPemilikRekening 		string 		`json:"namaPemilikRekening" gorm:"column:nama_pemilik_rekening"`
	NomorRekening 				string 		`json:"nomorRekening" gorm:"column:nomor_rekening"`
	FkLookupReportSettleCfg 	string 		`json:"fkLookupReportSettleCfg" gorm:"column:fk_lookup_report_settle_cfg"`
	SendReportUrl 				string 		`json:"sendEeportUrl" gorm:"column:send_report_url"`
	FkLookupSendReportVia  		string 		`json:"fkLookupSendReportVia" gorm:"column:fk_lookup_send_report_via"`
	FkLookupSettlementExecCfg 	string 		`json:"fkLookupSettlementExecCfg" gorm:"column:fk_lookup_settlement_exec_cfg"`
	TipeRekening 				string 		`json:"tipeRekening" gorm:"column:tipe_rekening"`
}

// TableName ..
func (q *MerchantGroupSetInfo) TableName() string {
	return "public.merchant_group_set_info"
}
