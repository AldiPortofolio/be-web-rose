package services

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"log"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dto"
)

// MonitoringActivationFDSService struct
type MonitoringActivationFDSService struct {
	General 							models.GeneralModel
	MonitoringActivationFDSRepository 	*db.MonitoringActivationFDSRepository
}

// InitMonitoringActivationFDSService ...
func InitMonitoringActivationFDSService(gen models.GeneralModel) *MonitoringActivationFDSService {
	return &MonitoringActivationFDSService{
		General:gen,
		MonitoringActivationFDSRepository: db.InitMonitoringActivationFDSRepository(),
	}
}

// Filter ...
func (service *MonitoringActivationFDSService) Filter(req dto.ReqMonitoringActivationFDSDto) models.Response {
	fmt.Println(">>> MonitoringActivationFDSService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MonitoringActivationFDSService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MonitoringActivationFDSService: Filter")
	defer span.Finish()

	var res models.Response
	list, total, err := service.MonitoringActivationFDSRepository.Filter(req)
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