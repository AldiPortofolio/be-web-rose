package services

import (
	"fmt"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// KategoriBisnisService struct
type KategoriBisnisService struct {
	General models.GeneralModel
	KategoriBisnisRepository *db.KategoriBisnisRepository
}

// InitKategoriBisnisService ...
func InitKategoriBisnisService(gen models.GeneralModel) *KategoriBisnisService {
	return &KategoriBisnisService{
		General:gen,
		KategoriBisnisRepository: db.InitKategoriBisnisRepository(),
	}
}

// FindAll ...
func (service *KategoriBisnisService) FindAll() models.Response {
	fmt.Println(">>> KategoriBisnisService - FindAll <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("KategoriBisnisService: FindAll",
		zap.Any("req", ""))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "KategoriBisnisService: FindAll")
	defer span.Finish()

	var res models.Response
	list, err := service.KategoriBisnisRepository.GetAll()
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.Contents = list
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.TotalData = len(list)
	res.Contents = list

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}