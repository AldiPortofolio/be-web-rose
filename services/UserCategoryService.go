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

type UserCategoryService struct {
	General                models.GeneralModel
	UserCategoryRepository *db.UserCategoryRepository
}

func InitUserCategoryService(gen models.GeneralModel) *UserCategoryService {
	return &UserCategoryService{
		General:                gen,
		UserCategoryRepository: db.InitUserCategoryRepository(),
	}
}

func (service *UserCategoryService) Filter(req dto.ReqUserCategoryDto) models.Response {
	fmt.Println(">>> UserCategoryService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("UserCategoryService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "UserCategoryService: Filter")
	defer span.Finish()

	var res models.Response
	list, total, err := service.UserCategoryRepository.Filter(req)
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

func (service *UserCategoryService) DropdownList() models.Response {
	fmt.Println(">>> UserCategoryService - DropdownList <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("UserCategoryService: DropdownList",
		zap.Any("req", ""))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "UserCategoryService: DropdownList")
	defer span.Finish()

	var res models.Response
	list, total, err := service.UserCategoryRepository.DropdownList()
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

func (service *UserCategoryService) Save(req dto.ReqUserCategoryDto) models.Response {
	fmt.Println(">>> UserCategoryService - save <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("UserCategoryService: save",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "UserCategoryService: save")
	defer span.Finish()

	var res models.Response

	var data dbmodels.UserCategory
	var err error

	if req.ID > 0 {
		data, err = service.UserCategoryRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return res
		}
	}

	data = dbmodels.UserCategory{
		ID:    req.ID,
		Code:  req.Code,
		Name:  req.Name,
		Notes: req.Notes,
		AppID: req.AppID,
		Logo:  req.Logo,
		Seq:   req.Seq,
	}

	if err := service.UserCategoryRepository.Save(&data); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}
