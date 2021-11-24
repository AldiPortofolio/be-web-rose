package db

import (
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type AcquititionsRepository struct {
}

func InitAcquititionsRepository() *AcquititionsRepository {
	return &AcquititionsRepository{}
}

func (repo *AcquititionsRepository) FindByID(id int64) (dbmodels.Acquititions, error) {
	db := GetDbCon()

	var data dbmodels.Acquititions

	err := db.Where(dbmodels.Acquititions{
		ID: id,
	}).First(&data).Error

	return data, err
}

func (repo *AcquititionsRepository) Save(req *dbmodels.Acquititions) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

func (r *AcquititionsRepository) Delete(id int) error {
	db := GetDbCon()
	var Acquititions dbmodels.Acquititions
	err := db.Where("id = ?", id).Delete(&Acquititions).Error

	if err != nil {
		return err
	}

	return nil
}
func (repo *AcquititionsRepository) Filter(req dto.ReqFilterAcquititionsDto) ([]dbmodels.Acquititions, int, error) {
	db := GetDbCon()
	var res []dbmodels.Acquititions
	var name string
	var merchantCategory string
	var businessType string
	var total int
	var err error
	page := req.Page
	limit := req.Limit

	if req.Name == "" {
		name = "%"
	} else {
		name = "%" + req.Name + "%"
	} 

	if req.MerchantCategory == "" {
		merchantCategory = "%"
	} else {
		merchantCategory = "%" + req.MerchantCategory + "%"
	}
	
	if req.BusinessType == "" {
		businessType = "%"
	} else {
		businessType = "%" + req.BusinessType + "%"
	}

	if req.MerchantGroup != 0 {
		err = db.Where("name ilike ? and merchant_category ilike ? and merchant_group_id = ? and business_type ilike ?", name, merchantCategory, req.MerchantGroup, businessType).Limit(limit).Offset((page - 1) * limit).Order("id DESC").Find(&res).Limit(-1).Offset(0).Count(&total).Error
	} else {
		err = db.Where("name ilike ? and merchant_category ilike ? and business_type ilike ?", name, merchantCategory, businessType).Limit(limit).Offset((page - 1) * limit).Order("id DESC").Find(&res).Limit(-1).Offset(0).Count(&total).Error
	}

	if err != nil {
		return res, 0, err
	}

	return res, total, nil
}
