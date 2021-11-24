package dto

// ReqPortalActivation ..
type ReqPortalActivation struct {
	MerchantName      string `json:"merchant_name"`
	OutletName        string `json:"outlet_name"`
	Name              string `json:"name"`
	MerchantGroupName string `json:"merchant_group_name"`
	MerchantGroupId   string `json:"merchant_group_id"`
	MID               string `json:"mid"`
	MPAN              string `json:"mpan"`
	TerminalId        string `json:"terminal_id"`
	Alamat            string `json:"address"`
	Kelurahan         string `json:"kelurahan"`
	Kecamatan         string `json:"kecamatan"`
	KabupatenKota     string `json:"kabupaten_kota"`
	Provinsi          string `json:"provinsi"`
	ProfilePict       string `json:"profile_pict"`
	Email             string `json:"email"`
	OwnerID           int64  `json:"owner_id"`
	OwnerName         string `json:"owner_name"`
	MerchantOutletId  string `json:"merchant_outlet_id"`
	TipeMerchant      string `json:"merchant_type"`
	StorePhoneNumber  string `json:"store_phone_number"`
	Password          string `json:"password"`
	Action            string `json:"action"`
	Category          string `json:"category"`
}

// ReqPortalCallback ..
type ReqPortalCallback struct {
	PortalStatus int    `json:"portal_status"`
	Mid          string `json:"merchant_id"`
	Type         string `json:"type"`
}

// BpActivationReq ..
type BpActivationReq struct {
	Id       int64  `json:"id"`
	Category int    `json:"category"`
	Email    string `json:"email"`
	Action   string `json:"action"`
	Password string `json:"password"`
}
