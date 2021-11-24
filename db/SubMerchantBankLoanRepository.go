package db

import "rose-be-go/models/dbmodels"

type SubMerchantBankLoanRepository struct {
	
}

func InitSubMerchantBankLoanRepository() *SubMerchantBankLoanRepository {
	return &SubMerchantBankLoanRepository{}
}

func (r *SubMerchantBankLoanRepository)FindByMasterLoanId(masterLoanID string) (data []dbmodels.SubMerchantBankLoan, err error)  {

	db := GetDbCon()

	err = db.Where("master_bank_loan_id = ?", masterLoanID).Find(&data).Error
	return data, err
}
