package dbmodels

import "time"

type MerchantSettlementConfig struct {
	Id                            int64     `json:"id"`
	Mid                           string    `json:"mid"`
	NoRekeningToko                string    `json:"noRekeningToko" gorm:"column:no_rekening_toko"`
	NamaBankTujuanSettlement      string    `json:"namaBankTujuanSettlement" gorm:"column:nama_bank_tujuan_settlement"`
	NamaPemilikRekening           string    `json:"namaPemilikRekening" gorm:"column:nama_pemilik_rekening"`
	TipeRekening                  string    `json:"tipeRekening" gorm:"column:tipe_rekening"`
	ReportSettlementConfigName    string    `json:"rptSetConfName" gorm:"column:report_settlement_config_by"`
	SettlementExecutionConfigName string    `json:"setExecConfName" gorm:"column:settlement_execution_config"`
	Status                        int64     `json:"status" gorm:"column:status"`
	Email                         string    `json:"email" gorm:"column:email"`
	SftpHost                      string    `json:"sftpHost" gorm:"column:sftp_host"`
	SftpUser                      string    `json:"sftpUser" gorm:"column:sftp_user"`
	SftpPassword                  string    `json:"sftpPassword" gorm:"column:sftp_password"`
	CreatedAt                     time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt                     time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (t *MerchantSettlementConfig) TableName() string {
	return "merchant_settlement_config"
}
