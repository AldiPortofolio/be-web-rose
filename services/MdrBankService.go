package services

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"log"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"time"
)

// MdrBankService struct
type MdrBankService struct {
	General models.GeneralModel
	MdrBankRepository *db.MdrBankRepository
}

// InitMdrBankService ...
func InitMdrBankService(gen models.GeneralModel) *MdrBankService {
	return &MdrBankService{
		General:gen,
		MdrBankRepository: db.InitMdrBankRepository(),
	}
}

// Save ...
func (service *MdrBankService) Save(req dto.ReqMdrDto) models.Response {
	fmt.Println(">> MdrBankService - Save <<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MdrBankService: save",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MdrBankService: save")
	defer span.Finish()

	var res models.Response

	var data dbmodels.MdrBank
	var err error

	data.UpdatedAt = time.Now()
	data.UpdatedBy = auth.UserLogin.Name
	if req.ID > 0  {
		data, err = service.MdrBankRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return res
		}
	}

	data.BankCode = req.BankCode
	data.BankName = req.BankName
	data.DokuBankCode = req.DokuBankCode
	data.Status = req.Status
	data.AcquiringStatus = req.AcquiringStatus
	data.Seq = req.Seq

	if err:=service.MdrBankRepository.Save(&data); err!=nil{
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res

}

func (service *MdrBankService) Filter(req dto.ReqMdrDto) models.Response {
	fmt.Println(">>> MdrBankService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MdrBankService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MdrBankService: Filter")
	defer span.Finish()

	var res models.Response
	list, total, err := service.MdrBankRepository.Filter(req)
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

func (service *MdrBankService) FindAll() models.Response {
	fmt.Println(">>> MdrBankService - FindAll <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MdrBankService: FindAll",
		zap.Any("req", ""))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MdrBankService: FindAll")
	defer span.Finish()

	var res models.Response
	list, err := service.MdrBankRepository.FindAll()
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.Contents = list
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.Contents = list

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}
