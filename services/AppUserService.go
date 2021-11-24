package services

import (
	"fmt"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	"rose-be-go/constants"
	"time"
)

// BankListService struct
type AppUserService struct {
	Ottolog logger.OttologInterface
	AppUserRepository *db.AppUserRepository
}

// InitAppUserService ...
func InitAppUserService(logs logger.OttologInterface)*AppUserService  {
	return &AppUserService{
		Ottolog:logs,
		AppUserRepository: db.InitAppUserRepository(logs),
	}
}

// Save ...
func (svc *AppUserService) Save(req dto.ReqAppUserDto, res *models.Response)  {
	svc.Ottolog.Info("AppUserService - Save")
	
	data, err := svc.AppUserRepository.FindByUserName(req.Username)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		return
	}

	if data.Password != req.OldPass {
		res.ErrCode = constants.EC_INVALID_PASSWORD
		res.ErrDesc = constants.EC_INVALID_PASSWORD_DESC
		return
	}

	data.Password = req.NewPass
	data.PwdExpiredDate = time.Now().AddDate(0,3,0)
	fmt.Println(data.PwdExpiredDate, "------ date after 3 month")

	if err := svc.AppUserRepository.Save(&data); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

}
