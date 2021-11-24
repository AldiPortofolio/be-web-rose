package db

import (
	"rose-be-go/constants/status_mdr_tenor"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type MdrTenorRepository struct {

}

func InitMdrTenorRepository() *MdrTenorRepository {
	return &MdrTenorRepository{}
}

func (r *MdrTenorRepository) Save(req *dbmodels.MdrTenor) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

// Get ..
func (r *MdrTenorRepository) Get(data dbmodels.MdrTenor) (dbmodels.MdrTenor, error) {
	db := GetDbCon()
	var res dbmodels.MdrTenor

	err := db.Where(&data).First(&res).Error

	return res, err
}


// FindByID ..
func (r *MdrTenorRepository) FindByID(id int64) (dbmodels.MdrTenor, error) {
	db := GetDbCon()

	var data dbmodels.MdrTenor

	err := db.Where(dbmodels.MdrTenor{
		ID: id,
	}).First(&data).Error

	return data, err
}


// Filter ..
func (r *MdrTenorRepository) Filter(req dto.ReqMdrTenorDto) ([]dbmodels.MdrTenor, int, error) {
	db := GetDbCon()
	var res []dbmodels.MdrTenor
	var total int
	page := req.Page
	limit := req.Limit

	db = db.Where("status = ?", status_mdr_tenor.ACTIVE)

	if req.TenorName != "" {
		db = db.Where("tenor_name ilike ?", "%" + req.TenorName + "%")
	}

	if req.TenorCode != "" {
		db = db.Where("tenor_code ilike ?", "%" + req.TenorCode + "%")
	}

	err := db.Limit(limit).Offset((page-1)*limit).Order("seq asc").Find(&res).Limit(-1).Offset(0).Count(&total).Error



	return res, total, err
}

func (repo *MdrTenorRepository) FindAll() ([]dbmodels.MdrTenor, error) {
	db := GetDbCon()
	var res []dbmodels.MdrTenor

	db = db.Where("status = ?", status_mdr_tenor.ACTIVE)


	err := db.Order("seq asc").Find(&res).Error



	return res, err

}