package db

import "rose-be-go/models/dbmodels"

// MerchantGroupFeeInfoRepository
type MerchantGroupFeeInfoRepository struct {
	// struct attributes
}

// InitMerchantGroupFeeInfoRepository ..
func InitMerchantGroupFeeInfoRepository() *MerchantGroupFeeInfoRepository {
	return &MerchantGroupFeeInfoRepository{}
}

func (repo *MerchantGroupFeeInfoRepository) Get(data dbmodels.MerchantGroupFeeInfo) (dbmodels.MerchantGroupFeeInfo, error) {
	db := GetDbCon()
	var res dbmodels.MerchantGroupFeeInfo

	err := db.Where(&data).First(&res).Error

	return res, err
}

func (repo *MerchantGroupFeeInfoRepository) GetById (id int64) (dbmodels.MerchantGroupFeeInfo, error) {
	db := GetDbCon()
	var res dbmodels.MerchantGroupFeeInfo

	err := db.Where("id = ? ", id).First(&res).Error

	return res, err
}
