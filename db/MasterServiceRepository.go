package db

import (
	"rose-be-go/models/dbmodels"
)

// MasterServiceRepository ..
type MasterServiceRepository struct {
	// struct attributes
}

// InitMasterServiceRepository ..
func InitMasterServiceRepository() *MasterServiceRepository {
	return &MasterServiceRepository{}
}

// Save ..
func (repo *MasterServiceRepository) Save(req *dbmodels.MasterService) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

// FindByID ..
func (repo *MasterServiceRepository) FindByID(id int64) (dbmodels.MasterService, error) {
	db := GetDbCon()

	var data dbmodels.MasterService

	err := db.Where(dbmodels.MasterService{
		ID: id,
	}).First(&data).Error

	return data, err
}

// FindAll ..
func (repo *MasterServiceRepository) FindAll() ([]dbmodels.MasterService, int, error) {
	db := GetDbCon()
	var res []dbmodels.MasterService
	var total int

	err := db.Order("id asc").Find(&res).Count(&total).Error

	return res, total, err
}
