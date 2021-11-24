package db

import (
	"log"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"

	"ottodigital.id/library/logger/v2"
)

type SettlementConfigRepository struct {
	Ottolog logger.OttologInterface

}

func InitSettlementConfigRepository(logs logger.OttologInterface) *SettlementConfigRepository {
	return &SettlementConfigRepository{
		Ottolog:logs,
	}
}

func (r *SettlementConfigRepository)Save(req *dbmodels.SettlementConfig) error {
	db := GetDbCon()

	if err:= db.Save(&req).Error; err !=nil{
		r.Ottolog.Error("Error save to db " + err.Error())
		return err
	}

	return nil
}

func (r *SettlementConfigRepository)FindByBankCode(bankCode string) (dbmodels.SettlementConfig, error) {
	db := GetDbCon()

	var res dbmodels.SettlementConfig

	var status = "Active"
	if err:= db.Where("bank_code = ? and status = ? ", bankCode, status).First(&res).Error; err !=nil{
		r.Ottolog.Error("Error find data " + err.Error())
		return res, err
	}

	return res, nil
}


func (r *SettlementConfigRepository)Filter(req dto.ReqSettlementConfigDto) ([]dto.ResSettlementConfigDto, int, error) {
	db := GetDbCon()

	var res []dto.ResSettlementConfigDto
	filterBank := req.Code
	filterStatus := req.Status
	limit := req.Limit
	page := req.Page
	type Count struct {
		Count int `json:"count"`
	}
	var total Count
	var queryFilter string

	if filterBank != "" && filterStatus != ""  {
		queryFilter = "Where a.bank_code =  '" + filterBank + "' and a.status = '" + filterStatus + "'"
	} else if filterBank == "" && filterStatus != "" {
		queryFilter = "Where a.status =  '" + filterStatus + "'" 
	} else if filterBank != "" && filterStatus == "" {
		queryFilter = "Where a.bank_code = '" + filterBank + "'"
	} 
	err := db.Raw("select a.*, b.full_name bank_name " +
			"from host_settlement_fee_config a " +
			"LEFT JOIN bank_list b ON b.code = a.bank_code " +
			queryFilter).Limit(limit).Offset((page-1)*limit).Order("updated_at desc").
			Scan(&res).Limit(-1).Offset(0).Error

	if err != nil {
		r.Ottolog.Error("error get data Settlement list "+ err.Error())
		return res, 0, err
	}

	err = db.Raw("select count (a.id) " +
		"from host_settlement_fee_config a " +
		"LEFT JOIN bank_list b ON b.code = a.bank_code " +
		queryFilter).Scan(&total).Error

	
	log.Println("ini total data :", total)
	if err != nil {
		log.Println("err get total data :", err, total)
		return res, 0, err
	}

	return res, total.Count, nil

}