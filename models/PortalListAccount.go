package models

type PortalListMerchantAccount struct {
	Merchant_id 	string 	`gorm:"column:id" json:"mid"`
	Merchant_name	string	`gorm:"column:store_name" json:"merchant_name"`
	Mpan			string	`gorm:"column:merchant_pan" json:"mpan"`
	Alamat			string	`gorm:"column:alamat" json:"address"`
	Kelurahan       string  `gorm:"column:kelurahan" json:"kelurahan"`
	Kecamatan       string  `gorm:"column:kecamatan" json:"kecamatan"`
	Kabupaten_kota   string  `gorm:"column:kabupaten" json:"kabupaten_kota"`
	Provinsi        string  `gorm:"column:provinsi" json:"provinsi"`
	Profile_pict     string  `gorm:"column:profile_pict" json:"profile_pict"`
	Merchant_type	string	`gorm:"column:merchant_type" json:"merchant_type"`
	Merchant_outlet_id	string	`gorm:"column:merchant_outlet_id" json:"merchant_outlet_id"`
	Store_phone		string	`gorm:"column:store_phone_number" json:"store_phone_number"`
	Portal_status	int	`gorm:"column:portal_status" json:"portal_status"`
	Owner_id		int		`gorm:"column:owner_id" json:"owner_id"`
	Email_owner		string	`gorm:"column:owner_email" json:"email"`
	OwnerName string `json:"owner_name" gorm:"owner_name"`
	Merchant_group_id		string	`gorm:"column:merchant_group_id" json:"merchant_group_id"`
	Merchant_group_name		string	`gorm:"column:merchant_group_name" json:"merchant_group_name"`
	PortalCategory 		string `gorm:"column:portal_category" json:"portal_category"`
}

type LogPortal struct {
	User		string	`gorm:"column:user" json:"from"`
	Mid			string	`gorm:"column:mid" json:"mid"`
	From		string	`gorm:"column:from" json:"from"`
	To			string	`gorm:"column:to" json:"to"`
	Action		string	`gorm:"column:action" json:"action"`
	Message		string	`gorm:"column:message" json:"message"`
}

type PortalOutletAccount struct {
	Id int `json:"id" gorm:"column:id"`
	MID string `json:"mid" gorm:"column:mid"`
	MerchantGroupId string `json:"merchant_group_id" gorm:"column:merchant_group_id"`
	MerchantOutletId  string `json:"merchant_outlet_id" gorm:"column:merchant_outlet_id"`
	MerchantPan string `json:"mpan" gorm:"column:merchant_pan"`
	OutletName  string `json:"outlet_name" gorm:"column:outlet_name"`
	MerchantName string `json:"merchant_name" gorm:"column:merchant_name"`
	MerchantGroupName string `json:"merchant_group_name" gorm:"column:merchant_group_name"`
	MerchantType string `json:"merchant_type" gorm:"column:merchant_type"`
	Alamat string `json:"alamat" gorm:"column:alamat"`
	Kelurahan       string  `gorm:"column:kelurahan" json:"kelurahan"`
	Kecamatan       string  `gorm:"column:kecamatan" json:"kecamatan"`
	KabupatenKota   string  `gorm:"column:kabupaten" json:"kabupaten_kota"`
	Provinsi        string  `gorm:"column:provinsi" json:"provinsi"`
	ProfilePict     string  `gorm:"column:profile_pict" json:"profile_pict"`
	StorePhoneNumber string `json:"store_phone_number" gorm:"column:store_phone_number"`
	OwnerName string `json:"owner_name" gorm:"column:owner_name"`
	PortalStatus int `json:"portal_status" gorm:"column:portal_status"`
	EmailOutlet string `json:"email" gorm:"column:email_outlet"`
	TerminalId string `json:"terminal_id" gorm:"column:terminal_id"`
}

// TableName ...
func (o *LogPortal) TableName() string {
	return "public.log_portal"
}