package dto

type MerchantSettlementConfigDto struct {
	Id                            int64  `json:"id"`
	Mid                           string `json:"mid"`
	NoRekeningToko                string `json:"noRekeningToko"`
	NamaBankTujuanSettlement      string `json:"namaBankTujuanSettlement"`
	NamaPemilikRekening           string `json:"namaPemilikRekening"`
	TipeRekening                  string `json:"tipeRekening"`
	ReportSettlementConfigName    string `json:"reportSettlementConfigName"`
	SettlementExecutionConfigName string `json:"settlementExecutionConfigName"`
	Status                        int64  `json:"status"`
	Email                         string `json:"email"`
	SftpHost                      string `json:"sftpHost"`
	SftpUser                      string `json:"sftpUser"`
	SftpPassword                  string `json:"sftpPassword"`
}
