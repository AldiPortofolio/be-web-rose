package db

import (
	"rose-be-go/models/dbmodels"
)

type ConfigVaMerchantGroupRepository struct {
}

func InitConfigVaMerchantGroupRepository() *ConfigVaMerchantGroupRepository {
	return &ConfigVaMerchantGroupRepository{}
}

func (repo *ConfigVaMerchantGroupRepository) FindByGroupId(id int64) (dbmodels.ConfigVaMerchantGroup, error) {
	db := GetDbCon()

	var data dbmodels.ConfigVaMerchantGroup

	err := db.Where(dbmodels.ConfigVaMerchantGroup{
		MerchantGroupId: id,
	}).First(&data).Error

	return data, err
}

func (repo *ConfigVaMerchantGroupRepository) Save(req *dbmodels.ConfigVaMerchantGroup) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}


