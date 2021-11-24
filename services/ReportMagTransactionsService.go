package services

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/opentracing/opentracing-go"
	"github.com/unknwon/com"
	"go.uber.org/zap"
	"io/ioutil"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	"strings"
)

type ReportMagTransactionsService struct {
	General models.GeneralModel
}

func InitReportMagTransactionsService(gen models.GeneralModel) *ReportMagTransactionsService {
	return &ReportMagTransactionsService{
		General: gen,
	}
}

// GetAllReportMagTransactions ...
func (service *ReportMagTransactionsService) GetAllReportMagTransactions(req dto.ReqReportMagTransactionsDto) models.Response {
	fmt.Println(">>> ReportMagTransactionsService - GetAllReportMagTransactions <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("ReportMagTransactionsService: GetAllReportMagTransactions", zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "ReportMagTransactionsService: GetAllReportMagTransactions")
	defer span.Finish()

	var res models.Response

	data, total, err := db.InitReportMagTransactionsRepository().GetAllDataMagTransaction(req)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.TotalData = total
	res.Contents = data

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}

// Export ...
func (service *ReportMagTransactionsService) Export(req dto.ReqReportMagTransactionsDto) models.Response {
	fmt.Println(">>> ReportMagTransactionsService - Export <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("ReportMagTransactionsService: Export", zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "ReportMagTransactionsService: Export")
	defer span.Finish()

	var res models.Response

	data, total, err := db.InitReportMagTransactionsRepository().GetAllDataMagTransaction(req)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		return res
	}

	file := db.InitReportMagTransactionsRepository().Export(data)
	csvContent, err := gocsv.MarshalString(&file) // Get all clients as CSV string
	if err != nil {
		res.ErrCode = constants.EC_FAIL_CSV_FILE
		res.ErrDesc = constants.EC_FAIL_CSV_FILE_DESC + com.ToStr(err)
		return res
	}

	r := strings.NewReader(csvContent)
	fileBase64, err := ioutil.ReadAll(r)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_CSV_FILE
		res.ErrDesc = constants.EC_FAIL_CSV_FILE_DESC + com.ToStr(err)
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.TotalData = total
	res.Contents = fileBase64

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}