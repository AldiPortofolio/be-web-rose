package dbmodels

import "time"

// MagTransactions ...
type MagTransactions struct {
	Id					int			`json:"id" gorm:"id"`
	User				string		`json:"user" gorm:"user"`
	Partner				string		`json:"partner" gorm:"partner"`
	BillingId			string		`json:"billing_id" gorm:"billing_id"`
	Channel				string		`json:"channel" gorm:"channel"`
	MerchantId			string		`json:"merchant_id" gorm:"merchant_id"`
	TerminalId			string		`json:"terminal_id" gorm:"terminal_id"`
	Amount				int			`json:"amount" gorm:"amount"`
	Tip					int			`json:"tip" gorm:"tip"`
	TotalAmount			int			`json:"total_amount" gorm:"total_amount"`
	ReqReferenceNo		string		`json:"req_reference_no" gorm:"req_reference_no"`
	QrCreatedAt			time.Time	`json:"qr_created_at" gorm:"qr_created_at"`
	MerchantPayStatus	string		`json:"merchant_pay_status" gorm:"merchant_pay_status"`
	MerchantPayRef		string		`json:"merchant_pay_ref" gorm:"merchant_pay_ref"`
	MerchantPayTime		time.Time	`json:"merchant_pay_time" gorm:"merchant_pay_time"`
	Issuer				string		`json:"issuer" gorm:"issuer"`
	IssuerCustAccount	string		`json:"issuer_cust_account" gorm:"issuer_cust_account"`
	IssuerRef			string		`json:"issuer_ref" gorm:"issuer_ref"`
	MagBillingId		string		`json:"mag_billing_id" gorm:"mag_billing_id"`
}

// TableName ...
func (q *MagTransactions) TableName() string {
	return "public.mag_transactions"
}

// ExportMagTransactions ...
type ExportMagTransactions struct {
	No					string
	Id					string
	User				string
	Partner				string
	BillingId			string		`csv:"Billing ID"`
	Channel				string
	MerchantId			string		`csv:"Merchant ID"`
	TerminalId			string		`csv:"Terminal ID"`
	Amount				int			`csv:"Amount"`
	Tip					int			`csv:"Tip"`
	TotalAmount			int			`csv:"Total Amount"`
	ReqReferenceNo		string		`csv:"Req Reference Number"`
	QrCreatedAt			string		`csv:"Qr Created At"`
	MerchantPayStatus	string		`csv:"Merchant Pay Status"`
	MerchantPayRef		string		`csv:"Merchant Pay Ref"`
	MerchantPayTime		string		`csv:"Merchant Pay Time"`
	Issuer				string
	IssuerCustAccount	string		`csv:"Issuer Customer Account"`
	IssuerRef			string		`csv:"Issuer Ref"`
	MagBillingId		string		`csv:"MAG Billing ID"`
}
