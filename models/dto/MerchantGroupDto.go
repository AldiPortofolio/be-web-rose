package dto

import "time"

type MerchantGroupDtoRes struct {
	ErrCode     				*string  		`json:"errCode"`
	ErrDesc     				*string  		`json:"errDesc"`
	Status      				*int  			`json:"status"`
	StatusDescription			*string  		`json:"statusDescription"`
	ApprovalStatus 				*int 			`json:"approvalStatus"`
	ApprovalStatusDescription	*string  		`json:"approvalStatusDescription"`
	LatestSuggestion 			*time.Time    	`json:"latestSuggestion"`
	LatestSuggestor 			*string 		`json:"latestSuggestor"`
	LatestApproval  			*time.Time 		`json:"latestApproval"`
	LatestApprover  			*string  		`json:"latestApprover"`
	Version 					*int  			`json:"version"`
	ActionType 					*string       	`json:"actionType"`
	Id 							*int64 			`json:"id"`
	TipeMerchantLookup 			*string  		`json:"tipeMerchantLookup"`
	TipeMerchantLookupName  	*string 		`json:"tipeMerchantLookupName"`
	GroupPhoto 					*string 		`json:"groupPhoto"`
	MerchantGroupName  			*string 		`json:"merchantGroupName"`
	NamaPT						*string  		`json:"namaPT"`
	JenisUsahaLookup 			*string 		`json:"jenisUsahaLookup"`
	JenisUsahaLookupName        *string 		`json:"jenisUsahaLookupName"`
	Alamat 						*string 		`json:"alamat"`
	Rt 							*string 		`json:"rt"`
	Rw 							*string 		`json:"rw"`
	Kelurahan 					*string 		`json:"kelurahan"`
	Kecamatan 					*string 		`json:"kecamatan"`
	ProvinsiLookup 				*string 		`json:"provinsiLookup"`
	KabupatenKota				*string 		`json:"kabupatenKota"`
	Negara						*string 		`json:"negara"`
	Siup 						*string  		`json:"siup"`
	SiupFlag 					*int64 			`json:"siupFlag"`
	Npwp 						*string 		`json:"npwp"`
	NpwpFlag 					*int64 			`json:"npwpFlag"`
	Pks  						*string 		`json:"pks"`
	KtpDireksi 					*string 		`json:"ktpDireksi"`
	KtpPenanggungJawab 			*string 		`json:"ktpPenanggungJawab"`
	AktaPendirian 				*string 		`json:"aktaPendirian"`
	TandaDaftarPerusahaan 		*string 		`json:"tandaDaftarPerusahaan"`
	PersetujuanMenkumham 		*string 		`json:"persetujuanMenkumham"`
	PicGroup 					*string 		`json:"picGroup"`
	NoTelpPic 					*string 		`json:"noTelpPic"`
	EmailPic					*string 		`json:"emailPic"`
	WebsitePerusahaan 			*string 		`json:"websitePerusahaan"`
	MasterDataApprovalId 		*string 		`json:"masterDataApprovalId"`
	MerchantGroupSettlementInfo MerchantGroupSettlementInfoDto 	`json:"merchantGroupSettlementInfo"`
	MerchantGroupFeeInfo 		MerchantGroupFeeInfoDto			`json:"merchantGroupFeeInfo"`
	InternalContactPerson 		InternalContactPersonDto 		`json:"internalContactPerson"`
	IdMda 						*string 		`json:"idMda"`
	PostalCode 					*string 		`json:"postalCode"`
	StatusSuspense 				*bool 			`json:"statusSuspense"`
	EnablePartnerCustomerId 	bool 			`json:"enablePartnerCustomerId"`
	PortalStatus 				int 			`json:"portalStatus"`
	PortalEmail 				string 			`json:"portalEmail"`
}

type MerchantGroupSettlementInfoDto struct {
	Id   						int64 			`json:"id"`
	NomorRekening 				string 			`json:"nomorRekening"`
	NamaBankTujuanSettlement 	string 			`json:"namaBankTujuanSettlement"`
	NamaPemilikRekening 		string 			`json:"namaPemilikRekening"`
	TipeRekening 				string 			`json:"tipeRekening"`
	ReportSettlementConfigLookup string 		`json:"reportSettlementConfigLookup"`
	SettlementExecutionConfigLookup string 		`json:"settlementExecutionConfigLookup"`
	SendReportViaLookup 		string 			`json:"sendReportViaLookup"`
	SendReportUrl 				string 			`json:"sendReportUrl"`
}

type MerchantGroupFeeInfoDto struct {
	Id 						int64 				`json:"id"`
	ProcessingFeeLookup 	string 				`json:"processingFeeLookup"`
	ProcessingFeeValue 		int 				`json:"processingFeeValue"`
	RentalEdcFee 			float64 			`json:"rentalEdcFee"`
	MdrLookup 				string 				`json:"mdrLookup"`
	MdrEmoneyOnUs 			float64 				`json:"mdrEmoneyOnUs"`
	MdrEmoneyOffUs 			float64 				`json:"mdrEmoneyOffUs"`
	MdrDebitOnUs 			float64 				`json:"mdrDebitOnUs"`
	MdrDebitOffUs 			float64 				`json:"mdrDebitOffUs"`
	MdrCreditOnUs  			float64 				`json:"mdrCreditOnUs"`
	MdrCreditOffUs			float64 				`json:"mdrCreditOffUs"`
	OtherFee				float64 				`json:"otherFee"`
	FmsFee					float64 				`json:"fmsFee"`
}

type InternalContactPersonDto struct {
	Id 						int64 				`json:"id"`
	BusinessPic 			string 				`json:"businessPic"`
	TechnicalPic 			string 				`json:"technicalPic"`
	SettleOperationPic 		string 				`json:"settleOperationPic"`
	Notes  					string 				`json:"notes"`
}

type MerchantGroupListIdDto struct {
	Id 						[]int				`json:"id"`
	
}
