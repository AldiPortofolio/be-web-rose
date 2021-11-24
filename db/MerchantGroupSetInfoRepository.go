package db

import "rose-be-go/models/dbmodels"

// MerchantGroupRepository
type MerchantGroupSetInfoRepository struct {
	// struct attributes
}

// InitMerchantGroupRepository ..
func InitMerchantGroupSetInfoRepository() *MerchantGroupSetInfoRepository {
	return &MerchantGroupSetInfoRepository{}
}

func (repo *MerchantGroupSetInfoRepository) Get(data dbmodels.MerchantGroupSetInfo) (dbmodels.MerchantGroupSetInfo, error) {
	db := GetDbCon()
	var res dbmodels.MerchantGroupSetInfo

	err := db.Where(&data).First(&res).Error

	return res, err
}

func (repo *MerchantGroupSetInfoRepository) GetById (id int64) (dbmodels.MerchantGroupSetInfo, error) {
	db := GetDbCon()
	var res dbmodels.MerchantGroupSetInfo

	err := db.Where("id = ?", id).First(&res).Error

	return res, err
}
