package db

import (
	"rose-be-go/constants/status_qrisconfig"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type QrisConfigRepository struct {

}

func InitQrisConfigRepository()  *QrisConfigRepository{
	return &QrisConfigRepository{}
}

func (repo *QrisConfigRepository) FindByID(id int64) (dbmodels.QrisConfig, error) {
	db := GetDbCon()

	var data dbmodels.QrisConfig

	err := db.Where(dbmodels.QrisConfig{
		ID: id,
	}).First(&data).Error

	return data, err
}

func (repo *QrisConfigRepository) Save(req *dbmodels.QrisConfig) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

func (repo *QrisConfigRepository) Filter(req dto.ReqQrisConfigDto) ([]dbmodels.QrisConfig, int, error) {
	db := GetDbCon()
	var res []dbmodels.QrisConfig
	var total int
	page := req.Page
	limit := req.Limit

	db = db.Where("status = ?", status_qrisconfig.ACTIVE)

	if req.IssuerName != "" {
		db = db.Where("issuer_name like ?", "%" + req.IssuerName + "%")
	}

	if req.InstitutionID != "" {
		db = db.Where("institution_id like ?", "%" + req.InstitutionID + "%")
	}

	if req.TransactionType != "" {
		db = db.Where("transaction_type = ?",  req.TransactionType)
	}

	err := db.Limit(limit).Offset((page-1)*limit).Order("issuer_name asc").Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		return res, 0, err
	}

	return res, total, nil
}