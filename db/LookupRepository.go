package db

import (
	"log"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type LookupRepository struct {

}

func InitLookupRepository() *LookupRepository {
	return &LookupRepository{}
}

func (repo *LookupRepository) Get(data dbmodels.Lookup) (dbmodels.Lookup, error) {
	db := GetDbCon()
	var res dbmodels.Lookup

	err := db.Where(&data).First(&res).Error

	return res, err
}


func (r *LookupRepository)FindByID(id int64) (dbmodels.Lookup, error) {
	db := GetDbCon()

	var res dbmodels.Lookup


	if err:= db.Where("id = ?", id).First(&res).Error; err !=nil{
		log.Println("Error find data " + err.Error())
		return res, err
	}

	return res, nil
}

func (r *LookupRepository) Save(req *dbmodels.Lookup) error {
	db := GetDbCon()

	if err:= db.Save(&req).Error; err !=nil{
		log.Println("Error save to db " + err.Error())
		return err
	}

	return nil
}


func (r *LookupRepository) Delete(lookupGroup dbmodels.Lookup) error {
	db := GetDbCon()

	if err:= db.Delete(&lookupGroup).Error; err !=nil{
		log.Println("Error delete froomm db " + err.Error())
		return err
	}

	return nil
}

func (r *LookupRepository)Filter(req dto.ReqLookupDto) ([]dbmodels.Lookup, int, error) {
	db := GetDbCon()

	var res []dbmodels.Lookup
	limit := req.Limit
	page := req.Page
	var total int

	if req.Name != "" {
		db = db.Where("name ilike ?", "%" + req.Name +"%")
	}

	if req.Code != "" {
		db = db.Where("code ilike ?", "%" + req.Code +"%")
	}

	if req.LookupGroup != "" {
		db = db.Where("lookup_group = ?",  req.LookupGroup)
	}

	err := db.Order("lookup_group asc, order_no asc").Limit(limit).Offset((page-1)*limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		log.Println("error get data Lookupgroup "+ err.Error())
		return res, 0, err
	}

	return res, total, nil

}