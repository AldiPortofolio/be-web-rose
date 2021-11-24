package db

import (
	"rose-be-go/models/dbmodels"
)

// MasterTypeRepository ..
type MasterTypeRepository struct {
	// struct attributes
}

// InitMasterTypeRepository ..
func InitMasterTypeRepository() *MasterTypeRepository {
	return &MasterTypeRepository{}
}

// Save ..
func (repo *MasterTypeRepository) Save(req *dbmodels.MasterType) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

// FindByID ..
func (repo *MasterTypeRepository) FindByID(id int64) (dbmodels.MasterType, error) {
	db := GetDbCon()

	var data dbmodels.MasterType

	err := db.Where(dbmodels.MasterType{
		ID: id,
	}).First(&data).Error

	return data, err
}

// FindAll ..
func (repo *MasterTypeRepository) FindAll() ([]dbmodels.MasterType, int, error) {
	db := GetDbCon()
	var res []dbmodels.MasterType
	var total int

	err := db.Order("id asc").Find(&res).Count(&total).Error

	return res, total, err
}
