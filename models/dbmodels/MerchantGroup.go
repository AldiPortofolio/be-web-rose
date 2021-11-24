package dbmodels

import "time"

// MerchantGroup ..
type MerchantGroup struct {
	ID                        int64     `json:"id" gorm:"column:id;"`
	ActionType                int       `json:"actionType" gorm:"column:action_type;"`
	ApprovalStatus            int       `json:"approvalStatus" gorm:"column:approval_status;"`
	LatestApproval            time.Time `json:"latestApproval" gorm:"column:latest_approval;"`
	LatestApprover            string    `json:"latestApprover" gorm:"column:latest_approver;"`
	LatestSuggestion          time.Time `json:"latestSuggestion" gorm:"column:latest_suggestion;"`
	LatestSuggestor           string    `json:"latestSuggestor" gorm:"column:latest_suggestor;"`
	Status                    int64     `json:"status" gorm:"column:status;"`
	Version                   int       `json:"version" gorm:"column:version;"`
	Alamat                    string    `json:"alamat" gorm:"column:alamat;"`
	EmailPic                  string    `json:"emailPic" gorm:"column:email_pic;"`
	GroupPhoto                string    `json:"groupPhoto" gorm:"column:group_photo;"`
	FkLookupJenisUsaha        string    `json:"fkLookupJenisUsaha" gorm:"column:fk_lookup_jenis_usaha;"`
	FkLookupKabupatenKota     string    `json:"fkLookupKabupatenKota" gorm:"column:fk_lookup_kabupaten_kota;"`
	Kecamatan                 string    `json:"kecamatan" gorm:"column:kecamatan;"`
	Kelurahan                 string    `json:"kelurahan" gorm:"column:kelurahan;"`
	KtpDireksi                string    `json:"ktpDireksi" gorm:"column:ktp_direksi;"`
	MerchantGroupName         string    `json:"merchantGroupName" gorm:"column:merchant_group_name;"`
	NamaPt                    string    `json:"namaPt" gorm:"column:nama_pt;"`
	Negara                    string    `json:"negara" gorm:"column:negara;"`
	NoTelpPic                 string    `json:"noTelpPic" gorm:"column:no_telp_pic;"`
	Npwp                      string    `json:"npwp" gorm:"column:npwp;"`
	NpwpFlag                  int64     `json:"npwpFlag" gorm:"column:npwp_flag;"`
	PicGroup                  string    `json:"picGroup" gorm:"column:pic_group;"`
	Pks                       string    `json:"pks" gorm:"column:pks;"`
	FkLookupProvinsi          string    `json:"fkLookupProvinsi" gorm:"column:fk_lookup_provinsi;"`
	Rt                        string    `json:"rt" gorm:"column:rt;"`
	Rw                        string    `json:"rw" gorm:"column:rw;"`
	Siup                      string    `json:"siup" gorm:"column:siup;"`
	SiupFlag                  int64     `json:"siupFlag" gorm:"column:siup_flag;"`
	FkLookupTipeMerchant      string    `json:"fkLookupTipeMerchant" gorm:"column:fk_lookup_tipe_merchant;"`
	WebsitePerusahaan         string    `json:"websitePerusahaan" gorm:"column:website_perusahaan;"`
	InternalContactPersonID   int64     `json:"internalContactPersonId" gorm:"column:internal_contact_person_id;"`
	MerchantGroupFeeInfoID    int64     `json:"merchantGroupFeeInfoId" gorm:"column:merchant_group_fee_info_id;"`
	MerchantGroupSettleInfoID int64     `json:"merchantGroupSettleInfoId" gorm:"column:merchant_group_settle_info_id;"`
	PostalCode                string    `json:"postalCode" gorm:"column:postal_code;"`
	KtpPenanggungJawab        string    `json:"ktpPenanggungJawab" gorm:"column:ktp_penanggung_jawab;"`
	AktaPendirian             string    `json:"aktaPendirian" gorm:"column:akta_pendirian;"`
	TandaDaftarPerusahaan     string    `json:"tandaDaftarPerusahaan" gorm:"column:tanda_daftar_perusahaan;"`
	PersetujuanMenhumkan      string    `json:"persetujuanMenhumkan" gorm:"column:persetujuan_menhumkan;"`
	StatusSuspense            bool      `json:"statusSuspense" gorm:"column:status_suspense;"`
	EnablePartnerCustomerId   bool 		`json:"enablePartnerCustomerId" gorm:"column:enable_partner_customer_id;"`
	PortalStatus    		  int 		`json:"portalStatus" gorm:"column:portal_status;"`
	EmailPortal   			  string 	`json:"emailPortal" gorm:"column:email_portal;"`
}

// TableName ..
func (q *MerchantGroup) TableName() string {
	return "public.merchant_group"
}
