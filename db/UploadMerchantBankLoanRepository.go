package db

import (
	"errors"
	"fmt"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type UploadMerchantBankLoanRepository struct {

}

func InitUploadMerchantBankLoanRepository() *UploadMerchantBankLoanRepository {
	return &UploadMerchantBankLoanRepository{}
}

func (repo *UploadMerchantBankLoanRepository) Save(uploadMerchant *dbmodels.UploadMerchantBankLoan) (error) {
	db := GetDbCon()
	if err := db.Save(&uploadMerchant).Error; err != nil {
		return  errors.New("Gagal Insert Upload Merchant Bank Loan " + err.Error())
	}
	return nil
}


func (repo *UploadMerchantBankLoanRepository) GetDataUploadMerchant(req dto.ReqUploadMerchantBankLoanDto) ([]dbmodels.UploadMerchantBankLoan, int, error) {
	db := GetDbCon()

	page := req.Page
	limit := req.Limit

	var uploadMerchant []dbmodels.UploadMerchantBankLoan
	var total int


	err := db.Limit(limit).Offset((page-1) * limit).Order("id desc").Order("id").Find(&uploadMerchant).Limit(-1).Offset(0).Count(&total).Error // query

	if err != nil {
		fmt.Println("<<< Error get data upload merchant Bank Loan >>>", err)
		return uploadMerchant, 0, err
	}

	return uploadMerchant, total, nil
}