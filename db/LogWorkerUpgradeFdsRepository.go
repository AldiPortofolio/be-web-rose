package db

import (
	"ottodigital.id/library/logger/v2"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type LogUpgradeFdsRepository struct {
	Ottolog logger.OttologInterface

}

func InitLogUpgradeFdsRepository(logs logger.OttologInterface) *LogUpgradeFdsRepository {
	return &LogUpgradeFdsRepository{
		Ottolog:logs,

	}
}


func (r *LogUpgradeFdsRepository)Save(req *dbmodels.LogUpgradeFds) error {
	db := GetDbCon()

	if err:= db.Save(&req).Error; err !=nil{
		r.Ottolog.Error("Error save to db " + err.Error())
		return err
	}

	return nil
}


func (r *LogUpgradeFdsRepository) Filter(req dto.ReqLogUpgradeFdsDto) ([]dbmodels.LogUpgradeFds, int, error) {
	db := GetDbCon()

	var res []dbmodels.LogUpgradeFds
	limit := req.Limit
	page := req.Page
	var total int

	if req.PhoneNumber != "" {
		db = db.Where("phone_number ilike ?", "%" + req.PhoneNumber +"%")
	}

	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}


	err := db.Limit(limit).Offset((page-1)*limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		r.Ottolog.Error("error get data log upgrade fds "+ err.Error())
		return res, 0, err
	}

	return res, total, nil

}


func (r *LogUpgradeFdsRepository)FindByID(id int64) (dbmodels.LogUpgradeFds, error) {
	db := GetDbCon()

	var res dbmodels.LogUpgradeFds


	if err:= db.Where("id = ?", id).First(&res).Error; err !=nil{
		r.Ottolog.Error("Error find data " + err.Error())
		return res, err
	}

	return res, nil
}