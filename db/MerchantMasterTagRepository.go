package db

import (
	"fmt"
	"log"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type MerchantMasterTagRepository struct {

}

func InitMerchantMasterTagRepository() *MerchantMasterTagRepository {
	return &MerchantMasterTagRepository{}
}

func (r *MerchantMasterTagRepository)GetAll() ([]dbmodels.MerchantMasterTag, error) {
	db := GetDbCon()

	var res []dbmodels.MerchantMasterTag

	err := db.Find(&res).Error

	return res, err
}

func (r *MerchantMasterTagRepository)Filter(req dto.ReqMerchantMasterTagDto) ([]dbmodels.MerchantMasterTag, int, error) {
	db := GetDbCon()

	var res []dbmodels.MerchantMasterTag
	limit := req.Limit
	page := req.Page
	var total int

	if req.Mid != "" {
		db = db.Where("mid ilike ?", "%" + req.Mid +"%")
	}

	if req.MasterTagCode != "" {
		db = db.Where("master_tag_code ilike ?", "%" + req.MasterTagCode +"%")
	}

	err := db.Order("updated_at desc").Limit(limit).Offset((page-1)*limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		fmt.Println("error get data master tag "+ err.Error())
		return res, 0, err
	}

	return res, total, nil

}

func (r *MerchantMasterTagRepository)FindByMid(req dto.ReqMerchantMasterTagByMidDto) ([]dto.ResMerchantMasterTag, int, error) {
	db := GetDbCon()

	var data []dto.ResMerchantMasterTag
	
	var total int

	err:=db.Table("merchant_master_tag a").Select("a.mid mid, a.master_tag_code master_tag_code, b.name").
		Joins("LEFT JOIN master_tag b on b.code = a.master_tag_code").
		Where("a.mid = ? and b.status = ?", req.Mid, true).
		// Where("b.status = true").
		Scan(&data).Count(&total).Error
	if err!= nil {
		log.Println("err get merchant master tag by mid", err)
	}
	return data, total, err

	

}