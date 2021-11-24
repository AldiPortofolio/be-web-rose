package services

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"log"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

// CategoryLevelFeatureService struct
type CategoryLevelFeatureService struct {
	General models.GeneralModel
	CategoryLevelFeatureRepository *db.CategoryLevelFeatureRepository
}

// InitCategoryLevelFeatureService ...
func InitCategoryLevelFeatureService(gen models.GeneralModel) *CategoryLevelFeatureService {
	return &CategoryLevelFeatureService{
		General:gen,
		CategoryLevelFeatureRepository: db.InitCategoryLevelFeatureRepository(),
	}
}

// Filter ...
func (service *CategoryLevelFeatureService) Filter(req dto.ReqCategoryLevelFeatureDto) models.Response {
	fmt.Println(">>> CategoryLevelFeatureService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("CategoryLevelFeatureService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "CategoryLevelFeatureService: Filter")
	defer span.Finish()

	var res models.Response
	list, total, err := service.CategoryLevelFeatureRepository.Filter(req)
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
func (service *CategoryLevelFeatureService) Save(req dto.ReqCategoryLevelFeatureDto) models.Response {
	fmt.Println(">>> CategoryLevelFeatureService - save <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("CategoryLevelFeatureService: save",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "CategoryLevelFeatureService: save")
	defer span.Finish()

	var res models.Response

	var data dbmodels.CategoryLevelFitur
	var err error

	if req.ID > 0  {
		data, err = service.CategoryLevelFeatureRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return res
		}
	}

	data = dbmodels.CategoryLevelFitur{
		ID:              req.ID,
		UserCategoryId:  req.UserCategoryId,
		LevelMerchantId: req.LevelMerchantId,
		FiturProductId:  req.FiturProductId,
		Status:          req.Status,
	}

	if err:=service.CategoryLevelFeatureRepository.Save(&data); err!=nil{
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}
