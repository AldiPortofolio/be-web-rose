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

type ProfileThemeService struct {
	General models.GeneralModel
	ProfileThemeRepository *db.ProfileThemeRepository
}

func InitProfileThemeService(gen models.GeneralModel) *ProfileThemeService {
	return &ProfileThemeService{
		General:gen,
		ProfileThemeRepository: db.InitProfileThemeRepository(),
	}
}

func (service *ProfileThemeService) Filter(req dto.ReqProfileThemeDto) models.Response {
	fmt.Println(">>> ProfileThemeService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("ProfileThemeService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "ProfileThemeService: Filter")
	defer span.Finish()

	var res models.Response
	list, total, err := service.ProfileThemeRepository.Filter(req)
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

func (service *ProfileThemeService) Save(req dto.ReqProfileThemeDto) models.Response {
	fmt.Println(">>> ProfileThemeService - save <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("ProfileThemeService: save",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "ProfileThemeService: save")
	defer span.Finish()

	var res models.Response

	var data dbmodels.ProfileTheme
	var err error

	if req.ID > 0  {
		data, err = service.ProfileThemeRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return res
		}
	}

	data = dbmodels.ProfileTheme{
		ID:                     req.ID,
		UserCategoryId:         req.UserCategoryId,
		LevelMerchantId:        req.LevelMerchantId,
		DashboardTopBackground: req.DashboardTopBackground,
		ThemeColor:             req.ThemeColor,
		DashboardLogo:          req.DashboardLogo,
		DashboardText:          req.DashboardText,
		ProfileBackgroundImage: req.ProfileBackgroundImage,
		Status:                 req.Status,
		FontColor: req.FontColor,
		Url: req.Url,
	}

	if err:=service.ProfileThemeRepository.Save(&data); err!=nil{
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}
