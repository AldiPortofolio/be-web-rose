package services

import (
	"fmt"
	"log"
	"net/http"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"strconv"
	"time"

	utilhttp "rose-be-go/utils/http"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	ottoutils "ottodigital.id/library/utils"
)

var bpAuthAddress string

func init()  {
	bpAuthAddress = ottoutils.GetEnv("BP_AUTH_ACTIVATION", "http://13.228.25.85:8000/bpauth/v0.1.0/activation/")
	

}

type MerchantGroupService struct {
	General models.GeneralModel
	MerchantGroupRepository *db.MerchantGroupRepository
	LookupRepository *db.LookupRepository
	MerchantGroupSetInfoRepository *db.MerchantGroupSetInfoRepository
	MerchantGroupInternalContactPersonRepository *db.MerchantGroupInternalContactPersonRepository
	MerchantGroupFeeInfoRepository *db.MerchantGroupFeeInfoRepository
}

func InitMerchantGroupService(gen models.GeneralModel) *MerchantGroupService{
	return &MerchantGroupService{
		General: gen,
		MerchantGroupRepository: db.InitMerchantGroupRepository(),
	}
}

var ActionType = map[int]string{
	0:"CREATE",
	1:"DELETE",
	2:"EDIT",
	3:"SUSPEND",
	4:"UNSUSPEND",
}

func (service *MerchantGroupService) GetDetail(id int64) dto.MerchantGroupDtoRes {
	fmt.Println(">>> MerchantGroupService - GetDetail <<<")

	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MerchantGroupService: GetDetail",
		zap.Any("req", id))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "ReportFinishService: Send")
	defer span.Finish()

	var res dto.MerchantGroupDtoRes

	merchantGroup, err := service.MerchantGroupRepository.GetDetail(dbmodels.MerchantGroup{ID:id})
	if err != nil {
		fmt.Println("error : ", err)
	}

	typeMerchantId, _ := strconv.Atoi(merchantGroup.FkLookupTipeMerchant)
	typeMerchant, _ := service.LookupRepository.Get(dbmodels.Lookup{
		Id:             int64(typeMerchantId),
		LookupGroup:      "TIPE_MERCHANT",
	})

	jenisUsahaId, _ := strconv.Atoi(merchantGroup.FkLookupJenisUsaha)
	jenisUsaha, _ := service.LookupRepository.Get(dbmodels.Lookup{
		Id:             int64(jenisUsahaId),
		LookupGroup:      "JENIS_USAHA",
	})



	mGrpSetInfo, _ := service.MerchantGroupSetInfoRepository.Get(dbmodels.MerchantGroupSetInfo{Id:merchantGroup.MerchantGroupSettleInfoID})

	contactPerson, _ := service.MerchantGroupInternalContactPersonRepository.Get(dbmodels.MerchantGroupInternalContactPerson{Id:merchantGroup.InternalContactPersonID})

	feeInfo, _ := service.MerchantGroupFeeInfoRepository.Get(dbmodels.MerchantGroupFeeInfo{Id:merchantGroup.MerchantGroupFeeInfoID})

	feeInfoId, _ := strconv.Atoi(feeInfo.FkLookupProcessingFee)
	feeValue, _ := service.LookupRepository.Get(dbmodels.Lookup{
		Id:             int64(feeInfoId),
		LookupGroup:      "PROCESSING_FEE",
	})

	actionType := ActionType[merchantGroup.ActionType]

	res = dto.MerchantGroupDtoRes{
		ErrCode:                     nil,
		ErrDesc:                     nil,
		Status:                      nil,
		StatusDescription:           nil,
		ApprovalStatus:              &merchantGroup.ApprovalStatus,
		ApprovalStatusDescription:   nil,
		LatestSuggestion:            &merchantGroup.LatestSuggestion,
		LatestSuggestor:             &merchantGroup.LatestSuggestor,
		LatestApproval:              &merchantGroup.LatestApproval,
		LatestApprover:              &merchantGroup.LatestApprover,
		Version:                     &merchantGroup.Version,
		ActionType:                  &actionType,
		Id:                          &merchantGroup.ID,
		TipeMerchantLookup:          &merchantGroup.FkLookupTipeMerchant,
		TipeMerchantLookupName:      &typeMerchant.Name,
		GroupPhoto:                  &merchantGroup.GroupPhoto,
		MerchantGroupName:           &merchantGroup.MerchantGroupName,
		NamaPT:                      &merchantGroup.NamaPt,
		JenisUsahaLookup:            &merchantGroup.FkLookupJenisUsaha,
		JenisUsahaLookupName:        &jenisUsaha.Name,
		Alamat:                      &merchantGroup.Alamat,
		Rt:                          &merchantGroup.Rt,
		Rw:                          &merchantGroup.Rw,
		Kelurahan:                   &merchantGroup.Kelurahan,
		Kecamatan:                   &merchantGroup.Kecamatan,
		ProvinsiLookup:              &merchantGroup.FkLookupProvinsi,
		KabupatenKota:               &merchantGroup.FkLookupKabupatenKota,
		Negara:                      &merchantGroup.Negara,
		Siup:                        &merchantGroup.Siup,
		SiupFlag:                    &merchantGroup.SiupFlag,
		Npwp:                        &merchantGroup.Npwp,
		NpwpFlag:                    &merchantGroup.NpwpFlag,
		Pks:                         &merchantGroup.Pks,
		KtpDireksi:                  &merchantGroup.KtpDireksi,
		KtpPenanggungJawab:          &merchantGroup.KtpPenanggungJawab,
		AktaPendirian:               &merchantGroup.AktaPendirian,
		TandaDaftarPerusahaan:       &merchantGroup.TandaDaftarPerusahaan,
		PersetujuanMenkumham:        &merchantGroup.PersetujuanMenhumkan,
		PicGroup:                    &merchantGroup.PicGroup,
		NoTelpPic:                   &merchantGroup.NoTelpPic,
		EmailPic:                    &merchantGroup.EmailPic,
		WebsitePerusahaan:           &merchantGroup.WebsitePerusahaan,
		MasterDataApprovalId:        nil,
		MerchantGroupSettlementInfo: dto.MerchantGroupSettlementInfoDto{
			Id:                              mGrpSetInfo.Id,
			NomorRekening:                   mGrpSetInfo.NomorRekening,
			NamaBankTujuanSettlement:        mGrpSetInfo.NamaBankTujuanSettlement,
			NamaPemilikRekening:             mGrpSetInfo.NamaPemilikRekening,
			TipeRekening:                    mGrpSetInfo.TipeRekening,
			ReportSettlementConfigLookup:    mGrpSetInfo.FkLookupReportSettleCfg,
			SettlementExecutionConfigLookup: mGrpSetInfo.FkLookupSettlementExecCfg,
			SendReportViaLookup:             mGrpSetInfo.FkLookupSendReportVia,
			SendReportUrl:                   mGrpSetInfo.SendReportUrl,
		},
		InternalContactPerson:       dto.InternalContactPersonDto{
			Id:                 contactPerson.Id,
			BusinessPic:        contactPerson.BusinessPic,
			TechnicalPic:       contactPerson.TechnicalPic,
			SettleOperationPic: contactPerson.SettleOperationPic,
			Notes:              contactPerson.Notes,
		},
		MerchantGroupFeeInfo:dto.MerchantGroupFeeInfoDto{
			Id:                  feeInfo.Id,
			ProcessingFeeLookup: feeInfo.FkLookupProcessingFee,
			ProcessingFeeValue:  feeValue.Version,
			RentalEdcFee:        feeInfo.RentalEdcFee,
			MdrLookup:           feeInfo.FkLookupMdr,
			MdrEmoneyOnUs:       feeInfo.MdrEmoneyOnUsValue,
			MdrEmoneyOffUs:      feeInfo.MdrEmoneyOffUsValue,
			MdrDebitOnUs:        feeInfo.MdrDebitOnUsValue,
			MdrDebitOffUs:       feeInfo.MdrDebitOffUsValue,
			MdrCreditOnUs:       feeInfo.MdrCreditOnUsValue,
			MdrCreditOffUs:      feeInfo.MdrEmoneyOffUsValue,
			OtherFee:            feeInfo.OtherFee,
			FmsFee:              feeInfo.FmsFee,
		},
		IdMda:                       nil,
		PostalCode:                  &merchantGroup.PostalCode,
		StatusSuspense:              &merchantGroup.StatusSuspense,
		EnablePartnerCustomerId:     merchantGroup.EnablePartnerCustomerId,
		PortalStatus:                merchantGroup.PortalStatus,
		PortalEmail:                 merchantGroup.EmailPortal,
	}

	return res
}


func (service *MerchantGroupService) MerchantGroupActivationService(id []int) models.Response {
	fmt.Println(">>> MerchantGroupService - MerchantGroupActivationService <<<")

	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MerchantGroupService: MerchantGroupActivationService",
		zap.Any("req", id))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MerchantGroupService - MerchantGroupActivationService")
	defer span.Finish()

	var res models.Response
	// err := service.MerchantGroupRepository.MerchantGroupActivationBp(id)
	// if err != nil {
	// 	res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
	// 	res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		
	// 	return res
	// }

	for i := 0 ; i < len(id) ; i++ {
		merchantId := id[i] 
		now := time.Now()
		email := strconv.Itoa( now.Day())+ strconv.Itoa( now.Nanosecond()) + "@mail.com"

		
		reqData := dto.BpActivationReq{
			Id: int64(merchantId),
			Category : 1,
			Action: "activation",
			Email : email,
			Password : "passwordRahasia1!",
		}
		
		_, key := BpActivation(reqData)

		req := dto.ReqSendBpActivationDto{
			Username: reqData.Email,
			Password: "passwordRahasia1!",
			NewPassword: "Asd#1234567",
		}
		resBpActivation, err := SendBpActivation( req, key )
		if err != nil {
			log.Println(req, "Err Res BP Auth Activation--> ", err)
		}
		log.Println("Res BP Auth Activation--> ", string(resBpActivation))
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}

func SendBpActivation(request interface{}, key string ) ([]byte, error) {
	url := bpAuthAddress + key
	header := make(http.Header)
	header.Add("Content-Type", "application/json")


	fmt.Println("Headerrrrrr======>", header)


	data, err := utilhttp.HTTPPostWithHeader(url, request, header)
	fmt.Println("xxxx-----------xxxx")
	fmt.Println("urlSvr", request)
	fmt.Println("header", header)
	fmt.Println("err", err)
	fmt.Println("xxxx-----------xxxx")


	return data, err
}
