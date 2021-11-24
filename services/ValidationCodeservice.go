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
	"rose-be-go/utils"
	"time"
)

type ValidationCodeService struct {
	Ottolog logger.OttologInterface
	ValidationRepository *db.ValidationRepository
}

func InitValidationCodeService(logs logger.OttologInterface) *ValidationCodeService {
	return &ValidationCodeService{
		Ottolog:logs,
		ValidationRepository: db.InitValidationRepository(logs),
	}
}

func (svc *ValidationCodeService) Save(req dto.ReqValidationCodeDto, res *models.Response)  {
	svc.Ottolog.Info("ValidationCodeService - Save")

	var data dbmodels.ValidationCode
	var err error

	data.CreatedAt = time.Now()
	data.CreatedBy = auth.UserLogin.Name

	if req.ID >0 {
		data, err = svc.ValidationRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return
		}
	}

	data.AppID = req.AppID
	data.UserCategoryCode = req.UserCategoryCode
	data.ValidationCode = req.ValidationCode
	data.ValidFrom = utils.ConverDateStringYYYYMMDDToTime(req.ValidFrom)
	data.ValidTo = utils.ConverDateStringYYYYMMDDToTime(req.ValidTo).Add(time.Hour * 24)
	data.UpdatedAt = time.Now()
	data.UpdatedBy = auth.UserLogin.Name

	if err := svc.ValidationRepository.Save(&data); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG


}


func (svc *ValidationCodeService) Filter(req dto.ReqValidationCodeDto, res *models.Response)  {
	svc.Ottolog.Info("ValidationCodeService - Filter")


	data, total, err := svc.ValidationRepository.Filter(req)
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