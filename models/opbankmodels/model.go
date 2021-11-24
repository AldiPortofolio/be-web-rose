package opbankmodels

type ReqInquiry struct {
	AccountNo string `json:"account_no"`
	BankCode string `json:"bank_code"`
}

type ResInquiryInternal struct {
	Rc   string       `json:"rc"`
	Msg  string       `json:"msg"`
	Data DataInternal `json:"data"`
}

type DataInternal struct {
	AccountNo    string `json:"accountNo"`
	AccountName  string `json:"accountName"`
	AccountType  string `json:"accountType"`
	CurrencyType string `json:"currencyType"`
}

type ResInquiryExternal struct {
	Rc   string       `json:"rc"`
	Msg  string       `json:"msg"`
	Data DataExternal `json:"data"`
}

type DataExternal struct {
	DestinationAccountNo   string `json:"destinationAccountNo"`
	DestinationBankName    string `json:"destinationBankName"`
	DestinationAccountName string `json:"destinationAccountName"`
}