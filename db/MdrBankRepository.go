package db

import (
	"rose-be-go/constants/status_mdr_bank"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type MdrBankRepository struct {

}

func InitMdrBankRepository() *MdrBankRepository {
	return &MdrBankRepository{}
}

func (repo *MdrBankRepository) Save(req *dbmodels.MdrBank) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

func (repo *MdrBankRepository) FindByID(id int64) (dbmodels.MdrBank, error) {
	db := GetDbCon()

	var data dbmodels.MdrBank

	err := db.Where(dbmodels.MdrBank{
		ID: id,
	}).First(&data).Error

	return data, err
}

func (repo *MdrBankRepository) Filter(req dto.ReqMdrDto) ([]dbmodels.MdrBank, int, error) {
	db := GetDbCon()
	var res []dbmodels.MdrBank
	var total int
	page := req.Page
	limit := req.Limit

	db = db.Where("status = ?", status_mdr_bank.ACTIVE)

	if req.BankName != "" {
		db = db.Where("bank_name ilike ?", "%" + req.BankName + "%")
	}

	if req.BankCode != "" {
		db = db.Where("bank_code ilike ?", "%" + req.BankCode + "%")
	}

	err := db.Limit(limit).Offset((page-1)*limit).Order("seq asc").Find(&res).Limit(-1).Offset(0).Count(&total).Error



	return res, total, err

}


func (repo *MdrBankRepository) FindAll() ([]dbmodels.MdrBank, error) {
	db := GetDbCon()
	var res []dbmodels.MdrBank

	db = db.Where("status = ?", status_mdr_bank.ACTIVE)


	err := db.Order("seq asc").Find(&res).Error



	return res, err

}

func (repo *MdrBankRepository) Get(data dbmodels.MdrBank) (dbmodels.MdrBank, error) {
	db := GetDbCon()
	var res dbmodels.MdrBank

	err := db.Where(&data).First(&res).Error

	return res, err
}
