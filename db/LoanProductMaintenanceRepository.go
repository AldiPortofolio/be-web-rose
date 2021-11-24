package db

import (
	"ottodigital.id/library/logger/v2"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type LoanProductMaintenanceRepository struct {
	Ottolog logger.OttologInterface
}

func InitLoanProductMaintenanceRepository(logs logger.OttologInterface) *LoanProductMaintenanceRepository {
	return &LoanProductMaintenanceRepository{
		Ottolog:logs,
	}
}

func (r *LoanProductMaintenanceRepository)Save(req *dbmodels.LoanProductMaintenance) error {
	db := GetDbCon()

	if err:= db.Save(&req).Error; err !=nil{
		r.Ottolog.Error("Error save to db " + err.Error())
		return err
	}

	return nil
}

func (r *LoanProductMaintenanceRepository)FindByID(id int64) (dbmodels.LoanProductMaintenance, error) {
	db := GetDbCon()

	var res dbmodels.LoanProductMaintenance


	if err:= db.Where("id = ?", id).First(&res).Error; err !=nil{
		r.Ottolog.Error("Error find data " + err.Error())
		return res, err
	}

	return res, nil
}


func (r *LoanProductMaintenanceRepository)Filter(req dto.ReqLoanProductMaintenanceDto) ([]dbmodels.LoanProductMaintenance, int, error) {
	db := GetDbCon()

	var res []dbmodels.LoanProductMaintenance
	limit := req.Limit
	page := req.Page
	var total int

	if req.BankCode != "" {
		db = db.Where("bank_code = ?",req.BankCode)
	}
	if req.Status != "" {
		db = db.Where("status = ?",req.Status)
	}
	if req.LoanProductCode != "" {
		db = db.Where("loan_product_code ilike ?","%" +req.LoanProductCode +"%")

	}

	if req.LoanProductName != "" {
		db = db.Where("loan_product_name = ?","%" +req.LoanProductName + "%")

	}

	err := db.Limit(limit).Offset((page-1)*limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		r.Ottolog.Error("error get data Validation Code master "+ err.Error())
		return res, 0, err
	}

	return res, total, nil

}