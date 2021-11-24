package db

import (
	"ottodigital.id/library/logger/v2"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type LookupGroupRepository struct {
	Ottolog logger.OttologInterface

}

func InitLookupGroupRepository(logs logger.OttologInterface) *LookupGroupRepository {
	return &LookupGroupRepository{
		Ottolog:logs,

	}

}

func (r *LookupGroupRepository)FindByID(id int64) (dbmodels.LookupGroup, error) {
	db := GetDbCon()

	var res dbmodels.LookupGroup


	if err:= db.Where("id = ?", id).First(&res).Error; err !=nil{
		r.Ottolog.Error("Error find data " + err.Error())
		return res, err
	}

	return res, nil
}

func (r *LookupGroupRepository) Save(req *dbmodels.LookupGroup) error {
	db := GetDbCon()

	if err:= db.Save(&req).Error; err !=nil{
		r.Ottolog.Error("Error save to db " + err.Error())
		return err
	}

	return nil
}

func (r *LookupGroupRepository) Delete(lookupGroup dbmodels.LookupGroup) error {
	db := GetDbCon()

	if err:= db.Delete(&lookupGroup).Error; err !=nil{
		r.Ottolog.Error("Error delete froomm db " + err.Error())
		return err
	}

	return nil
}

func (r *LookupGroupRepository) FindAll() ([]dbmodels.LookupGroup, error) {
	db:=GetDbCon()
	var data []dbmodels.LookupGroup
	err:= db.Find(&data).Order("name asc").Error
	if err!= nil {
		r.Ottolog.Error("Error Get data lookupgroup from db" + err.Error())
	}
	return data, err
}

func (r *LookupGroupRepository)Filter(req dto.ReqLookupGroupDto) ([]dbmodels.LookupGroup, int, error) {
	db := GetDbCon()

	var res []dbmodels.LookupGroup
	limit := req.Limit
	page := req.Page
	var total int

	if req.Name != "" {
		db = db.Where("name ilike ?", "%" + req.Name +"%")
	}

	err := db.Order("name asc").Limit(limit).Offset((page-1)*limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		r.Ottolog.Error("error get data Lookupgroup "+ err.Error())
		return res, 0, err
	}

	return res, total, nil

}