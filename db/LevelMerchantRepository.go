package db

import (
	"rose-be-go/models/dbmodels"
)

type LevelMerchantRepository struct {

}

func InitLevelMerchantRepository()  *LevelMerchantRepository{
	return &LevelMerchantRepository{}
}

func (repo *LevelMerchantRepository) GetAll(limit int64) ([]dbmodels.LevelMerchant, error) {
	db := GetDbCon()

	var data []dbmodels.LevelMerchant

	err := db.Limit(limit).Find(&data).Error

	return data, err
}