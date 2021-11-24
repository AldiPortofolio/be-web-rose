package db

import (
	
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"

	"ottodigital.id/library/logger/v2"
)

type PartnerLinkRepository struct {
	Ottolog logger.OttologInterface

}

func InitPartnerLinkRepository(logs logger.OttologInterface) *PartnerLinkRepository {
	return &PartnerLinkRepository{
		Ottolog:logs,
	}
}

func (r *PartnerLinkRepository)Save(req *dbmodels.PartnerLink) error {
	db := GetDbCon()

	if err:= db.Save(&req).Error; err !=nil{
		r.Ottolog.Error("Error save to db " + err.Error())
		return err
	}

	return nil
}

func (r *PartnerLinkRepository)FindByIdAndCode(merchantId int, code string) (dbmodels.PartnerLink, error) {
	db := GetDbCon()

	var res dbmodels.PartnerLink
	err := db.Where("merchant_id = ? and code = ?", merchantId, code).First(&res).Error
	return res, err

}


func (r *PartnerLinkRepository)Filter(req dto.ReqPartnerLinkDto) ([]dbmodels.PartnerLink, int, error) {
	db := GetDbCon()

	var res []dbmodels.PartnerLink
	
	limit := req.Limit
	page := req.Page

	var total int

	if req.MerchantId != "" {
		db = db.Where("merchant_id = ?",  req.MerchantId )
	}

	err := db.Limit(limit).Offset((page-1)*limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		r.Ottolog.Error("error get data partner list "+ err.Error())
		return res, 0, err
	}

	
	return res, total, nil

}

func (r *PartnerLinkRepository)Delete(id int) error {
	db := GetDbCon()
	var partnerLink dbmodels.PartnerLink
	err:= db.Where("id = ?", id).Delete(&partnerLink).Error
	if err != nil {
		r.Ottolog.Error("Error save to db " + err.Error() )
		return err
	}

	return nil
}