package db

import (
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type ImageManagementRepository struct {

}

func InitImageManagementRepository()  *ImageManagementRepository{
	return &ImageManagementRepository{}
}

func (repo *ImageManagementRepository) FindByID(id int64) (dbmodels.ImageManagement, error) {
	db := GetDbCon()

	var data dbmodels.ImageManagement

	err := db.Where(dbmodels.ImageManagement{
		ID: id,
	}).First(&data).Error

	return data, err
}

func (repo *ImageManagementRepository) Save(req *dbmodels.ImageManagement) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

func (repo *ImageManagementRepository) Filter(req dto.ReqImageManagementDto) ([]dbmodels.ImageManagement, int, error) {
	db := GetDbCon()
	var res []dbmodels.ImageManagement
	var total int
	page := req.Page
	limit := req.Limit

	if req.ID  != 0 {
		db = db.Where("id = ?",  req.ID )
	}

	if req.Name != "" {
		db = db.Where("name like ?", "%" + req.Name + "%")
	}

	if req.URL != "" {
		db = db.Where("url like ?", "%" + req.URL + "%")
	}

	if req.Notes != "" {
		db = db.Where("notes like ?", "%" + req.Notes + "%")
	}

	err := db.Limit(limit).Offset((page-1)*limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		return res, 0, err
	}

	return res, total, nil
}