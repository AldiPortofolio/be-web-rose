package dto

type ReqMerchantBankLoanDto struct {
	StoreName   string `json:"storeName" example:"Toko Fulan"`
	PhoneNumber string `json:"phoneNumber" example:"0823131231"`
	Status        string    `json:"status"`
	Mid         string `json:"mid"`
	Limit       int    `json:"limit" example:"10"`
	Page        int    `json:"page" example:"1"`
}

type ReqSubMerchantBankloanDto struct {
	MasterBankLoanId string `json:"masterBankLoanId"`

}