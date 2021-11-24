package db

import (
	"log"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"time"
)

type ValidationRepository struct {
	Ottolog logger.OttologInterface


}

func InitValidationRepository(logs logger.OttologInterface) *ValidationRepository {
	return &ValidationRepository{
		Ottolog:logs,

	}
}

func (r *ValidationRepository)Save(req *dbmodels.ValidationCode) error {
	db := GetDbCon()

	if err:= db.Save(&req).Error; err !=nil{
		r.Ottolog.Error("Error save to db " + err.Error())
		return err
	}

	return nil
}

func (r *ValidationRepository)Filter(req dto.ReqValidationCodeDto) ([]dto.ResValidationCodeDto, int, error) {
	db := GetDbCon()

	var res []dto.ResValidationCodeDto
	limit := req.Limit
	page := req.Page
	var total int

	if req.ValidationCode != "" {
		db = db.Where("validation_code ilike ?", "%" + req.ValidationCode +"%")
	}

	if req.AppID != "" {
		db = db.Where("app_id = ?", req.AppID)
	}

	if req.UserCategoryCode != "" {
		db = db.Where("user_category_code = ?", req.UserCategoryCode)
	}

	db = db.Table("validation_code_self_register a").Select("a.*, b.name app_name").
		Joins("LEFT JOIN lookup b on b.code = a.app_id").
		Where("b.lookup_group='APP_ID'")


	err := db.Limit(limit).Offset((page-1)*limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		r.Ottolog.Error("error get data Validation Code "+ err.Error())
		return res, 0, err
	}

	return res, total, nil

}

func (r *ValidationRepository)FindByID(id int64) (dbmodels.ValidationCode, error) {
	db := GetDbCon()

	var res dbmodels.ValidationCode


	if err:= db.Where("id = ?", id).First(&res).Error; err !=nil{
		r.Ottolog.Error("Error find data " + err.Error())
		return res, err
	}

	return res, nil
}



func (repo *ValidationRepository) CountValidationCode(code string) (int, error) {


	var count int
	db := DbCon
	err := db.Table("validation_code_self_register").
		Where("validation_code = ?", code).
		Where("valid_from <= ?", time.Now()).
		Where("valid_to >=?", time.Now()).

		Count(&count).Error

	log.Println(count)
	log.Println(err)
	return count, err
}
