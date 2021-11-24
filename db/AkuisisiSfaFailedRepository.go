package db

import (
	"fmt"
	"log"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type AkuisisiSfaFailedRepository struct {
	
}

func InitAkuisisiSfaFailedRepository()*AkuisisiSfaFailedRepository  {
	return &AkuisisiSfaFailedRepository{}
}

func (repo *AkuisisiSfaFailedRepository) Filter(req dto.ReqAkuisisiSfaFailed) ([]dbmodels.AkuisisiSfaFailed, int, error ){
	db := GetDbCon()
	var res []dbmodels.AkuisisiSfaFailed
	page := req.Page
	limit := req.Limit

	
	var keySearch string

	if req.Key != "" && req.StartDate != "" && req.EndDate != "" {
		keySearch = "where store_name ilike '%" + req.Key + "%' or store_phone_number ilike '%" + req.Key + "%' or merchant_outlet_id ilike '%" + req.Key + "%' and updated_at >= '" + req.StartDate + "' and updated_at <= '" + req.EndDate + "' "
	} else if req.Key == "" && req.StartDate != "" && req.EndDate != "" {
		keySearch = "where updated_at >= '" + req.StartDate + "' and updated_at <= '" + req.EndDate + "' "
	} else if req.Key != "" && req.StartDate == "" && req.EndDate == "" {
		keySearch = "where store_name ilike '%" + req.Key + "%' or store_phone_number ilike '%" + req.Key + "%' or merchant_outlet_id ilike '%" + req.Key + "%' "
	}
	fmt.Println("----------->",keySearch)
	err := db.Raw("select * from akuisisi_sfa_failed " + keySearch).Limit(limit).Offset((page-1)*limit).Order("id desc").
			Scan(&res).Limit(-1).Offset(0).Error
	if err != nil {
		log.Println("Error Get Data Akuisisi SFA Failed : ", err)
		return res, 0 , err
	}

	type Count struct {
		Count int `json:"count"`
	}
	var total Count

	err = db.Raw("select count(id) from akuisisi_sfa_failed " + keySearch).Scan(&total).Error

	return res, total.Count , err

}


