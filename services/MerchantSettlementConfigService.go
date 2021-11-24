package services

import (
	"fmt"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"strconv"
	"time"

	"ottodigital.id/library/logger/v2"
)

// BankListService struct
type MerchantSettlementConfigService struct {
	Ottolog logger.OttologInterface
	MerchantSettlementConfigRepository *db.MerchantSettlementConfigRepository
	MerchantRepository *db.MerchantRepository
}

// InitAppUserService ...
func InitMerchantSettlementConfigService(logs logger.OttologInterface) *MerchantSettlementConfigService  {
	return &MerchantSettlementConfigService{
		Ottolog:logs,
		MerchantSettlementConfigRepository: db.InitMerchantSettlementConfigRepository(logs),
		MerchantRepository: db.InitMerchantRepository(),
	}
}

// Save ...
func (svc *MerchantSettlementConfigService) Save(req dbmodels.MerchantSettlementConfig, res *models.Response)  {
	svc.Ottolog.Info("MerchantSettlementConfigService - Save")
	id , _ := strconv.Atoi(req.Mid)

	merchant, err := svc.MerchantRepository.FindById(id)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		return
	}

	data, err := svc.MerchantSettlementConfigRepository.FindByMid(merchant.MerchantOutletID)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		fmt.Println("err find data exist : ", err.Error())
	}
	fmt.Println(data.Mid, " == ", merchant.MerchantOutletID)
	if data.Mid == merchant.MerchantOutletID {
		fmt.Println("masuk update data")
		data.NamaBankTujuanSettlement = req.NamaBankTujuanSettlement
		data.NamaPemilikRekening = req.NamaPemilikRekening
		data.NoRekeningToko = req.NoRekeningToko
		data.TipeRekening = req.TipeRekening
		data.ReportSettlementConfigName = req.ReportSettlementConfigName
		data.SettlementExecutionConfigName = req.SettlementExecutionConfigName
		data.SftpHost = req.SftpHost
		data.SftpUser = req.SftpUser
		data.SftpPassword = req.SftpPassword
		data.Email = req.Email
		data.Status = req.Status
		data.UpdatedAt = time.Now()
		if err := svc.MerchantSettlementConfigRepository.Update(&data); err != nil {
			res.ErrCode = constants.EC_FAIL_SAVE
			res.ErrDesc = constants.EC_FAIL_SAVE_DESC
			return
		}

	} else {
		fmt.Println("masuk save data baru")
		data = req
		data.Mid = merchant.MerchantOutletID
		data.CreatedAt = time.Now()
		data.UpdatedAt = time.Now()
		if err := svc.MerchantSettlementConfigRepository.Save(&data); err != nil {
			res.ErrCode = constants.EC_FAIL_SAVE
			res.ErrDesc = constants.EC_FAIL_SAVE_DESC
			return
		}
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

}
