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

// LogUpgradeFdsService struct
type LogUpgradeFdsService struct {
	Ottolog logger.OttologInterface
	LogUpgradeFdsRepository *db.LogUpgradeFdsRepository
}

// InitLogUpgradeFdsService ...
func InitLogUpgradeFdsService(logs logger.OttologInterface) *LogUpgradeFdsService {
	return &LogUpgradeFdsService{
		Ottolog:logs,
		LogUpgradeFdsRepository: db.InitLogUpgradeFdsRepository(logs),
	}
}

// Filter ...
func (svc *LogUpgradeFdsService) Filter(req dto.ReqLogUpgradeFdsDto, res *models.Response)  {
	svc.Ottolog.Info("LogUpgradeFdsRepository - Filter")


	data, total, err := svc.LogUpgradeFdsRepository.Filter(req)
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

// Save ...
func (svc *LogUpgradeFdsService) Save(req dbmodels.LogUpgradeFds, res *models.Response)  {
	svc.Ottolog.Info("LogUpgradeFdsService - Save")

	var data dbmodels.LogUpgradeFds
	var err error


	data, err = svc.LogUpgradeFdsRepository.FindByID(req.ID)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		return
	}

	data.RetryAt = time.Now()
	data.RetryBy = auth.UserLogin.Name
	data.Status = constants.RETRY

	if err := svc.LogUpgradeFdsRepository.Save(&data); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

}