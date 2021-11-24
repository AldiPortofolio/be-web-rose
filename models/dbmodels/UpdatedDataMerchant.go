package dbmodels

import "time"

// UpdatedDataMerchant ...
type UpdatedDataMerchant struct {
	ID                int64     `json:"id"`
	Mid               string    `json:"mid"`
	LoanBankCode      string    `json:"loanBankCode"`
	StoreName         string    `json:"storeName"`
	ExpireDate        time.Time `json:"expireDate"`
	Address           string    `json:"address"`
	PostalCode        string    `json:"postalCode"`
	Province          string    `json:"province"`
	City              string    `json:"city"`
	District          string    `json:"district"`
	Village           string    `json:"village"`
	Rt                string    `json:"rt"`
	Rw                string    `json:"rw"`
	LoanBankAccount   string    `json:"loanBankAccount"`
	MerchantGroupId   int64     `json:"merchantGroupId"`
	UserCategoryCode  string    `json:"userCategoryCode"`
	PartnerCode       string    `json:"partnerCode"`
	PartnerCustomerId string    `json:"partnerCustomerId"`

	StorePhoneNumber  string    `json:"storePhoneNumber"`
	TipeBisnis        string    `json:"tipeBisnis" example:""`
	SrId              string    `json:"srId" example:"1"`
	Longitude         string    `json:"longitude" example:"1.98765"`
	Latitude          string    `json:"latitude" example:"12.12345"`
	LokasiBisnis      string    `json:"lokasiBisnis" example:"Pasar"`
	JenisLokasiBisnis string    `json:"jenisLokasiBisnis" example:"4"`
	CategoryBisnis    string    `json:"categoryBisnis" example:"14"`
	OperationHour     string    `json:"operationHour" example:"08:00 - 17:00"`
	BestVisit         string    `json:"bestVisit" example:"15:00 - 17:00"`
	UserCategory      string    `json:"userCategory" example:"idm"`
	Patokan           string    `json:"patokan" example:"ATM BRI"`
	OwnerName         string    `json:"ownerName" example:"Joni"`
	OwnerPhoneNumber  string    `json:"ownerPhoneNumber" example:"08123456789"`
	OwnerAddress      string    `json:"ownerAddress" example:"Jl. Merdaka"`
	OwnerNoId         string    `json:"ownerNoId" example:"11111111111111"`
	OwnerTanggalLahir time.Time `json:"ownerTanggalLahir" example:"2021-10-11"`
	OwnerProvinsi     string    `json:"ownerProvinsi" example:"13"`
	OwnerCity         string    `json:"ownerCity" example:"1303"`
	OwnerKecamatan    string    `json:"ownerKecamatan" example:"1303051"`
	OwnerKelurahan    string    `json:"ownerKelurahan" example:"1303051002"`
	OwnerPostalCode   string    `json:"ownerPostalCode" example:"10102"`

	StoreNameBefore         string    `json:"storeNameBefore"`
	ExpireDateBefore        time.Time `json:"expireDateBefore"`
	AddressBefore           string    `json:"addressBefore"`
	PostalCodeBefore        string    `json:"postalCodeBefore"`
	ProvinceBefore          string    `json:"provinceBefore"`
	CityBefore              string    `json:"cityBefore"`
	DistrictBefore          string    `json:"districtBefore"`
	VillageBefore           string    `json:"villageBefore"`
	RtBefore                string    `json:"rtBefore"`
	RwBefore                string    `json:"rwBefore"`
	LoanBankAccountBefore   string    `json:"loanBankAccountBefore"`
	MerchantGroupIdBefore   int64     `json:"merchantGroupIdBefore"`
	UserCategoryCodeBefore  string    `json:"userCategoryCodeBefore"`
	PartnerCodeBefore       string    `json:"partnerCodeBefore"`
	PartnerCustomerIdBefore string    `json:"partnerCustomerIdBefore"`

	StorePhoneNumberBefore  string    `json:"storePhoneNumberBefore"`
	TipeBisnisBefore        string    `json:"tipeBisnisBefore" example:""`
	SrIdBefore              string    `json:"srIdBefore" example:"1"`
	LongitudeBefore         string    `json:"longitudeBefore" example:"1.98765"`
	LatitudeBefore          string    `json:"latitudeBefore" example:"12.12345"`
	LokasiBisnisBefore      string    `json:"lokasiBisnisBefore" example:"Pasar"`
	JenisLokasiBisnisBefore string    `json:"jenisLokasiBisnisBefore" example:"4"`
	CategoryBisnisBefore    string    `json:"categoryBisnisBefore" example:"14"`
	OperationHourBefore     string    `json:"operationHourBefore" example:"08:00 - 17:00"`
	BestVisitBefore         string    `json:"bestVisitBefore" example:"15:00 - 17:00"`
	UserCategoryBefore      string    `json:"userCategoryBefore" example:"idm"`
	PatokanBefore           string    `json:"patokanBefore" example:"ATM BRI"`
	OwnerNameBefore         string    `json:"ownerNameBefore" example:"Joni"`
	OwnerPhoneNumberBefore  string    `json:"ownerPhoneNumberBefore" example:"08123456789"`
	OwnerAddressBefore      string    `json:"ownerAddressBefore" example:"Jl. Merdaka"`
	OwnerNoIdBefore         string    `json:"ownerNoIdBefore" example:"11111111111111"`
	OwnerTanggalLahirBefore time.Time `json:"ownerTanggalLahirBefore" example:"2021-10-11"`
	OwnerProvinsiBefore     string    `json:"ownerProvinsiBefore" example:"13"`
	OwnerCityBefore         string    `json:"ownerCityBefore" example:"1303"`
	OwnerKecamatanBefore    string    `json:"ownerKecamatanBefore" example:"1303051"`
	OwnerKelurahanBefore    string    `json:"ownerKelurahanBefore" example:"1303051002"`
	OwnerPostalCodeBefore   string    `json:"ownerPostalCodeBefore" example:"10102"`

	Status    string    `json:"status"`
	UpdatedBy string    `json:"updatedBy"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (t *UpdatedDataMerchant) TableName() string {
	return "public.updated_data_merchant"
}
