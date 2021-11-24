package db

import (
	"rose-be-go/models/dbmodels"
)

type LimitTransactionDepositRepository struct {

}

func InitLimitTransactionDepositRepository() *LimitTransactionDepositRepository {
	return &LimitTransactionDepositRepository{}
}

func (repo *LimitTransactionDepositRepository) GetUserCategory() ([]dbmodels.UserCategory, error) {
	db := GetDbCon()

	var data []dbmodels.UserCategory

	err := db.Where(dbmodels.UserCategory{}).Find(&data).Error

	return data, err
}

func (repo *LimitTransactionDepositRepository) GetLevelMerchant() ([]dbmodels.LevelMerchant, error) {
	db := GetDbCon()

	var data []dbmodels.LevelMerchant

	err := db.Where(dbmodels.LevelMerchant{}).Find(&data).Error

	return data, err
}

