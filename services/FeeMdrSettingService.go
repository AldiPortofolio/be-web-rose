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

// FeeMdrSettingService ..
type FeeMdrSettingService struct {
	General                 models.GeneralModel
	FeeMdrSettingRepository *db.FeeMdrSettingRepository
	MerchantRepository      *db.MerchantRepository
	MdrBankRepository       *db.MdrBankRepository
	MdrTenorRepository      *db.MdrTenorRepository
}

// InitFeeMdrSettingService ..
func InitFeeMdrSettingService(gen models.GeneralModel) *FeeMdrSettingService {
	return &FeeMdrSettingService{
		General:                 gen,
		FeeMdrSettingRepository: db.InitFeeMdrSettingRepository(),
		MerchantRepository:      db.InitMerchantRepository(),
		MdrBankRepository:       db.InitMdrBankRepository(),
		MdrTenorRepository:      db.InitMdrTenorRepository(),
	}
}

// Save ..
func (service *FeeMdrSettingService) Save(req dto.ReqFeeMdrSettingDto) models.Response {
	fmt.Println(">> FeeMdrSettingService - Save <<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("FeeMdrSettingService: Save",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "FeeMdrSettingService: Save")
	defer span.Finish()

	var res models.Response

	var data dbmodels.FeeMdrSetting
	var err error

	if req.ID > 0 {
		data, err = service.FeeMdrSettingRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return res
		}
	}

	merchantChan := make(chan dbmodels.Merchant, 1)
	merchantChan <- service.getMerchant(req.MidMerchant)

	mdrBankChan := make(chan dbmodels.MdrBank, 1)
	mdrBankChan <- service.getMdrBank(req.Bank)

	mdrTenorChan := make(chan dbmodels.MdrTenor, 1)
	mdrTenorChan <- service.getMdrTenor(req.Tenor)

	if merchant := <-merchantChan; merchant.ID == 0 || req.MidMerchant == "" {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = "Merchant not found"
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

	data.MidMerchant = req.MidMerchant
	data.MidBank = req.MidBank
	data.SecretID = req.SecretID
	data.Bank = req.Bank
	data.Tenor = req.Tenor
	data.BankMdr = req.BankMdr
	data.BankMdrCredit = req.BankMdrCredit
	data.PlanId = req.PlanId
	data.MerchantMdr = req.MerchantMdr
	data.MerchantMdrCredit = req.MerchantMdrCredit
	data.MerchantFeeType = req.MerchantFeeType
	data.MerchantFee = req.MerchantFee
	data.CustomerFeeType = req.CustomerFeeType
	data.CustomerFee = req.CustomerFee
	data.Status = req.Status
	data.UpdatedAt = time.Now()
	data.UpdatedBy = auth.UserLogin.Name

	if err := service.FeeMdrSettingRepository.Save(&data); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}

func (service *FeeMdrSettingService) getMerchant(merchantOutletId string) dbmodels.Merchant {
	data := dbmodels.Merchant{
		MerchantOutletID: merchantOutletId,
	}
	merchant, getMerchantErr := service.MerchantRepository.Get(data)
	if getMerchantErr != nil {
		sugarLogger := service.General.OttoZaplog
		sugarLogger.Error(fmt.Sprintf("Failed to get merchant: %v", getMerchantErr))
	}
	return merchant
}

func (service *FeeMdrSettingService) getMdrBank(bankCode string) dbmodels.MdrBank {
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

func (service *FeeMdrSettingService) getMdrTenor(tenorCode string) dbmodels.MdrTenor {
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

// Filter ..
func (service *FeeMdrSettingService) Filter(req dto.ReqFeeMdrSettingDto) models.Response {
	fmt.Println(">>> FeeMdrSettingService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("FeeMdrSettingService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "FeeMdrSettingService: Filter")
	defer span.Finish()

	var res models.Response
	list, total, err := service.FeeMdrSettingRepository.Filter(req)
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

// FindAll ..
func (service *FeeMdrSettingService) FindAll() models.Response {
	fmt.Println(">>> FeeMdrSettingService - FindAll <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("FeeMdrSettingService: FindAll",
		zap.Any("req", ""))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "FeeMdrSettingService: FindAll")
	defer span.Finish()

	var res models.Response
	list, total, err := service.FeeMdrSettingRepository.FindAll()
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
