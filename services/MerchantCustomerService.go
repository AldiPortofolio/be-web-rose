package services

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	"strings"
	"github.com/unknwon/com"
)

// MerchantCustomerService ..
type MerchantCustomerService struct {
	General                    models.GeneralModel
	MerchantCustomerRepository *db.MerchantCustomerRepository
}

// InitMerchantCustomerService ..
func InitMerchantCustomerService(gen models.GeneralModel) *MerchantCustomerService {
	return &MerchantCustomerService{
		General:                    gen,
		MerchantCustomerRepository: db.InitMerchantCustomerRepository(),
	}
}

// Filter ..
func (service *MerchantCustomerService) Filter(req dto.ReqMerchantCustomerDto) models.Response {
	fmt.Println(">>> MerchantCustomerService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MerchantCustomerService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MerchantCustomerService: Filter")
	defer span.Finish()

	var res models.Response
	list, total, err := service.MerchantCustomerRepository.Filter(req)
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

// FindAll ..
func (service *MerchantCustomerService) FindAll() models.Response {
	fmt.Println(">>> MerchantCustomerService - FindAll <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MerchantCustomerService: FindAll",
		zap.Any("req", ""))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MerchantCustomerService: FindAll")
	defer span.Finish()

	var res models.Response
	list, total, err := service.MerchantCustomerRepository.FindAll()
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
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

// Export ..
func (service *MerchantCustomerService) Export(req dto.ReqMerchantCustomerDto) models.Response {
	fmt.Println(">>> MerchantCustomerService - Export <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MerchantCustomerService: Export",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MerchantCustomerService: Export")
	defer span.Finish()

	var res models.Response
	list, total, err := service.MerchantCustomerRepository.Filter(req)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.TotalData = total
		res.Contents = list
		return res
	}

	file := service.MerchantCustomerRepository.Export(list)
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
