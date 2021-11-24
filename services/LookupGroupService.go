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

// LookupGroupService struct ...
type LookupGroupService struct {
	Ottolog logger.OttologInterface
	LookupGroupRepository *db.LookupGroupRepository
}

// InitLookupGroupService ...
func InitLookupGroupService(logs logger.OttologInterface) *LookupGroupService {
	return &LookupGroupService{
		Ottolog: logs,
		LookupGroupRepository: db.InitLookupGroupRepository(logs),
	}
}

// Save ...
func (svc *LookupGroupService) Save(req dto.ReqLookupGroupDto, res *models.Response)  {
	svc.Ottolog.Info("LookupGroupService - Save")

	var data dbmodels.LookupGroup
	var err error

	data.CreatedAt = time.Now()
	data.CreatedBy = auth.UserLogin.Name

	if req.ID > 0 {
		data, err = svc.LookupGroupRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return
		}
	}

	data.UpdatedAt = time.Now()
	data.UpdatedBy = auth.UserLogin.Name
	data.Name = req.Name

	if err:=svc.LookupGroupRepository.Save(&data); err!=nil{
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return
	}


	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

}

// Filter ...
func (svc *LookupGroupService) Filter(req dto.ReqLookupGroupDto, res *models.Response)  {
	svc.Ottolog.Info("LookupGroupService - Filter")


	data, total, err := svc.LookupGroupRepository.Filter(req)
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

// FindAll ...
func (svc *LookupGroupService) FindAll( res *models.Response)  {
	svc.Ottolog.Info("LookupGroupService - FindAll")


	data, err := svc.LookupGroupRepository.FindAll()
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