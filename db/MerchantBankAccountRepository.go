package db

import (
	"errors"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type MerchantBankAccountRepository struct {
	Ottolog logger.OttologInterface
}

func InitMerchantBankAccountRepository(logs logger.OttologInterface) *MerchantBankAccountRepository {
	return &MerchantBankAccountRepository{
		Ottolog:logs,
	}
}


func (r *MerchantBankAccountRepository) Save(account *dbmodels.MerchantBankAccount) error {
	db := GetDbCon()

	if err := db.Save(&account).Error; err != nil {
		r.Ottolog.Error("failed Save to MerchantBank Account " + err.Error())
		return errors.New("failed Save to MerchantBank Account " + err.Error())
	}

	return nil

}

func (r *MerchantBankAccountRepository) GetDataByMid(mid string) ([]dto.ResMerchantBankAccount, error) {

	db := GetDbCon()
	var res []dto.ResMerchantBankAccount
	err := db.Table("merchant_bank_account a").
		Select("a.*, b.full_name, b.short_name, b.url_image", ).
		Joins("LEFT JOIN bank_list b on b.code = a.bank_code").
		Where("a.mid = ?", mid).Order("updated_at desc").Scan(&res).Error
	return res, err

}

func (r *MerchantBankAccountRepository) FindByID(id int64) (dbmodels.MerchantBankAccount, error) {

	db := GetDbCon()
	var res dbmodels.MerchantBankAccount
	err := db.Where("id = ?", id).First(&res).Error
	return res, err


}

func (r *MerchantBankAccountRepository) FilterApproval(req dto.ReqMerchantBankAccountDto) ([]dto.ResMerchantBankAccount, int, error) {
	db := GetDbCon()

	var res []dto.ResMerchantBankAccount
	limit := req.Limit
	page := req.Page
	var total int

	db = db.Table("merchant_bank_account a").
		Select("a.*, b.full_name, b.short_name, b.url_image", ).
		Joins("LEFT JOIN bank_list b on b.code = a.bank_code")

	if req.Status != "" {
		db = db.Where("a.status = ?", req.Status)
	}

	if req.Mid != "" {
		db = db.Where("a.mid ilike ?", "%" + req.Mid +"%")

	}

	if req.BankCode != "" {
		db = db.Where("a.bank_code ilike ?", "%" + req.BankCode +"%")

	}

	if req.AccountNumber != "" {
		db = db.Where("a.account_number ilike ?", "%" + req.AccountNumber +"%")

	}

	if req.AccountName != "" {
		db = db.Where("a.account_name ilike ?", "%" + req.AccountName +"%")

	}

	err := db.Limit(limit).Offset((page-1)*limit).Order("a.updated_at asc").Scan(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		r.Ottolog.Error("error get data Approval Merchant Bank Account "+ err.Error())
		return res, 0, err
	}

	return res, total, nil
}

