package dbmodels

type MerchantWip struct {
	ID                    int64       `json:"id" gorm:"column:id"`
	ActionType            int         `json:"actionType" gorm:"action_type"`
	ApprovalStatus        int         `json:"approvalStatus" gorm:"approval_status"`
	LatestApproval        string      `json:"latestApproval" gorm:"latest_approval"`
	LatestApprover        string      `json:"latestApprover" gorm:"latest_approver"`
	LatestSuggestion      string      `json:"latestSuggestion" gorm:"latest_suggestion"`
	LatestSuggestor       string      `json:"latestSuggestor" gorm:"latest_suggestor"`
	Status                int64       `json:"status" gorm:"status"`
	Version               int         `json:"version" gorm:"version"`
	AgentID               string      `json:"agentId" gorm:"agent_id"`
	AgentName             string      `json:"agentName" gorm:"agent_name"`
	Alamat                string      `json:"alamat" gorm:"alamat"`
	APIKey                string      `json:"apiKey" gorm:"api_key"`
	ApprovalNote          string      `json:"approvalNote" gorm:"approval_note"`
	HostStatus            string      `json:"hostStatus" gorm:"host_status"`
	InstitutionID         string      `json:"institutionId" gorm:"institution_id"`
	JamOperasional        string      `json:"jamOperasional" gorm:"jam_operasional"`
	JenisLokasiBisnis     string      `json:"jenisLokasiBisnis" gorm:"jenis_lokasi_bisnis"`
	JenisUsaha            string      `json:"jenisUsaha" gorm:"jenis_usaha"`
	KabupatenKota         string      `json:"kabupatenKota" gorm:"kabupaten_kota"`
	Kecamatan             string      `json:"kecamatan" gorm:"kecamatan"`
	Kelurahan             string      `json:"kelurahan" gorm:"kelurahan"`
	KtpPath               string      `json:"ktpPath" gorm:"ktp_path"`
	Latitude              string      `json:"latitude" gorm:"latitude"`
	LogoPath              string      `json:"logoPath" gorm:"logo_path"`
	LokasiBisnis          string      `json:"lokasiBisnis" gorm:"lokasi_bisnis"`
	Longitude             string      `json:"longitude" gorm:"longitude"`
	MerchantCategoryCode  string      `json:"merchantCategoryCode" gorm:"merchant_category_code"`
	MerchantGroupID       int         `json:"merchantGroupId" gorm:"merchant_group_id"`
	MerchantOutletID      string      `json:"merchantOutletId" gorm:"merchant_outlet_id"`
	MerchantPan           string      `json:"merchantPan" gorm:"merchant_pan"`
	MerchantPhoto2Path    string      `json:"merchantPhoto2Path" gorm:"merchant_photo_2_path"`
	MerchantPhotoPath     string      `json:"merchantPhotoPath" gorm:"merchant_photo_path"`
	MerchantType          string      `json:"merchantType" gorm:"merchant_type"`
	Notes                 string      `json:"notes" gorm:"notes"`
	PostalCode            string      `json:"postalCode" gorm:"postal_code"`
	Provinsi              string      `json:"provinsi" gorm:"provinsi"`
	Reason                string      `json:"reason" gorm:"reason"`
	ReferralCode          string      `json:"referralCode" gorm:"referral_code"`
	SecretID              string      `json:"secretId" gorm:"secret_id"`
	SecretQuestion        string      `json:"secretQuestion" gorm:"secret_question"`
	SecretQuestionAnswer  string      `json:"secretQuestionAnswer" gorm:"secret_question_answer"`
	SelfiePath            string      `json:"selfiePath" gorm:"selfie_path"`
	SignPath              string      `json:"signPath" gorm:"sign_path"`
	StatusRegistration    int         `json:"statusRegistration" gorm:"status_registration"`
	StoreName             string      `json:"storeName" gorm:"store_name"`
	StorePhoneNumber      string      `json:"storePhoneNumber" gorm:"store_phone_number"`
	WfID                  int         `json:"wfId" gorm:"wf_id"`
	WorkflowExecutionID   int         `json:"workflowExecutionId" gorm:"workflow_execution_id"`
	OwnerWipID            int         `json:"ownerWipId" gorm:"owner_wip_id"`
	SettlementConfigWipID int         `json:"settlementConfigWipId" gorm:"settlement_config_wip_id"`
	IDMerchant            int64       `json:"idMerchant" gorm:"id_merchant"`
	Level                 string 	  `json:"level" gorm:"level"`
	Mcc                   string      `json:"mcc" gorm:"mcc"`
	NMid                  string 	  `json:"nMid" gorm:"n_mid"`
	KategoriBisnis        string 	  `json:"kategoriBisnis" gorm:"kategori_bisnis"`
}

// TableName ...
func (q *MerchantWip) TableName() string {
	return "public.merchant_wip"
}
