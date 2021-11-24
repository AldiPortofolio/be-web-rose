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

// LoanProductMaintenanceService struct 
type LoanProductMaintenanceService struct {
	Ottolog logger.OttologInterface
	LoanProductMaintenanceRepository *db.LoanProductMaintenanceRepository
	BankListRepository *db.BankListRepository
}

// InitLoanProductMaintenanceService ...
func InitLoanProductMaintenanceService(logs logger.OttologInterface)  *LoanProductMaintenanceService{
	return &LoanProductMaintenanceService{
		Ottolog:logs,
		LoanProductMaintenanceRepository: db.InitLoanProductMaintenanceRepository(logs),
		BankListRepository: db.InitBankListRepository(logs),
	}
}

// Save ...
func (svc *LoanProductMaintenanceService) Save(req dto.ReqSaveLoanProductMaintenanceDto, res *models.Response)  {
	svc.Ottolog.Info("LoanProductMaintenanceService - Save")

	var err error

	var data dbmodels.LoanProductMaintenance
	data.CreatedAt = time.Now()
	data.CreatedBy = auth.UserLogin.Name

	if req.ID >0 {
		data, err = svc.LoanProductMaintenanceRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return
		}
	}
	bankData,err := svc.BankListRepository.FindByCode(req.BankCode)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		return
	}

	data.BankName = bankData.FullName
	data.BankCode = req.BankCode
	data.LoanProductName = req.LoanProductName
	data.LoanProductCode = req.LoanProductCode
	data.AdminFeeValue = req.AdminFeeValue
	data.Description = req.Description
	data.AdminFeeType = req.AdminFeeType
	data.Status = req.Status
	data.UpdatedAt = time.Now()
	data.UpdatedBy = auth.UserLogin.Name

	if err := svc.LoanProductMaintenanceRepository.Save(&data); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG


}

// Filter ...
func (svc *LoanProductMaintenanceService) Filter(req dto.ReqLoanProductMaintenanceDto, res *models.Response)  {
	svc.Ottolog.Info("LoanProductMaintenanceService - Filter")


	data, total, err := svc.LoanProductMaintenanceRepository.Filter(req)
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