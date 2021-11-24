package dbmodels

import "time"

// MerchantCustomer ..
type MerchantCustomer struct {
	ID          int64     `json:"id" gorm:"column:id;"`
	ServiceID   int64     `json:"serviceId" gorm:"column:service_id;"`
	Name        string    `json:"name" gorm:"column:name;"`
	Phone       string    `json:"phone" gorm:"column:phone;"`
	TypeID      int64     `json:"typeId" gorm:"column:type_id;"`
	CitizenIdNo string    `json:"citizenIdNo" gorm:"column:citizen_id_no;"`
	Merchant    string    `json:"merchant" gorm:"column:merchant;"`
	Pob         string    `json:"pob" gorm:"column:pob;"`
	Dob         time.Time `json:"dob" gorm:"column:dob;"`
	Gender      string    `json:"gender" gorm:"column:gender;"`
	Address     string    `json:"address" gorm:"column:address;"`
	PostalCode  string    `json:"postalCode" gorm:"column:postal_code;"`
	ProvinceID  int64     `json:"provinceId" gorm:"column:province_id;"`
	CityID      int64     `json:"cityId" gorm:"column:city_id;"`
	DistrictID  int64     `json:"districtId" gorm:"column:district_id;"`
	VillageID   int64     `json:"villageId" gorm:"column:village_id;"`
	RW          string    `json:"rw" gorm:"column:rw;"`
	RT          string    `json:"rt" gorm:"column:rt;"`
	Occupation  string    `json:"occupation" gorm:"column:occupation;"`
	Longitude   float64   `json:"longitude" gorm:"column:longitude;"`
	Latitude    float64   `json:"latitude" gorm:"column:latitude;"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"updated_at"`
}

// TableName ..
func (q *MerchantCustomer) TableName() string {
	return "public.merchant_customer"
}

// CustomMerchantCustomer ..
type CustomMerchantCustomer struct {
	ID           int64     `json:"id" gorm:"column:id;"`
	ServiceID    int64     `json:"serviceId" gorm:"column:service_id;"`
	ServiceName  string    `json:"serviceName" gorm:"column:service_name;"`
	Name         string    `json:"name" gorm:"column:name;"`
	Phone        string    `json:"phone" gorm:"column:phone;"`
	TypeID       int64     `json:"typeId" gorm:"column:type_id;"`
	TypeName     string    `json:"typeName" gorm:"column:type_name;"`
	CitizenIdNo  string    `json:"citizenIdNo" gorm:"column:citizen_id_no;"`
	Merchant     string    `json:"merchant" gorm:"column:merchant;"`
	Pob          string    `json:"pob" gorm:"column:pob;"`
	Dob          time.Time `json:"dob" gorm:"column:dob;"`
	Gender       string    `json:"gender" gorm:"column:gender;"`
	Address      string    `json:"address" gorm:"column:address;"`
	PostalCode   string    `json:"postalCode" gorm:"column:postal_code;"`
	ProvinceID   int64     `json:"provinceId" gorm:"column:province_id;"`
	ProvinceName string    `json:"provinceName" gorm:"column:province_name;"`
	CityID       int64     `json:"cityId" gorm:"column:city_id;"`
	CityName     string    `json:"cityName" gorm:"column:city_name;"`
	DistrictID   int64     `json:"districtId" gorm:"column:district_id;"`
	DistrictName string    `json:"districtName" gorm:"column:district_name;"`
	VillageID    int64     `json:"villageId" gorm:"column:village_id;"`
	VillageName  string    `json:"villageName" gorm:"column:village_name;"`
	RW           string    `json:"rw" gorm:"column:rw;"`
	RT           string    `json:"rt" gorm:"column:rt;"`
	Occupation   string    `json:"occupation" gorm:"column:occupation;"`
	Longitude    float64   `json:"longitude" gorm:"column:longitude;"`
	Latitude     float64   `json:"latitude" gorm:"column:latitude;"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"updated_at"`
}

// ExportCustomMerchants ...
type ExportCustomMerchants struct {
	Service				string		`csv:"Service"`
	Name				string		`csv:"Nama"`
	PhoneNumber			string		`csv:"No handphone"`
	Type				string		`csv:"Tipe"`
	CitizenIdNo			string		`csv:"No KTP"`
	Merchant			string		`csv:"Merchant"`
	Dob					string		`csv:"Tempat dan tanggal lahir"`
	Gender				string		`csv:"Jenis kelamin"`
	Address				string		`csv:"Alamat sesuai KTP"`
	Province			string		`csv:"Provinsi sesuai KTP"`
	City				string		`csv:"Kota/kabupaten sesuai KTP"`
	District			string		`csv:"Kecamatan sesuai KTP"`
	Village				string		`csv:"Kelurahan sesuai KTP"`
	VillageId			string		`csv:"Kelurahan ID"`
	RTRW				string		`csv:"RT/RW sesuai KTP"`
	Occupation			string		`csv:"Pekerjaan"`
	Longitude			string		`csv:"Longitude"`
	Latitude			string		`csv:"Latitude"`
}