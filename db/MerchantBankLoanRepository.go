package db

import (
	"fmt"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type MerchantBankLoanRepository struct {

}

func InitMerchantBankLoanRepository()*MerchantBankLoanRepository  {
	return &MerchantBankLoanRepository{}
}

func (r *MerchantBankLoanRepository)Save(loan *dbmodels.MerchantBankLoan) error {
	db:=GetDbCon()
	return db.Save(&loan).Error
}

func (r *MerchantBankLoanRepository) FindByMidAndCode(mid, code string) (data dbmodels.MerchantBankLoan, err error)  {

	db:= GetDbCon()
	err =db.Where("mid = ?", mid).Where("bank_code =?", code).First(&data).Error
	return data, err
}

func (r *MerchantBankLoanRepository)Filter(req dto.ReqMerchantBankLoanDto) ([]dbmodels.MerchantBankLoan, int, error) {
	db := GetDbCon()

	var res []dbmodels.MerchantBankLoan
	limit := req.Limit
	page := req.Page
	var total int

	if req.Mid != "" {
		db = db.Where("mid ilike ?", "%" + req.Mid +"%")
	}

	if req.StoreName != "" {
		db = db.Where("store_name ilike ?", "%" + req.StoreName +"%")
	}

	if req.Status !="" {
		db = db.Where("store_name ilike =", req.Status)

	}
	if req.PhoneNumber != "" {
		db = db.Where("phone_number ilike ?", "%" + req.PhoneNumber +"%")
	}

	err := db.Limit(limit).Offset((page-1)*limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		fmt.Println("error get data merchant bank loan"+ err.Error())
		return res, 0, err
	}

	return res, total, nil

}