package services

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"time"
)

// MasterServicesService ..
type MasterServicesService struct {
	General                 models.GeneralModel
	MasterServiceRepository *db.MasterServiceRepository
}

// InitMasterServicesService ..
func InitMasterServicesService(gen models.GeneralModel) *MasterServicesService {
	return &MasterServicesService{
		General:                 gen,
		MasterServiceRepository: db.InitMasterServiceRepository(),
	}
}

// Save ..
func (service *MasterServicesService) Save(req dto.ReqMasterServiceDto) models.Response {
	fmt.Println(">> MasterServicesService - Save <<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MasterServicesService: Save",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MasterServicesService: Save")
	defer span.Finish()

	var res models.Response

	var data dbmodels.MasterService
	var err error

	if req.ID > 0 {
		data, err = service.MasterServiceRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return res
		}
	}

	data.Name = req.Name
	data.UpdatedAt = time.Now()
	data.UpdatedBy = auth.UserLogin.Name

	if err := service.MasterServiceRepository.Save(&data); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}

// FindAll ..
func (service *MasterServicesService) FindAll() models.Response {
	fmt.Println(">>> MasterServicesService - FindAll <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MasterServicesService: FindAll",
		zap.Any("req", ""))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MasterServicesService: FindAll")
	defer span.Finish()

	var res models.Response
	list, total, err := service.MasterServiceRepository.FindAll()
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
