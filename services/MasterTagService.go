package services

import (
	"fmt"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"time"
)

// MasterTagService struct
type MasterTagService struct {
	Ottolog logger.OttologInterface
	MasterTagRepository *db.MasterTagRepository
}

// InitMasterTagService ...
func InitMasterTagService(logs logger.OttologInterface) *MasterTagService {
	return &MasterTagService{
		Ottolog:logs,
		MasterTagRepository: db.InitMasterTagRepository(logs),
	}
}

// Save ...
func (svc *MasterTagService) Save(req dto.ReqMasterTagDto, res *models.Response)  {
	svc.Ottolog.Info("MasterTagService - Save")

	var data dbmodels.MasterTag
	var err error

	data.CreatedAt = time.Now()
	data.CreatedBy = auth.UserLogin.Name
	if req.ID > 0 {
		data, err = svc.MasterTagRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return
		}
	}

	data.Name = req.Name
	data.Code = req.Code
	data.Description = req.Description
	switch req.Status {
	case constants.IN_ACTIVE:
		data.Status = false
		
	case constants.ACTIVE:
		data.Status = true
		
	}
	data.UpdatedAt = time.Now()
	data.UpdatedBy = auth.UserLogin.Name

	if err := svc.MasterTagRepository.Save(&data); err!= nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

}

// Filter ...
func (svc *MasterTagService) Filter(req dto.ReqMasterTagDto, res *models.Response)  {
	svc.Ottolog.Info("MasterTagService - Filter")


	data, total, err := svc.MasterTagRepository.Filter(req)
	if err != nil {

		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND

		svc.Ottolog.Error(fmt.Sprintf("Failed to get data from database: %s", fmt.Sprintf("ERR:%s", err.Error())))
		return
	}

	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.ErrCode = constants.ERR_SUCCESS
	res.Contents = data
	res.TotalData = total

}

// GetAll ...
func (svc *MasterTagService) GetAll(res *models.Response)  {
	svc.Ottolog.Info("MasterTagService - GetAll")


	data, err := svc.MasterTagRepository.GetAll()
	if err != nil {

		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND

		svc.Ottolog.Error(fmt.Sprintf("Failed to get data from database: %s", fmt.Sprintf("ERR:%s", err.Error())))
		return
	}

	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.ErrCode = constants.ERR_SUCCESS
	res.Contents = data

}
