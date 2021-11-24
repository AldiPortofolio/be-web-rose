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

// DepositBankService struct
type DepositBankService struct {
	General 	models.GeneralModel
}

// InitDepositBankService ...
func InitDepositBankService(gen models.GeneralModel) *DepositBankService {
	return &DepositBankService{
		General: gen,
	}
}

// GetAllInfo ...
func (service *DepositBankService) GetAllInfo(req dto.ReqDepositBank) models.Response {
	fmt.Println(">>> DepositBankService - GetAllInfo <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("DepositBankService: GetAllInfo",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "DepositBankService: GetAllInfo")
	defer span.Finish()

	var res models.Response

	data, total, err := db.InitDepositBankRepository().GetAllInfo(req)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.Contents = data
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.TotalData = total
	res.Contents = data

	sugarLogger.Info("Response", zap.Any("res", res))

	return res

}
