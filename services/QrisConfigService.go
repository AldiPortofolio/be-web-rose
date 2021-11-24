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

type QrisConfigService struct {
	General models.GeneralModel
	QrisConfigRepository *db.QrisConfigRepository
	QrisConfigHistoryRepository *db.QrisConfigHistoryRepository
}

func InitQrisConfigService(gen models.GeneralModel) *QrisConfigService {
	return &QrisConfigService{
		General:gen,
		QrisConfigRepository: db.InitQrisConfigRepository(),
		QrisConfigHistoryRepository: db.InitQrisConfigHistoryRepository(),
	}
}

func (service *QrisConfigService) Filter(req dto.ReqQrisConfigDto) models.Response {
	fmt.Println(">>> QrisConfigService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("QrisConfigService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "QrisConfigService: Filter")
	defer span.Finish()

	var res models.Response
	list, total, err := service.QrisConfigRepository.Filter(req)
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

func (service *QrisConfigService) Save(req dto.ReqQrisConfigDto) models.Response {
	fmt.Println(">>> QrisConfigService - save <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("QrisConfigService: save",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "QrisConfigService: save")
	defer span.Finish()

	var res models.Response

	var data dbmodels.QrisConfig
	var err error

	data.CreatedAt = time.Now()
	data.CreatedBy = auth.UserLogin.Name
	if req.ID > 0  {
		data, err = service.QrisConfigRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return res
		}
	}

	data.InstitutionID = req.InstitutionID
	data.IssuerName = req.IssuerName
	data.TransactionType = req.TransactionType
	data.Status = req.Status
	data.UpdatedAt = time.Now()
	data.UpdatedBy = auth.UserLogin.Name

	if err:=service.QrisConfigRepository.Save(&data); err!=nil{
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}
	service.InsertLog(data) // insert log

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}

func (service *QrisConfigService) InsertLog(req dbmodels.QrisConfig)  {
	data := dbmodels.QrisConfigHistory{
		QrisConfigID: req.ID,
		InstitutionID: req.InstitutionID,
		IssuerName: req.IssuerName,
		TransactionType: req.TransactionType,
		Status: req.Status,
		CreatedAt: time.Now(),
		CreatedBy: auth.UserLogin.Name,
	}
	if err := service.QrisConfigHistoryRepository.Save(data); err!= nil{
		log.Println("err save to qris config history", err)
	}

}

