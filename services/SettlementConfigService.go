package services

import (
	"fmt"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"time"

	"ottodigital.id/library/logger/v2"
)

// SettlementConfigService struct
type SettlementConfigService struct {
	Ottolog logger.OttologInterface
	SettlementConfigRepository *db.SettlementConfigRepository
}

// InitSettlementConfigService ...
func InitSettlementConfigService(logs logger.OttologInterface)*SettlementConfigService  {
	return &SettlementConfigService{
		Ottolog:logs,
		SettlementConfigRepository: db.InitSettlementConfigRepository(logs),
	}
}

// Save ...
func (svc *SettlementConfigService) Save(req dto.ReqSettlementConfigDto, res *models.Response)  {
	svc.Ottolog.Info("SettlementConfigService - Save")

	var data dbmodels.SettlementConfig
	var oldData dbmodels.SettlementConfig
	var responseContents = [2]dbmodels.SettlementConfig{}
	var err error
	data.CreatedAt = time.Now()
	
	if req.Code != "" {
		oldData, err = svc.SettlementConfigRepository.FindByBankCode(req.Code)
		if err == nil {
			oldData.Status = "Non Active"
			oldData.UpdatedAt = time.Now()
			oldData.UpdatedBy = auth.UserLogin.Name

			err = svc.SettlementConfigRepository.Save(&oldData)
			if err != nil {
				res.ErrCode = constants.EC_FAIL_SAVE
				res.ErrDesc = constants.EC_FAIL_SAVE_DESC
	
			}
			
		}
	}
	data.BankCode = req.Code
	data.SettlementType  = req.SettlementType
	data.Status  		= req.Status
	data.UpdatedAt = time.Now()
	data.UpdatedBy = auth.UserLogin.Name

	if err := svc.SettlementConfigRepository.Save(&data); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return
	}

	responseContents[0] = oldData
	responseContents[1] = data
	
	res.Contents = responseContents
	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

}

// Filter ...
func (svc *SettlementConfigService) Filter(req dto.ReqSettlementConfigDto, res *models.Response)  {
	svc.Ottolog.Info("SettlementConfigService - Filter")


	data, total, err := svc.SettlementConfigRepository.Filter(req)
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