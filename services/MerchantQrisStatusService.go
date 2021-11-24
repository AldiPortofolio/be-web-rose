package services

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dto"
)

type MerchantQrisStatusService struct {
	General models.GeneralModel
	MerchantQrisStatusRepository *db.MerchantQrisStatusRepository
}

func InitMerchantQrisStatusService(gen models.GeneralModel) *MerchantQrisStatusService{
	return &MerchantQrisStatusService{
		General: gen,
		MerchantQrisStatusRepository: db.InitMerchantQrisStatusRepository(),
	}
}

func (service *MerchantQrisStatusService) Filter(req dto.ReqFilterDto) models.Response {

	fmt.Println(">>> MerchantQrisStatusService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MerchantQrisStatusService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MerchantQrisStatusService: Filter")
	defer span.Finish()

	var res models.Response

	list, total, err := service.MerchantQrisStatusRepository.Filter(req)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.TotalData = total
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
