package services

import (
	"fmt"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dto"
)

type MerchantBankLoanService struct {
	MerchantBankLoanRepository *db.MerchantBankLoanRepository
	SubMerchantBankLoanRepository *db.SubMerchantBankLoanRepository
}

func InitMerchantBankLoanService() *MerchantBankLoanService {
	return &MerchantBankLoanService{
		MerchantBankLoanRepository: db.InitMerchantBankLoanRepository(),
		SubMerchantBankLoanRepository: db.InitSubMerchantBankLoanRepository(),
	}
}


func (svc *MerchantBankLoanService) Filter(req dto.ReqMerchantBankLoanDto, res *models.Response)  {
	fmt.Println("MerchantBankLoanService - Filter")


	data, total, err := svc.MerchantBankLoanRepository.Filter(req)
	if err != nil {

		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND

		fmt.Println(fmt.Printf("Failed to get data from database: %s", fmt.Sprintf("ERR:%s", err.Error())))
		return
	}

	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.ErrCode = constants.ERR_SUCCESS
	res.Contents = data
	res.TotalData = total

}

func (svc *MerchantBankLoanService) FindSubMerchantBankLoan(req dto.ReqSubMerchantBankloanDto, res *models.Response)  {
	fmt.Println("MerchantBankLoanService - Filter")


	data, err := svc.SubMerchantBankLoanRepository.FindByMasterLoanId(req.MasterBankLoanId)
	if err != nil {

		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND

		fmt.Println("Failed to get data from database: ", fmt.Sprintf("ERR:%s", err.Error()))
		return
	}

	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.ErrCode = constants.ERR_SUCCESS
	res.Contents = data

	
}

