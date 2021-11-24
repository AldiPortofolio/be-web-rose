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

type ValidationCodeMasterTagService struct {
	Ottolog logger.OttologInterface
	ValidationCodeMasterTagRepository *db.ValidationCodeMasterTagRepository
}

func InitValidationCodeMasterTagService(logs logger.OttologInterface)  *ValidationCodeMasterTagService{
	return &ValidationCodeMasterTagService{
		Ottolog:logs,
		ValidationCodeMasterTagRepository: db.InitValidationCodeMasterTagRepository(logs),
	}
}

func (svc *ValidationCodeMasterTagService) Save(req dto.ReqSaveValidationCodeMasterTagDto, res *models.Response)  {
	svc.Ottolog.Info("ValidationCodeMasterTagService - Save")

	var err error

	var data dbmodels.ValidationCodeMasterTag
	data.CreatedAt = time.Now()
	data.CreatedBy = auth.UserLogin.Name

	if req.ID >0 {
		data, err = svc.ValidationCodeMasterTagRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return
		}
	}
	data.ValidationCodeID = req.ValidationCodeID
	data.MasterTagID = req.MasterTagID
	data.UpdatedAt = time.Now()
	data.UpdatedBy = auth.UserLogin.Name

	if err := svc.ValidationCodeMasterTagRepository.Save(&data); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG


}



func (svc *ValidationCodeMasterTagService) Filter(req dto.ReqValidationCodeMasterTagDto, res *models.Response)  {
	svc.Ottolog.Info("ValidationCodeMasterTagService - Filter")


	data, total, err := svc.ValidationCodeMasterTagRepository.Filter(req)
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