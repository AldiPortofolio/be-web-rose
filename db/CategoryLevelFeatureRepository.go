package db

import (
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type CategoryLevelFeatureRepository struct {

}

func InitCategoryLevelFeatureRepository()  *CategoryLevelFeatureRepository{
	return &CategoryLevelFeatureRepository{}
}

func (repo *CategoryLevelFeatureRepository) FindByID(id int64) (dbmodels.CategoryLevelFitur, error) {
	db := GetDbCon()

	var data dbmodels.CategoryLevelFitur

	err := db.Where(dbmodels.CategoryLevelFitur{
		ID: id,
	}).First(&data).Error

	return data, err
}

func (repo *CategoryLevelFeatureRepository) Save(req *dbmodels.CategoryLevelFitur) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

func (repo *CategoryLevelFeatureRepository) Filter(req dto.ReqCategoryLevelFeatureDto) ([]dto.ResCategoryLevelFeatureDto, int, error) {
	db := GetDbCon()
	var res []dto.ResCategoryLevelFeatureDto
	var total int
	page := req.Page
	limit := req.Limit

	db = db.Table("category_level_fitur CLF ").
		Select("CLF.*," +
			" UC.name as user_category_name," +
			" LM.name as level_merchant_name," +
			" FP.name as fitur_product_name").
		Joins("LEFT JOIN user_category UC on UC.id = CLF.user_category_id").
		Joins("LEFT JOIN level_merchant LM on LM.id = CLF.level_merchant_id").
		Joins("LEFT JOIN fitur_product FP on FP.id = CLF.fitur_product_id").
		Order("clf.id desc")

	if req.ID  != 0 {
		db = db.Where("CLF.id = ?",  req.ID )
	}

	if req.UserCategoryId  != 0 {
		db = db.Where("CLF.user_category_id = ?",  req.UserCategoryId )
	}

	if req.LevelMerchantId  != 0 {
		db = db.Where("CLF.level_merchant_id = ?",  req.LevelMerchantId )
	}

	if req.FiturProductId  != 0 {
		db = db.Where("CLF.fitur_product_id = ?",  req.FiturProductId )
	}

	if req.Status != "" {
		db = db.Where("CLF.status like ?", "%" + req.Status + "%")
	}

	err := db.Limit(limit).Offset((page-1)*limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		return res, 0, err
	}

	return res, total, nil
}