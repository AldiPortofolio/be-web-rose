package services

import (
	"fmt"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/auth"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"rose-be-go/constants"
	"time"
)

// BankListService struct
type BankListService struct {
	Ottolog logger.OttologInterface
	BankListRepository *db.BankListRepository
	SettlementConfigRepository *db.SettlementConfigRepository
}

// InitBankListService ...
func InitBankListService(logs logger.OttologInterface)*BankListService  {
	return &BankListService{
		Ottolog:logs,
		BankListRepository: db.InitBankListRepository(logs),
		SettlementConfigRepository: db.InitSettlementConfigRepository(logs),
	}
}

// Save ...
func (svc *BankListService) Save(req dto.ReqBankListDto, res *models.Response)  {
	svc.Ottolog.Info("BankListService - Save")

	var data dbmodels.BankList
	var err error

	data.CreatedAt = time.Now()
	data.CreatedBy = auth.UserLogin.Name
	if req.ID >0 {
		data, err = svc.BankListRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return
		}
	}
	data.ShortName = req.ShortName
	data.FullName  = req.FullName
	data.Code  		= req.Code
	data.Seq 		= req.Seq
	data.UrlImage 	= req.UrlImage
	data.Status		= req.Status
	data.UpdatedAt = time.Now()
	data.UpdatedBy = auth.UserLogin.Name

	if err := svc.BankListRepository.Save(&data); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return
	}

	if data.ID == 0 {
		var settlementConfig dbmodels.SettlementConfig
		settlementConfig.BankCode = req.Code
		settlementConfig.SettlementType = req.SettlementFeeConfig
		settlementConfig.Status = "Active"
		settlementConfig.UpdatedBy = auth.UserLogin.Name
		
		if err := svc.SettlementConfigRepository.Save(&settlementConfig); err != nil {
			res.ErrCode = constants.EC_FAIL_SAVE
			res.ErrDesc = constants.EC_FAIL_SAVE_DESC
			return
		}
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

}

// Filter ...
func (svc *BankListService) Filter(req dto.ReqBankListDto, res *models.Response)  {
	svc.Ottolog.Info("BankListService - Filter")


	data, total, err := svc.BankListRepository.Filter(req)
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