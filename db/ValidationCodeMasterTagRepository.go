package db

import (
	"ottodigital.id/library/logger/v2"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type ValidationCodeMasterTagRepository struct {
	Ottolog logger.OttologInterface
}

func InitValidationCodeMasterTagRepository(logs logger.OttologInterface) *ValidationCodeMasterTagRepository {
	return &ValidationCodeMasterTagRepository{
		Ottolog:logs,
	}
}

func (r *ValidationCodeMasterTagRepository)Save(req *dbmodels.ValidationCodeMasterTag) error {
	db := GetDbCon()

	if err:= db.Save(&req).Error; err !=nil{
		r.Ottolog.Error("Error save to db " + err.Error())
		return err
	}

	return nil
}

func (r *ValidationCodeMasterTagRepository)FindByID(id int64) (dbmodels.ValidationCodeMasterTag, error) {
	db := GetDbCon()

	var res dbmodels.ValidationCodeMasterTag


	if err:= db.Where("id = ?", id).First(&res).Error; err !=nil{
		r.Ottolog.Error("Error find data " + err.Error())
		return res, err
	}

	return res, nil
}

func (r *ValidationCodeMasterTagRepository)Filter(req dto.ReqValidationCodeMasterTagDto) ([]dto.ResValidationCodeMasterTagDto, int, error) {
	db := GetDbCon()

	var res []dto.ResValidationCodeMasterTagDto
	limit := req.Limit
	page := req.Page
	var total int

	if req.ValidationCodeID != 0 {
		db = db.Where("a.validation_code_id = ?",req.ValidationCodeID)
	}
	if req.MasterTagCode != "" {
		db = db.Where("b.code ilike ?", "%" + req.MasterTagCode +"%")

	}

	db = db.Table("validation_code_master_tag a").
		Select(" a.*, c.validation_code validation_code, a.master_tag_id master_tag_id, b.code master_tag_code ").
		Joins("left join master_tag b on a.master_tag_id = b.id").
		Joins("left join validation_code_self_register c on a.validation_code_id = c.id ")

	err := db.Limit(limit).Offset((page-1)*limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		r.Ottolog.Error("error get data Validation Code master "+ err.Error())
		return res, 0, err
	}

	return res, total, nil

}