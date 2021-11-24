package db

import (
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type AkuisisiSfaRepository struct {
	
}

func InitAkuisisiSfaRepository()*AkuisisiSfaRepository  {
	return &AkuisisiSfaRepository{}
}

func (repo *AkuisisiSfaRepository) Filter(req dto.ReqAkuisisiSfa) ([]dbmodels.AkuisisiSfa, int, error ){
	db := GetDbCon()
	var res []dbmodels.AkuisisiSfa
	var total int
	page := req.Page
	limit := req.Limit

	db = db.Where("status = 'n'")
	keySearch := "%" + req.Key + "%"
	if req.Key != "" && req.StartDate != "" && req.EndDate != "" {
		db = db.Where("store_name ilike ? or store_phone_number ilike ? or merchant_outlet_id ilike ? and updated_at >= ? and updated_at <= ?", keySearch, keySearch, keySearch, req.StartDate, req.EndDate)
	} else if req.Key == "" && req.StartDate != "" && req.EndDate != "" {
		db = db.Where("updated_at >= ? and updated_at <= ?", req.StartDate, req.EndDate)
	} else if req.Key != "" && req.StartDate == "" && req.EndDate == "" {
		db = db.Where("store_name ilike ? or store_phone_number ilike ? or merchant_outlet_id ilike ? ", keySearch, keySearch, keySearch)
	}

	err := db.Order("id desc").Limit(limit).Offset((page-1)*limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error

	return res, total, err

}

