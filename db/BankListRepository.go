package db

import (
	"ottodigital.id/library/logger/v2"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type BankListRepository struct {
	Ottolog logger.OttologInterface

}

func InitBankListRepository(logs logger.OttologInterface) *BankListRepository {
	return &BankListRepository{
		Ottolog:logs,
	}
}

func (r *BankListRepository)Save(req *dbmodels.BankList) error {
	db := GetDbCon()

	if err:= db.Save(&req).Error; err !=nil{
		r.Ottolog.Error("Error save to db " + err.Error())
		return err
	}

	return nil
}

func (r *BankListRepository)FindByID(id int64) (dbmodels.BankList, error) {
	db := GetDbCon()

	var res dbmodels.BankList


	if err:= db.Where("id = ?", id).First(&res).Error; err !=nil{
		r.Ottolog.Error("Error find data " + err.Error())
		return res, err
	}

	return res, nil
}

func (r *BankListRepository)FindByCode(code string) (dbmodels.BankList, error) {
	db := GetDbCon()

	var res dbmodels.BankList


	if err:= db.Where("code = ?", code).First(&res).Error; err !=nil{
		r.Ottolog.Error("Error find data " + err.Error())
		return res, err
	}

	return res, nil
}

func (r *BankListRepository)Filter(req dto.ReqBankListDto) ([]dto.ResBankList, int, error) {
	db := GetDbCon()

	var res []dto.ResBankList
	limit := req.Limit
	page := req.Page

	// if req.ShortName != "" {
	// 	db = db.Where("short_name ilike ?", "%" + req.ShortName +"%")
	// }

	// if req.FullName != "" {
	// 	db = db.Where("full_name ilike ?", "%" + req.FullName +"%")
	// }

	// if req.Code != "" {
	// 	db = db.Where("code ilike ?", "%" + req.Code +"%")
	// }
	// "Where a.portal_status =  '" + filterStatus + "' and e.name = '" + filterType +"'"
	var queryFilter string
	if req.ShortName != "" && req.FullName != "" && req.Code != "" {
		queryFilter = "where short_name ilike '%" + req.ShortName + "%' AND full_name ilike '%" + req.FullName + "%' AND code ilike '%" + req.Code +"%'"
	} else if req.ShortName != "" && req.FullName != "" && req.Code == "" {
		queryFilter = "where short_name ilike '%" + req.ShortName + "%' AND full_name ilike '%" + req.FullName + "%'"
	} else if req.ShortName != "" && req.FullName == "" && req.Code != "" {
		queryFilter = "where short_name ilike '%" + req.ShortName + "%' AND code ilike '%" + req.Code +"%'"
	} else if req.ShortName == "" && req.FullName != "" && req.Code != "" {
		queryFilter = "where full_name ilike '%" + req.FullName + "%' AND code ilike '%" + req.Code +"%'"
	} else if req.ShortName == "" && req.FullName == "" && req.Code != "" {
		queryFilter = "where code ilike '%" + req.Code +"%'"
	} else if req.ShortName != "" && req.FullName == "" && req.Code == "" {
		queryFilter = "where short_name ilike '%" + req.ShortName +"%'"
	} else if req.ShortName == "" && req.FullName != "" && req.Code == "" {
		queryFilter = "where full_name ilike '%" + req.FullName +"%'"
	} 


	err := db.Raw( "select a.*, b.settlement_type " +
			"from bank_list a " +
			"LEFT JOIN host_settlement_fee_config b ON b.bank_code = a.code and b.status = 'Active' " +
			queryFilter).Limit(limit).Offset((page-1)*limit).
			Scan(&res).Limit(-1).Offset(0).Error

	// err := db.Limit(limit).Offset((page-1)*limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error

	type Count struct {
		Count int `json:"count"`
	}
	var total Count

	err = db.Raw( "select count (a.id) " +
			"from bank_list a " +
			"LEFT JOIN host_settlement_fee_config b ON b.bank_code = a.code and b.status = 'Active' " +
			queryFilter).Scan(&total).Error
	

	if err != nil {
		r.Ottolog.Error("error get data bank list "+ err.Error())
		return res, 0, err
	}

	return res, total.Count , nil

}