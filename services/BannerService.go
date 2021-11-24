package services

import (
	"fmt"
	"log"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// BannerService struct
type BannerService struct {
	General          models.GeneralModel
	BannerRepository *db.BannerRepository
}

// InitBannerService ...
func InitBannerService(gen models.GeneralModel) *BannerService {
	return &BannerService{
		General:          gen,
		BannerRepository: db.InitBannerRepository(),
	}
}

// Filter ...
func (service *BannerService) Filter(req dto.ReqBannerDto) models.Response {
	fmt.Println(">>> BannerService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("BannerService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "BannerService: Filter")
	defer span.Finish()

	var res models.Response
	list, total, err := service.BannerRepository.Filter(req)
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
func (service *BannerService) Save(req dto.ReqBannerDto) models.Response {
	fmt.Println(">>> BannerService - save <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("BannerService: save",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "BannerService: save")
	defer span.Finish()

	var res models.Response

	var data dbmodels.Banner
	var err error

	if req.ID > 0 {
		data, err = service.BannerRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return res
		}
	}

	data = dbmodels.Banner{
		ID:              req.ID,
		UserCategoryId:  req.UserCategoryId,
		LevelMerchantId: req.LevelMerchantId,
		Name:            req.Name,
		AdsImage:        req.AdsImage,
		AdsLink:         req.AdsLink,
		Seq:             req.Seq,
		Status:          req.Status,
		BannerName:      req.BannerName,
		DetailBanner:    req.DetailBanner,
	}

	if err := service.BannerRepository.Save(&data); err != nil {
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
func (service *BannerService) Delete(id int) models.Response {
	fmt.Println(">>> BannerService - Delete <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("BannerService: save",
		zap.Any("req", id))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "BannerService: delete")
	defer span.Finish()

	var res models.Response

	if err := service.BannerRepository.Delete(id); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}
