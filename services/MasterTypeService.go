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

// MasterTypeService ..
type MasterTypeService struct {
	General              models.GeneralModel
	MasterTypeRepository *db.MasterTypeRepository
}

// InitMasterTypeService ..
func InitMasterTypeService(gen models.GeneralModel) *MasterTypeService {
	return &MasterTypeService{
		General:              gen,
		MasterTypeRepository: db.InitMasterTypeRepository(),
	}
}

// Save ..
func (service *MasterTypeService) Save(req dto.ReqMasterTypeDto) models.Response {
	fmt.Println(">> MasterTypeService - Save <<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MasterTypeService: Save",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MasterTypeService: Save")
	defer span.Finish()

	var res models.Response

	var data dbmodels.MasterType
	var err error

	if req.ID > 0 {
		data, err = service.MasterTypeRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return res
		}
	}

	data.Name = req.Name
	data.UpdatedAt = time.Now()
	data.UpdatedBy = auth.UserLogin.Name

	if err := service.MasterTypeRepository.Save(&data); err != nil {
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
func (service *MasterTypeService) FindAll() models.Response {
	fmt.Println(">>> MasterTypeService - FindAll <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MasterTypeService: FindAll",
		zap.Any("req", ""))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MasterTypeService: FindAll")
	defer span.Finish()

	var res models.Response
	list, total, err := service.MasterTypeRepository.FindAll()
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
