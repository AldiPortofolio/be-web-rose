package dbmodels

import "time"

type MerchantDetail struct {
	ErrCode                   string              `json:"errCode"`
	ErrDesc                   string              `json:"errDesc"`
	Status                    string              `json:"status"`
	StatusDescription         string              `json:"statusDescription"`
	ApprovalStatus            string              `json:"approvalStatus"`
	ApprovalStatusDescription string              `json:"approvalStatusDescription"`
	LatestSuggestion          string              `json:"latestSuggestion"`
	LatestSuggestor           string              `json:"latestSuggestor"`
	LatestApproval            string              `json:"latestApproval"`
	LatestApprover            string              `json:"latestApprover"`
	Version                   string              `json:"version"`
	ActionType                string              `json:"actionType"`
	Id                        int64               `json:"id"`
	MerchantOutlets           string              `json:"merchantOutlets"`
	StoreName                 string              `json:"storeName"`
	StoreNamePreprinted       string              `json:"storeNamePreprinted"`
	PhotoLocationLeft         string              `json:"photoLocationLeft"`
	PhotoLocationRight        string              `json:"photoLocationRight"`
	FotoPreprinted            string              `json:"fotoPreprinted"`
	PartnerCustomerId         string              `json:"partnerCustomerId"`
	Patokan                   string              `json:"patokan"`
	AgentSalesRetail          string              `json:"agentSalesRetail"`
	ProfilePictureUrl         string              `json:"profilePictureUrl"`
	OtpOption                 string              `json:"otpOption"`
	NotificationChannel       string              `json:"notificationChannel"`
	ChannelPembayaran         string              `json:"channelPembayaran"`
	MerchantGroupID           int64               `json:"merchantGroupId"`
	MerchantGroup             MerchantGroupDetail `json:"merchantGroup"`
	KtpPath                   string              `json:"ktpPath"`
	SelfiePath                string              `json:"selfiePath"`
	MerchantPhotoPath         string              `json:"merchantPhotoPath"`
	MerchantPhoto2Path        string              `json:"merchantPhoto2Path"`
	SignPath                  string              `json:"signPath"`
	LogoPath                  string              `json:"logoPath"`
	MerchantType              string              `json:"merchantType"`
	JenisUsaha                string              `json:"jenisUsaha"`
	JenisUsahaName            string              `json:"jenisUsahaName"`
	Alamat                    string              `json:"alamat"`
	Kelurahan                 string              `json:"kelurahan"`
	Kecamatan                 string              `json:"kecamatan"`
	Provinsi                  string              `json:"provinsi"`
	ProvinsiName              string              `json:"provinsiName"`
	KabupatenKota             string              `json:"kabupatenKota"`
	KabupatenKotaName         string              `json:"kabupatenKotaName"`
	PostalCode                string              `json:"postalCode"`
	Longitude                 string              `json:"longitude"`
	Latitude                  string              `json:"latitude"`
	StorePhoneNumber          string              `json:"storePhoneNumber"`
	LokasiBisnis              string              `json:"lokasiBisnis"`
	JenisLokasiBisnis         string              `json:"jenisLokasiBisnis"`
	JamOperasional            string              `json:"jamOperasional"`
	SettlementConfigID        int64               `json:"settlementConfigId"`
	SettlementConfig          SettlementDetail    `json:"settlementConfig"`
	OwnerID                   int64               `json:"ownerId"`
	Owner                     OwnerDetail               `json:"owner"`
	HostStatus                string              `json:"hostStatus"`
	HostType                  string              `json:"hostType"`
	ReferralCode              string              `json:"referralCode"`
	AgentName                 string              `json:"agentName"`
	AgentID                   string              `json:"agentID"`
	AgentCompanyID            string              `json:"agentCompanyID"`
	AgentPhoneNumber          string              `json:"agentPhoneNumber"`
	InstitutionID             string              `json:"institutionID"`
	MerchantCategoryCode      string              `json:"merchantCategoryCode"`
	MerchantOutletID          string              `json:"merchantOutletID"`
	MerchantPan               string              `json:"merchantPan"`
	ApiKey                    string              `json:"apiKey"`
	SecretID                  string              `json:"secretID"`
	SecretQuestion            string              `json:"secretQuestion"`
	SecretQuestionAnswer      string              `json:"secretQuestionAnswer"`
	Notes                     string              `json:"notes"`
	Reason                    string              `json:"reason"`
	StatusSuspense            string              `json:"statusSuspense"`
	Level                     string              `json:"level"`
	StatusRegist              string              `json:"statusRegist"`
	MerchantCriteria          string              `json:"merchantCriteria"`
	Mcc                       string              `json:"mcc"`
	Nmid                      string              `json:"nmid"`
	KategoriBisnis            string              `json:"kategoriBisnis"`
	MidInfinitium             string              `json:"midInfinitium"`
	ExistingQrValue           string              `json:"existingQrValue"`
	Category                  string              `json:"category"`
	MallIdDoku                string              `json:"mallIdDoku"`
	SharedKeyDoku             string              `json:"sharedKeyDoku"`
	MerchantUrl               string              `json:"merchantUrl"`
	SuccessUrl                string              `json:"successUrl"`
	FailedUrl                 string              `json:"failedUrl"`
	CallbackUrl               string              `json:"callbackUrl"`
	VaUrl                     string              `json:"vaUrl"`
	VaBca                     string              `json:"vaBca"`
	VaMandiri                 string              `json:"vaMandiri"`
	VaBri                     string              `json:"vaBri"`
	VaLain                    string              `json:"vaLain"`
	VaOttoCash                string              `json:"vaOttoCash"`
	VaTransactionType         string              `json:"vaTransactionType"`
	VaOttoCashCompanyCode     string              `json:"vaOttoCashCompanyCode"`
	VaBcaCompanyCode          string              `json:"vaBcaCompanyCode"`
	VaMandiriCompanyCode      string              `json:"vaMandiriCompanyCode"`
	VaBriCompanyCode          string              `json:"vaBriCompanyCode"`
	VaLainCompanyCode         string              `json:"vaLainCompanyCode"`
	VaBcaSubCompanyCode       string              `json:"vaBcaSubCompanyCode"`
	VaMandiriSubCompanyCode   string              `json:"vaMandiriSubCompanyCode"`
	VaBriSubCompanyCode       string              `json:"vaBriSubCompanyCode"`
	VaLainSubCompanyCode      string              `json:"vaLainSubCompanyCode"`
	VaBcaFee                  string              `json:"vaBcaFee"`
	VaMandiriFee              string              `json:"vaMandiriFee"`
	VaBriFee                  string              `json:"vaBriFee"`
	VaLainFee                 string              `json:"vaLainFee"`
	InquiryUrl                string              `json:"inquiryUrl"`
	PaymentUrl                string              `json:"paymentUrl"`
	CreditPayment             string              `json:"creditPayment"`
	DebitPayment              string              `json:"debitPayment"`
	Qris                      string              `json:"qris"`
	SrId                      string              `json:"srId"`
	PairVerifySimilarity      string              `json:"pairVerifySimilarity"`
	PortalStatus              string              `json:"portalStatus"`
	StatusRegistration        string              `json:"statusRegistration"`
	SelfRegister              string              `json:"selfRegister"`
	RegistrationCode          string              `json:"registrationCode"`
	TanggalSalesAkuisisi      time.Time           `json:"tanggalSalesAkuisisi"`
	TanggalUploadNmid         time.Time           `json:"tanggalUploadNmid"`
	No                        string              `json:"no"`
	StoreName2                string              `json:"storeName2"`
	JmlTerminal               string              `json:"jmlTerminal"`
	Npwp                      string              `json:"npwp"`
	Ktp                       string              `json:"ktp"`
}

// TableName ...
func (q *MerchantDetail) TableName() string {
	return "public.merchant"
}

type MerchantGroupDetail struct {
	ErrCode                   	  string 							 `json:"errCode"`
	ErrDesc                   	  string 							 `json:"errDesc"`
	Status                    	  string 							 `json:"status"`
	StatusDescription         	  string 							 `json:"statusDescription"`
	ApprovalStatus            	  string 							 `json:"approvalStatus"`
	ApprovalStatusDescription 	  string 							 `json:"approvalStatusDescription"`
	LatestSuggestion          	  string 							 `json:"latestSuggestion"`
	LatestSuggestor           	  string 							 `json:"latestSuggestor"`
	LatestApproval            	  string 							 `json:"latestApproval"`
	LatestApprover            	  string 							 `json:"latestApprover"`
	Version                   	  string 							 `json:"version"`
	ActionType                	  string 							 `json:"actionType"`
	ID                        	  int64							     `json:"id"`
	TipeMerchantLookup        	  string 							 `json:"tipeMerchantLookup" gorm:"column:fk_lookup_tipe_merchant"`
	// TipeMerchantLookupName     string                             `json:"tipeMerchantLookupName"`
	GroupPhoto        			  string 							 `json:"groupPhoto"`
	MerchantGroupName 			  string 							 `json:"merchantGroupName"`
	NamaPT            			  string 							 `json:"namaPT"`
	JenisUsahaLookup  			  string 							 `json:"jenisUsahaLookup" gorm:"column:fk_lookup_jenis_usaha"`
	// JenisUsahaLookupName       string                             `json:"jenisUsahaLookupName"`
	Alamat                        string                             `json:"alamat"`
	Rt                            string                             `json:"rt"`
	Rw                            string                             `json:"rw"`
	Kelurahan                     string                             `json:"kelurahan"`
	Kecamatan                     string                             `json:"kecamatan"`
	ProvinsiLookup                string                             `json:"provinsiLookup" gorm:"column:fk_lookup_provinsi"`
	KabupatenKota                 string                             `json:"kabupatenKota" gorm:"column:fk_lookup_kabupaten_kota"`
	Negara                        string                             `json:"negara"`
	Siup                          string                             `json:"siup"`
	SiupFlag                      string                             `json:"siupFlag"`
	Npwp                          string                             `json:"npwp"`
	NpwpFlag                      string                             `json:"npwpFlag"`
	Pks                           string                             `json:"pks"`
	KtpDireksi                    string                             `json:"ktpDireksi"`
	KtpPenanggungJawab            string                             `json:"ktpPenanggungJawab"`
	AktaPendirian                 string                             `json:"aktaPendirian"`
	TandaDaftarPerusahaan         string                             `json:"tandaDaftarPerusahaan"`
	PersetujuanMenkumham          string                             `json:"persetujuanMenkumham"`
	PicGroup                      string                             `json:"picGroup"`
	NoAkta                        string                             `json:"noAkta"`
	NoTelpPic                     string                             `json:"noTelpPic"`
	EmailPic                      string                             `json:"emailPic"`
	WebsitePerusahaan             string                             `json:"websitePerusahaan"`
	MasterDataApprovalId          string                             `json:"masterDataApprovalId"`
	MerchantGroupSettlementInfoID int64                              `json:"merchantGroupSettleInfoId"`
	MerchantGroupSettlementInfo   MerchantGroupSetInfo               `json:"merchantGroupSettlementInfo"`
	MerchantGroupFeeInfoID        int64                              `json:"merchantGroupFeeInfoId"`
	MerchantGroupFeeInfo          MerchantGroupFeeInfo               `json:"merchantGroupFeeInfo"`
	InternalContactPersonID       int64                              `json:"internalContactPersonId"`
	InternalContactPerson         MerchantGroupInternalContactPerson `json:"internalContactPerson"`
	IdMda                         string                             `json:"idMda"`
	PostalCode                    string                             `json:"postalCode"`
	StatusSuspense                string                             `json:"statusSuspense"`
	EnablePartnerCustomerId       string                             `json:"enablePartnerCustomerId"`
	PortalStatus                  string                             `json:"portalStatus"`
	EmailPortal                   string                             `json:"emailPortal"`
}

func (q *MerchantGroupDetail) TableName() string {
	return "public.merchant_group"
}

type OwnerDetail struct {
	ErrCode               string    `json:"errCode"`
	ErrDesc               string    `json:"errDesc"`
	ID                    int64     `json:"id"`
	OwnerTitle            string    `json:"ownerTitle"`
	OwnerFirstName        string    `json:"ownerFirstName"`
	OwnerLastName         string    `json:"ownerLastName"`
	OwnerAddress          string    `json:"ownerAddress"`
	OwnerRt               string    `json:"ownerRt"`
	OwnerRw               string    `json:"ownerRw"`
	OwnerKelurahan        string    `json:"ownerKelurahan"`
	OwnerKecamatan        string    `json:"ownerKecamatan"`
	OwnerKabupaten        string    `json:"ownerKabupaten"`
	OwnerKabupatenName    string    `json:"ownerKabupatenName"`
	OwnerProvinsi         string    `json:"ownerProvinsi"`
	OwnerProvinsiName     string    `json:"ownerProvinsiName"`
	OwnerKodePos          string    `json:"ownerKodePos"`
	OwnerTipeID           string    `json:"ownerTipeID"`
	OwnerNoID             string    `json:"ownerNoID"`
	OwnerTanggalExpiredID time.Time `json:"ownerTanggalExpiredID"`
	OwnerJenisKelamin     string    `json:"ownerJenisKelamin"`
	OwnerNoTelp           string    `json:"ownerNoTelp"`
	OwnerTelpLain         string    `json:"ownerTelpLain"`
	OwnerPekerjaan        string    `json:"ownerPekerjaan"`
	OwnerEmail            string    `json:"ownerEmail"`
	OwnerTempatLahir      string    `json:"ownerTempatLahir"`
	OwnerTanggalLahir     time.Time `json:"ownerTanggalLahir"`
	OwnerNamaIbuKandung   string    `json:"ownerNamaIbuKandung"`
	OwnerNpwp             string    `json:"ownerNpwp"`
}

func (q *OwnerDetail) TableName() string {
	return "public.owner"
}

type SettlementDetail struct {
	ID                            int64  `json:"id"`
	SettlementConfig              string `json:"settlementConfig"`
	SettlementConfigName          string `json:"settlementConfigName"`
	NoRekeningToko                string `json:"noRekeningToko"`
	NamaBankTujuanSettlement      string `json:"namaBankTujuanSettlement"`
	NamaPemilikRekening           string `json:"namaPemilikRekening"`
	TipeRekening                  string `json:"tipeRekening"`
	ReportSettlementConfig        string `json:"reportSettlementConfig"`
	ReportSettlementConfig2       string `json:"reportSettlementConfig2"`
	ReportSettlementConfigName    string `json:"reportSettlementConfigName"`
	SettlementExecutionConfig     string `json:"settlementExecutionConfig"`
	SettlementExecutionConfigName string `json:"settlementExecutionConfigName"`
	SendReportVia                 string `json:"sendReportVia"`
	SendReportUrl                 string `json:"sendReportUrl"`
	ProcessingConfiguration       string `json:"processingConfiguration"`
	ProcessingConfigurationName   string `json:"processingConfigurationName"`
	ProcessingFee                 string `json:"processingFee"`
	ProcessingFeeValue            string `json:"processingFeeValue"`
	RentalEdcFee                  string `json:"rentalEdcFee"`
	Mdr                           string `json:"mdr"`
	MdrEmoneyOnUs                 string `json:"mdrEmoneyOnUs"`
	MdrEmoneyOffUs                string `json:"mdrEmoneyOffUs"`
	MdrDebitOnUs                  string `json:"mdrDebitOnUs"`
	MdrDebitOffUs                 string `json:"mdrDebitOffUs"`
	MdrCreditOnUs                 string `json:"mdrCreditOnUs"`
	MdrCreditOffUs                string `json:"mdrCreditOffUs"`
	OtherFee                      string `json:"otherFee"`
	FmsFee                        string `json:"fmsFee"`
	Status                        string `json:"status"`
	Email                         string `json:"email"`
	SftpHost                      string `json:"sftpHost"`
	SftpUser                      string `json:"sftpUser"`
	SftpPassword                  string `json:"sftpPassword"`
}

func (q *SettlementDetail) TableName() string {
	return "public.settlement_config"
}
