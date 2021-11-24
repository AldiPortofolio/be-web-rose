package services

import (
	"fmt"
	"log"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"time"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// Acquititions struct
type AcquititionsService struct {
	General                   models.GeneralModel
	AcquititionsRepository *db.AcquititionsRepository
}

// InitAcquititions ...
func InitAcquititionsService(gen models.GeneralModel) *AcquititionsService {
	return &AcquititionsService{
		General:          gen,
		AcquititionsRepository: db.InitAcquititionsRepository(),
	}
}

// Filter ...
func (service *AcquititionsService) Filter(req dto.ReqFilterAcquititionsDto) models.Response {
	fmt.Println(">>> Acquititions - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("Acquititions: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "Acquititions: Filter")
	defer span.Finish()

	var res models.Response
	list, total, err := service.AcquititionsRepository.Filter(req)
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

// Save ...
func (service *AcquititionsService) Save(req dto.ReqAcquititionsDto) models.Response {
	fmt.Println(">>> Acquititions - save <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("Acquititions: save",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "Acquititions: save")
	defer span.Finish()

	var res models.Response

	var data dbmodels.Acquititions
	var err error

	if req.ID > 0 {
		data, err = service.AcquititionsRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return res
		}
	}

	var createdAt time.Time
	if req.ID == 0 {
		createdAt = time.Now()
	} else {
		createdAt = req.CreatedAt
	}

	data = dbmodels.Acquititions{
		ID               : req.ID,
		MerchantType     : req.MerchantType    ,
		MerchantGroupId  : req.MerchantGroupId ,
		MerchantCategory : req.MerchantCategory,
		Name             : req.Name            ,
		LogoUrl          : req.LogoUrl         ,
		RegisterUsingId  : req.RegisterUsingId ,
		Sequence         : req.Sequence        ,
		ShowInApp        : req.ShowInApp       ,
		SalesRetails     : req.SalesRetails    ,
		BusinessType     : req.BusinessType    ,
		CreatedAt        : createdAt           ,
		UpdatedAt        : time.Now()          ,
		UpdatedBy        : auth.UserLogin.Name ,
		
	}

	if err := service.AcquititionsRepository.Save(&data); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}

// Save ...
func (service *AcquititionsService) Delete(id int) models.Response {
	fmt.Println(">>> Acquititions - Delete <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("Acquititions: save",
		zap.Any("req", id))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "Acquititions: delete")
	defer span.Finish()

	var res models.Response

	if err := service.AcquititionsRepository.Delete(id); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}
