package services

import (
	"encoding/json"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"log"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/hosts/ottomart"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/ottomartmodels"
	"time"
)

// ClearSessionService struct
type ClearSessionService struct {
	General models.GeneralModel
	OttomartHost *ottomart.OttomartHost
	 *db.HistoryClearSessionRepository

}

// InitClearSessionService ...
func InitClearSessionService(gen models.GeneralModel) *ClearSessionService{
	return &ClearSessionService{
		General: gen,
		OttomartHost: ottomart.InitOttomartHost(),
		HistoryClearSessionRepository: db.InitHistoryClearSessionRepository(),

	}
}

// Clear ...
func (service *ClearSessionService) Clear() models.Response  {
	fmt.Println(">>> ClearSessionService - Clear <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("ClearSessionService: Clear",
		zap.Any("req", ""))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "ClearSessionService: Clear")
	defer span.Finish()

	var res models.Response

	data, err := service.OttomartHost.Send("", constants.OttomartClearSession)
	if err != nil {
		log.Println("err -> ", err)
		res.ErrCode = constants.EC_FAIL_SEND_TO_HOST
		res.ErrDesc = constants.EC_FAIL_SEND_TO_HOST_DESC
		return res
	}

	resOttomart := ottomartmodels.ClearSessionRes{}
	if err = json.Unmarshal(data, &resOttomart); err != nil {
		go sugarLogger.Error(fmt.Sprintf("Failed to unmarshal json response from 'ottomart' host: %v", err),
			zap.String("SpanID:", service.General.SpanId))

		res.ErrCode = constants.EC_FAIL_SEND_TO_HOST
		res.ErrDesc = constants.EC_FAIL_SEND_TO_HOST_DESC
		return res
	}

	service.InsertLog()


	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}

// GetLastUpdated ...
func (service *ClearSessionService) GetLastUpdated() models.Response  {
	fmt.Println(">>> ClearSessionService - GetLastUpdated <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("ClearSessionService: GetLastUpdated",
		zap.Any("req", ""))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "ClearSessionService: GetLastUpdated")
	defer span.Finish()

	var res models.Response

	data, err := service.HistoryClearSessionRepository.GetLastUpdated()
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.Contents = data

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}

// InsertLog ...
func (service *ClearSessionService) InsertLog() {
	fmt.Println("<< ClearSessionService - Insert log to DB >> ")
	req:=dbmodels.HistoryClearSession{
		CreatedBy: auth.UserLogin.Name,
		CreatedAt:time.Now(),
	}
	log.Println(req)

	service.HistoryClearSessionRepository.Save(req)

}


