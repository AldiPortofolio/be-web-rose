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

type MdrTenorService struct {
	General models.GeneralModel
	MdrTenorRepository *db.MdrTenorRepository
}

func InitMdrTenorService(gen models.GeneralModel) *MdrTenorService {
	return &MdrTenorService{
		General:gen,
		MdrTenorRepository: db.InitMdrTenorRepository(),
	}
}

func (service *MdrTenorService) Save(req dto.ReqMdrTenorDto) models.Response {
	fmt.Println(">> MdrTenorService - Save <<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MdrTenorService: save",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MdrTenorService: save")
	defer span.Finish()

	var res models.Response

	var data dbmodels.MdrTenor
	var err error

	data.UpdatedAt = time.Now()
	data.UpdatedBy = auth.UserLogin.Name
	if req.ID > 0  {
		data, err = service.MdrTenorRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return res
		}
	}

	data.TenorCode = req.TenorCode
	data.DokuTenorCode = req.DokuTenorCode
	data.TenorName = req.TenorName
	data.Status = req.Status
	data.Seq = req.Seq


	if err:=service.MdrTenorRepository.Save(&data); err!=nil{
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res

}

func (service *MdrTenorService) Filter(req dto.ReqMdrTenorDto) models.Response {
	fmt.Println(">>> MdrTenorService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MdrTenorService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MdrTenorService: Filter")
	defer span.Finish()

	var res models.Response
	list, total, err := service.MdrTenorRepository.Filter(req)
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


func (service *MdrTenorService) FindAll() models.Response {
	fmt.Println(">>> MdrTenorService - FindAll <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MdrTenorService: FindAll",
		zap.Any("req", ""))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MdrTenorService: FindAll")
	defer span.Finish()

	var res models.Response
	list, err := service.MdrTenorRepository.FindAll()
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