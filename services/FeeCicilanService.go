package services

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"time"
)

// FeeCicilanService struct
type FeeCicilanService struct {
	General models.GeneralModel
	FeeCicilanRepository *db.FeeCicilanRepository
}

// InitFeeCicilanService ...
func InitFeeCicilanService(gen models.GeneralModel) *FeeCicilanService {
	return &FeeCicilanService{
		General:gen,
		FeeCicilanRepository: db.InitFeeCicilanRepository(),
	}
}

// Save ...
func (service *FeeCicilanService) Save(req dto.ReqFeeCicilanSettingDto) models.Response  {
	fmt.Println(">> FeeCicilanService - Save <<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("FeeCicilanService: Save",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "FeeCicilanService: Save")
	defer span.Finish()

	var res models.Response

	data := dbmodels.FeeCicilanSetting{
		VaLainnyaFee: req.VaLainnyaFee,
		VaMandiriFee: req.VaMandiriFee,
		VaBcaFee: req.VaBcaFee,
		AdminFeeInfinitium: req.AdminFeeInfinitium,
		AdminFeeDoku: req.AdminFeeDoku,
		User: auth.UserLogin.Name,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	if err := service.FeeCicilanRepository.Save(&data); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res

}

// Find ...
func (service *FeeCicilanService) Find() models.Response {
	fmt.Println(">>> FeeCicilanService - Find <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("FeeCicilanService: Find",
		zap.Any("req", ""))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "FeeCicilanService: Find")
	defer span.Finish()

	var res models.Response
	data, err := service.FeeCicilanRepository.Find()
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.Contents = data
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.Contents = data

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}
