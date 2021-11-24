package dto

import "time"

// ReqMerchantCustomerDto ..
type ReqMerchantCustomerDto struct {
	ID          int64     `json:"id"`
	ServiceID   int64     `json:"serviceId"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	TypeID      int64     `json:"typeId"`
	CitizenIdNo string    `json:"citizenIdNo"`
	Merchant    string    `json:"merchant"`
	Pob         string    `json:"pob"`
	Dob         time.Time `json:"dob"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	PostalCode  string    `json:"postalCode"`
	ProvinceID  int64     `json:"provinceId"`
	CityID      int64     `json:"cityId"`
	DistrictID  int64     `json:"districtId"`
	VillageID   int64     `json:"villageId"`
	RW          string    `json:"rw"`
	RT          string    `json:"rt"`
	Occupation  string    `json:"occupation"`
	Longitude   float64   `json:"longitude"`
	Latitude    float64   `json:"latitude"`
	Limit       int       `json:"limit"`
	Page        int       `json:"page"`
}
