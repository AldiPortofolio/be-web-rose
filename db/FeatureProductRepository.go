package db

import (
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type FeatureProductRepository struct {

}

func InitFeatureProductRepository()  *FeatureProductRepository{
	return &FeatureProductRepository{}
}

func (repo *FeatureProductRepository) FindByID(id int64) (dbmodels.FiturProduct, error) {
	db := GetDbCon()

	var data dbmodels.FiturProduct

	err := db.Where(dbmodels.FiturProduct{
		ID: id,
	}).First(&data).Error

	return data, err
}

func (repo *FeatureProductRepository) Save(req *dbmodels.FiturProduct) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

func (repo *FeatureProductRepository) Filter(req dto.ReqFiturProductDto) ([]dto.ResFiturProductDto, int, error) {
	db := GetDbCon()
	var res []dto.ResFiturProductDto
	var total int
	page := req.Page
	limit := req.Limit

	db = db.Table("fitur_product FP ").
		Select("FP.*," +
			"P.name as product_name").
		Joins("LEFT JOIN product P on P.id = FP.product_id")

	if req.ID  != 0 {
		db = db.Where("FP.id = ?",  req.ID )
	}

	if req.ProductID  != 0 {
		db = db.Where("FP.product_id = ?",  req.ProductID )
	}

	if req.Code != "" {
		db = db.Where("FP.code like ?", "%" + req.Code + "%")
	}

	if req.Icon != "" {
		db = db.Where("FP.icon like ?", "%" + req.Icon + "%")
	}

	if req.Name != "" {
		db = db.Where("FP.name like ?", "%" + req.Name + "%")
	}

	if req.Notes != "" {
		db = db.Where("FP.notes like ?", "%" + req.Notes + "%")
	}

	if req.Seq  != 0 {
		db = db.Where("FP.seq = ?",  req.Seq )
	}

	err := db.Limit(limit).Offset((page-1)*limit).Order("name ASC").Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		return res, 0, err
	}

	return res, total, nil
}