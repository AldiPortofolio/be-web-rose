package dbmodels

type AkuisisiSfa struct {
	ID                       int64  `json:"id"`
	PhotoKtp                 string `json:"photo_ktp"`
	PhotoLocation            string `json:"photo_location"`
	PhotoLocation2           string `json:"photo_location_2" gorm:"column:photo_location_2"`
	PhotoSelfie              string `json:"photo_selfie"`
	PhotoSign                string `json:"photo_sign"`
	MerchantGroupName        string `json:"merchant_group_name"`
	StoreName                string `json:"store_name"`
	StoreJenisUsaha          string `json:"store_jenis_usaha"`
	KategoriBisnis           string `json:"kategori_bisnis"`
	StoreAlamat              string `json:"store_alamat"`
	StoreKelurahan           string `json:"store_kelurahan"`
	StoreKecamatan           string `json:"store_kecamatan"`
	StoreJamOperasional      string `json:"store_jam_operasional"`
	StoreJenisLokasiBisnis   string `json:"store_jenis_lokasi_bisnis"`
	StoreKabupatenKota       string `json:"store_kabupaten_kota"`
	StorePostalCode          string `json:"store_postal_code"`
	StoreProvinsi            string `json:"store_provinsi"`
	StoreLatitude            string `json:"store_latitude"`
	StoreLongitude           string `json:"store_longitude"`
	StoreLokasiBisnis        string `json:"store_lokasi_bisnis"`
	StorePhoneNumber         string `json:"store_phone_number"`
	AgentID                  string `json:"agent_id"`
	AgentName                string `json:"agent_name"`
	AgentCompanyID           string `json:"agent_company_id"`
	AgentPhoneNumber         string `json:"agent_phone_number"`
	OwnerAddress             string `json:"owner_address"`
	OwnerFirstname           string `json:"owner_firstname"`
	OwnerJenisKelamin        string `json:"owner_jenis_kelamin"`
	OwnerKabupatenKota       string `json:"owner_kabupaten_kota"`
	OwnerKecamatan           string `json:"owner_kecamatan"`
	OwnerKelurahan           string `json:"owner_kelurahan"`
	OwnerKodePos             string `json:"owner_kode_pos"`
	OwnerLastname            string `json:"owner_lastname"`
	OwnerNamaGadisIbuKandung string `json:"owner_nama_gadis_ibu_kandung"`
	OwnerNoID                string `json:"owner_no_id"`
	OwnerNoTelf              string `json:"owner_no_telf"`
	OwnerNoTelfLainnya       string `json:"owner_no_telf_lainnya"`
	OwnerPekerjaan           string `json:"owner_pekerjaan"`
	OwnerProvinsi            string `json:"owner_provinsi"`
	OwnerRt                  string `json:"owner_rt"`
	OwnerRw                  string `json:"owner_rw"`
	OwnerTanggalLahir        string `json:"owner_tanggal_lahir"`
	OwnerTempatLahir         string `json:"owner_tempat_lahir"`
	OwnerTglExpiredID        string `json:"owner_tgl_expired_id"`
	OwnerTipeID              string `json:"owner_tipe_id"`
	OwnerTitle               string `json:"owner_title"`
	DeviceType               string `json:"device_type"`
	MetodePembayaran         string `json:"metode_pembayaran"`
	DeviceGroup              string `json:"device_group"`
	DeviceBrand              string `json:"device_brand"`
	OutletID                 string `json:"outlet_id"`
	TerminalPhoneNumber      string `json:"terminal_phone_number"`
	TerminalProvider         string `json:"terminal_provider"`
	InstitutionID            string `json:"institution_id"`
	Notes                    string `json:"notes"`
	MerchantOutletId         string `json:"merchant_outlet_id"`
	Nmid                     string `json:"nmid"`
	Level                    string `json:"level"`
	ExistingQrValue          string `json:"existing_qr_value"`
	MerchantPan              string `json:"merchant_pan"`
	Category                 string `json:"category"`
	StoreNamePreprinted      string `json:"store_name_preprinted"`
	PhotoLocationLeft        string `json:"photo_location_left"`
	PhotoLocationRight       string `json:"photo_location_right"`
	PartnerCustomerId        string `json:"partner_customer_id"`
	FotoPreprinted           string `json:"foto_preprinted"`
	PriorityLevel            string `json:"priority_level"`
	Status                   string `json:"status"`
	UpdatedAt                string `json:"updated_at"`
	SrId                     string `json:"srId"`
	PairVerifySimilarity     string `json:"pairVerifySimilarity"`
	Description              string `json:"description"`
}

func (t *AkuisisiSfa) TableName() string {
	return "akuisisi_sfa"
}
