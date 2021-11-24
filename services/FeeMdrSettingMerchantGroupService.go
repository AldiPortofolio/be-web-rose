package services

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"log"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"time"
)

// FeeMdrSettingMerchantGroupService struct
type FeeMdrSettingMerchantGroupService struct {
	General                              models.GeneralModel
	FeeMdrSettingMerchantGroupRepository *db.FeeMdrSettingMerchantGroupRepository
	MerchantGroupRepository              *db.MerchantGroupRepository
	MdrBankRepository                    *db.MdrBankRepository
	MdrTenorRepository                   *db.MdrTenorRepository
}

// InitFeeMdrSettingMerchantGroupService ...
func InitFeeMdrSettingMerchantGroupService(gen models.GeneralModel) *FeeMdrSettingMerchantGroupService {
	return &FeeMdrSettingMerchantGroupService{
		General:                              gen,
		FeeMdrSettingMerchantGroupRepository: db.InitFeeMdrSettingMerchantGroupRepository(),
		MerchantGroupRepository:              db.InitMerchantGroupRepository(),
		MdrBankRepository:                    db.InitMdrBankRepository(),
		MdrTenorRepository:                   db.InitMdrTenorRepository(),
	}
}

// Save ...
func (service *FeeMdrSettingMerchantGroupService) Save(req dto.ReqFeeMdrSettingMerchantGroupDto) models.Response {
	fmt.Println(">> FeeMdrSettingMerchantGroupService - Save <<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("FeeMdrSettingMerchantGroupService: Save",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "FeeMdrSettingMerchantGroupService: Save")
	defer span.Finish()

	var res models.Response

	var data dbmodels.FeeMdrSettingMerchantGroup
	var err error

	if req.ID > 0 {
		data, err = service.FeeMdrSettingMerchantGroupRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return res
		}
	}

	merchantGroupChan := make(chan dbmodels.MerchantGroup, 1)
	merchantGroupChan <- service.getMerchantGroup(req.IdMerchantGroup)

	mdrBankChan := make(chan dbmodels.MdrBank, 1)
	mdrBankChan <- service.getMdrBank(req.Bank)

	mdrTenorChan := make(chan dbmodels.MdrTenor, 1)
	mdrTenorChan <- service.getMdrTenor(req.Tenor)

	if merchantGroup := <-merchantGroupChan; merchantGroup.ID == 0 || req.IdMerchantGroup == 0 {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = "Merchant group not found"
		return res
	}

	if mdrBank := <-mdrBankChan; mdrBank.ID == 0 || req.Bank == "" {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = "Bank not found"
		return res
	}

	if mdrTenor := <-mdrTenorChan; mdrTenor.ID == 0 || req.Tenor == "" {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = "Tenor not found"
		return res
	}

	data.IdMerchantGroup = req.IdMerchantGroup
	data.MidBank = req.MidBank
	data.SecretID = req.SecretID
	data.Bank = req.Bank
	data.Tenor = req.Tenor
	data.BankMdr = req.BankMdr
	data.MerchantMdr = req.MerchantMdr
	data.MerchantFeeType = req.MerchantFeeType
	data.MerchantFee = req.MerchantFee
	data.CustomerFeeType = req.CustomerFeeType
	data.CustomerFee = req.CustomerFee
	data.UpdatedAt = time.Now()
	data.UpdatedBy = auth.UserLogin.Name

	if err := service.FeeMdrSettingMerchantGroupRepository.Save(&data); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}

func (service *FeeMdrSettingMerchantGroupService) getMerchantGroup(merchantGroupId int64) dbmodels.MerchantGroup {
	data := dbmodels.MerchantGroup{
		ID: merchantGroupId,
	}
	merchantGroup, getMerchantGroupErr := service.MerchantGroupRepository.Get(data)
	if getMerchantGroupErr != nil {
		sugarLogger := service.General.OttoZaplog
		sugarLogger.Error(fmt.Sprintf("Failed to get merchant group: %v", getMerchantGroupErr))
	}
	return merchantGroup
}

func (service *FeeMdrSettingMerchantGroupService) getMdrBank(bankCode string) dbmodels.MdrBank {
	data := dbmodels.MdrBank{
		BankCode: bankCode,
	}
	mdrBank, getMdrBankErr := service.MdrBankRepository.Get(data)
	if getMdrBankErr != nil {
		sugarLogger := service.General.OttoZaplog
		sugarLogger.Error(fmt.Sprintf("Failed to get mdr bank: %v", getMdrBankErr))
	}
	return mdrBank
}

func (service *FeeMdrSettingMerchantGroupService) getMdrTenor(tenorCode string) dbmodels.MdrTenor {
	data := dbmodels.MdrTenor{
		TenorCode: tenorCode,
	}
	mdrTenor, getMdrTenorErr := service.MdrTenorRepository.Get(data)
	if getMdrTenorErr != nil {
		sugarLogger := service.General.OttoZaplog
		sugarLogger.Error(fmt.Sprintf("Failed to get mdr tenor: %v", getMdrTenorErr))
	}
	return mdrTenor
}

// Filter ...
func (service *FeeMdrSettingMerchantGroupService) Filter(req dto.ReqFeeMdrSettingMerchantGroupDto) models.Response {
	fmt.Println(">>> FeeMdrSettingMerchantGroupService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("FeeMdrSettingMerchantGroupService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "FeeMdrSettingMerchantGroupService: Filter")
	defer span.Finish()

	var res models.Response
	list, total, err := service.FeeMdrSettingMerchantGroupRepository.Filter(req)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.TotalData = total
		res.Contents = list
		return res
	}

	log.Println("total -->", total)

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.TotalData = total
	res.Contents = list

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}

// FindAll ...
func (service *FeeMdrSettingMerchantGroupService) FindAll() models.Response {
	fmt.Println(">>> FeeMdrSettingMerchantGroupService - FindAll <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("FeeMdrSettingMerchantGroupService: FindAll",
		zap.Any("req", ""))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "FeeMdrSettingMerchantGroupService: FindAll")
	defer span.Finish()

	var res models.Response
	list, total, err := service.FeeMdrSettingMerchantGroupRepository.FindAll()
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.Contents = list
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.TotalData = total
	res.Contents = list

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}
