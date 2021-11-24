package db

import (
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type LimitTransactionRepository struct {

}

func InitLimitTransactionRepository() *LimitTransactionRepository {
	return &LimitTransactionRepository{}
}

func (repo *LimitTransactionRepository) FindByID(id int64) (dbmodels.LimitTransaction, error) {
	db := GetDbCon()

	var data dbmodels.LimitTransaction

	err := db.Where(dbmodels.LimitTransaction{
		ID: id,
	}).First(&data).Error

	return data, err
}

func (repo *LimitTransactionRepository) Save(req *dbmodels.LimitTransaction) error  {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

func (repo *LimitTransactionRepository) Filter(req dto.ReqLimitTransactionDto) ([]dbmodels.LimitTransaction, int, error) {
	db := GetDbCon()
	var res []dbmodels.LimitTransaction
	var total int
	page := req.Page
	limit := req.Limit

	db = db.Table("limit_transaction LT ").
		Select("LT.*")

	if req.UserCategory  != "" {
		db = db.Where("LT.user_category = ?",  req.UserCategory )
	}

	if req.LevelMerchant  != "" {
		db = db.Where("LT.level_merchant = ?",  req.LevelMerchant )
	}

	if req.FeatureProduct  != "" {
		db = db.Where("LT.feature_product = ?",  req.FeatureProduct )
	}

	if req.TimeFrame != "" {
		db = db.Where("LT.time_frame = ?", req.TimeFrame)
	}

	err := db.Limit(limit).Offset((page-1)*limit).Order("user_category, level_merchant, feature_product, time_frame asc").Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		return res, 0, err
	}

	return res, total, nil
}