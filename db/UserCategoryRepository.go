package db

import (
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type UserCategoryRepository struct {
}

func InitUserCategoryRepository() *UserCategoryRepository {
	return &UserCategoryRepository{}
}

func (repo *UserCategoryRepository) FindByID(id int64) (dbmodels.UserCategory, error) {
	db := GetDbCon()

	var data dbmodels.UserCategory

	err := db.Where(dbmodels.UserCategory{
		ID: id,
	}).First(&data).Error

	return data, err
}

func (repo *UserCategoryRepository) Save(req *dbmodels.UserCategory) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

func (repo *UserCategoryRepository) Filter(req dto.ReqUserCategoryDto) ([]dto.ResUserCategoryDto, int, error) {
	db := GetDbCon()
	var res []dto.ResUserCategoryDto
	var total int
	page := req.Page
	limit := req.Limit

	filter := "seq asc"
	if limit == 999 {
		filter = "code asc"
	}

	if req.ID != 0 {
		db = db.Where("a.id = ?", req.ID)
	}

	if req.Code != "" {
		db = db.Where("a.code like ?", "%"+req.Code+"%")
	}

	if req.Name != "" {
		db = db.Where("a.name like ?", "%"+req.Name+"%")
	}

	if req.Notes != "" {
		db = db.Where("a.notes like ?", "%"+req.Notes+"%")
	}

	if req.Seq != 0 {
		db = db.Where("a.seq = ?", req.Seq)
	}

	db = db.Table("user_category a").Select("a.*, b.name app_name").
		Joins("LEFT JOIN lookup b on b.code = a.app_id").
		Where("b.lookup_group='APP_ID'")

	err := db.Limit(limit).Offset((page - 1) * limit).Order(filter).Scan(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		return res, 0, err
	}

	return res, total, nil
}

func (repo *UserCategoryRepository) DropdownList() ([]dto.ResUserCategoryDropdownLisDto, int, error) {
	db := GetDbCon()
	var res []dto.ResUserCategoryDropdownLisDto
	var total int

	db = db.Table("user_category a").Select("a.id, a.code, a.name").
		Joins("LEFT JOIN lookup b on b.code = a.app_id").
		Where("b.lookup_group='APP_ID'").Order("lower(a.name) ASC")

	err := db.Scan(&res).Count(&total).Error

	if err != nil {
		return res, 0, err
	}

	return res, total, nil
}
