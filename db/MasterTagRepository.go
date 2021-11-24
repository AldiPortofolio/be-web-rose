package db

import (
	"ottodigital.id/library/logger/v2"
	"rose-be-go/constants"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type MasterTagRepository struct {
	Ottolog logger.OttologInterface

}

func InitMasterTagRepository(logs logger.OttologInterface) *MasterTagRepository {
	return &MasterTagRepository{
		Ottolog:logs,
	}
}

func (r *MasterTagRepository)Save(req *dbmodels.MasterTag) error {
	db := GetDbCon()

	if err:= db.Save(&req).Error; err !=nil{
		r.Ottolog.Error("Error save to db " + err.Error())
		return err
	}

	return nil
}

func (r *MasterTagRepository)FindByID(id int64) (dbmodels.MasterTag, error) {
	db := GetDbCon()

	var res dbmodels.MasterTag


	if err:= db.Where("id = ?", id).First(&res).Error; err !=nil{
		r.Ottolog.Error("Error find data " + err.Error())
		return res, err
	}

	return res, nil
}


func (r *MasterTagRepository)FindByCode(code string) (dbmodels.MasterTag, error) {
	db := GetDbCon()

	var res dbmodels.MasterTag


	if err:= db.Where("code = ?", code).First(&res).Error; err !=nil{
		r.Ottolog.Error("Error find data " + err.Error())
		return res, err
	}

	return res, nil
}

func (r *MasterTagRepository)Filter(req dto.ReqMasterTagDto) ([]dbmodels.MasterTag, int, error) {
	db := GetDbCon()

	var res []dbmodels.MasterTag
	limit := req.Limit
	page := req.Page
	var total int

	if req.Code != "" {
		db = db.Where("code ilike ?", "%" + req.Code +"%")
	}

	if req.Name != "" {
		db = db.Where("name ilike ?", "%" + req.Name +"%")
	}

	if req.Status != "" {
		switch req.Status {
		case constants.IN_ACTIVE:
			db = db.Where("status = false")
			
		case constants.ACTIVE:
			db = db.Where("status = true")
			
		}
	}

	err := db.Limit(limit).Offset((page-1)*limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		r.Ottolog.Error("error get data master tag "+ err.Error())
		return res, 0, err
	}

	return res, total, nil

}


func (r *MasterTagRepository)GetAll() ([]dbmodels.MasterTag, error) {
	db := GetDbCon()

	var res []dbmodels.MasterTag

	err := db.Find(&res).Error

	if err != nil {
		r.Ottolog.Error("error get data master tag "+ err.Error())
		return res, err
	}

	return res, nil

}