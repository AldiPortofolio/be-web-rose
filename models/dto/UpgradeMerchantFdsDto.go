package dto

type ReqUpgradeMerchantFdsDto struct {
	CustomerPhone      string `json:"customerPhone"`
	IDCard             string `json:"idCard"`
	MerchantAddress    string `json:"merchantAddress"`
	MerchantCity       string `json:"merchantCity"`
	MerchantPostalCode string `json:"merchantPostalCode"`
	BirthPlace         string `json:"birthPlace"`
	MerchantBod        string `json:"merchantBod"`
}
