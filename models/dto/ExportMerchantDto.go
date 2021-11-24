package dto

type ReqReportExportMerchantDto struct {
	Topic        string `json:"topic"`
	Name         string `json:"name"`
	PortalStatus string `json:"portalStatus"`
	TipeMerchant string `json:"tipeMerchant"`
	User         string `json:"user"`
}
