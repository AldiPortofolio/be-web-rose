package db

import (
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type ProductRepository struct {

}

func InitProductRepository()  *ProductRepository{
	return &ProductRepository{}
}

func (repo *ProductRepository) FindByID(id int64) (dbmodels.Product, error) {
	db := GetDbCon()

	var data dbmodels.Product

	err := db.Where(dbmodels.Product{
		ID: id,
	}).First(&data).Error

	return data, err
}

func (repo *ProductRepository) Save(req *dbmodels.Product) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

func (repo *ProductRepository) Filter(req dto.ReqProductDto) ([]dbmodels.Product, int, error) {
	db := GetDbCon()
	var res []dbmodels.Product
	var total int
	page := req.Page
	limit := req.Limit

	if req.ID  != 0 {
		db = db.Where("id = ?",  req.ID )
	}

	if req.Code != "" {
		db = db.Where("code like ?", "%" + req.Code + "%")
	}

	if req.Name != "" {
		db = db.Where("name like ?", "%" + req.Name + "%")
	}

	if req.Title != "" {
		db = db.Where("title like ?", "%" + req.Title + "%")
	}

	if req.Desc != "" {
		db = db.Where("product.desc like ?", "%" + req.Desc + "%")
	}

	if req.Notes != "" {
		db = db.Where("notes like ?", "%" + req.Notes + "%")
	}

	if req.Seq  != 0 {
		db = db.Where("seq = ?",  req.Seq )
	}

	err := db.Limit(limit).Offset((page-1)*limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		return res, 0, err
	}

	return res, total, nil
}