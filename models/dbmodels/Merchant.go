package dbmodels

type Merchant struct {
	ID                int64  `json:"id" gorm:"column:id"`
	StoreName         string `json:"storeName" gorm:"store_name"`
	MerchantPan       string `json:"merchantPan" gorm:"merchant_pan"`
	MerchantOutletID  string `json:"merchantOutletId" gorm:"merchant_outlet_id"`
	NMid              string `json:"nMid" gorm:"n_mid"`
	ActionType        int    `json:"actionType" gorm:"column:action_type"`
	ApprovalStatus    int    `json:"approvalStatus" gorm:"approval_status"`
	MerchantGroupId   int64  `json:"merchantGroupId" gorm:"column:merchant_group_id"`
	Alamat            string `json:"alamat" gorm:"column:alamat"`
	Kelurahan         string `json:"kelurahan" gorm:"column:kelurahan"`
	Kecamatan         string `json:"kecamatan" gorm:"column:kecamatan"`
	KabupatenKota     string `json:"kabupatenKota" gorm:"column:kabupaten_kota"`
	Provinsi          string `json:"provinsi" gorm:"column:provinsi"`
	SelfiePath        string `json:"selfiePath" gorm:"column:selfie_path"`
	StorePhoneNumber  string `json:"storePhoneNumber" gorm:"column:store_phone_number"`
	PortalStatus      int    `json:"portal_status" gorm:"column:portal_status"`
	OwnerId           int64  `json:"ownerId" gorm:"column:owner_id"`
	MerchantType      string `json:"merchantType" gorm:"column:merchant_type"`
	Category          string `json:"category" gorm:"column:category"`
	TipeBisnis        string `json:"tipeBisnis" gorm:"column:jenis_usaha"`
	SrId              string `json:"srId" gorm:"column:sr_id"`
	Longitude         string `json:"longitude" gorm:"column:longitude"`
	Latitude          string `json:"latitude" gorm:"column:latitude"`
	LokasiBisnis      string `json:"lokasiBisnis" gorm:"column:lokasi_bisnis"`
	JenisLokasiBisnis string `json:"jenisLokasiBisnis" gorm:"column:jenis_lokasi_bisnis"`
	CategoryBisnis    string `json:"categoryBisnis" gorm:"column:kategori_bisnis"`
	OperationHour     string `json:"operationHour" gorm:"column:jam_operasional"`
	BestVisit         string `json:"bestVisit" gorm:"column:best_visit"`
	Patokan           string `json:"patokan" gorm:"column:patokan"`
}

// TableName ...
func (q *Merchant) TableName() string {
	return "public.merchant"
}

// DashboardMerchant ...
type DashboardMerchant struct {
	Id             string `json:"id"`
	StoreName      string `json:"store_name"`
	MerchantGroup  string `json:"merchant_group"`
	MerchantType   string `json:"merchant_type"`
	JenisUsaha     string `json:"jenis_usaha"`
	Mid            string `json:"mid"`
	Mpan           string `json:"mpan"`
	Email          string `json:"email"`
	StatusSuspense bool   `json:"status_suspense"`
	ApprovalStatus string `json:"approval_status"`
	Portal_status  string `json:"portal_status"`
}
