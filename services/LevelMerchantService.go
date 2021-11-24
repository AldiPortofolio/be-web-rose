package services

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
)

// LevelMerchantService struct
type LevelMerchantService struct {
	General models.GeneralModel
	LevelMerchantRepository *db.LevelMerchantRepository
}

// InitLevelMerchantService ...
func InitLevelMerchantService(gen models.GeneralModel) *LevelMerchantService {
	return &LevelMerchantService{
		General:gen,
		LevelMerchantRepository: db.InitLevelMerchantRepository(),
	}
}

// GetAll ...
func (service *LevelMerchantService) GetAll(limit int64) models.Response {
	fmt.Println(">>> LevelMerchantService - GetAll <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("LevelMerchantService: GetAll",
		zap.Any("limit", limit))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "LevelMerchantService: GetAll")
	defer span.Finish()

	var res models.Response
	data, err := service.LevelMerchantRepository.GetAll(limit)
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
