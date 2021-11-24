package db

import "rose-be-go/models/dbmodels"

// MerchantGroupRepository
type MerchantGroupInternalContactPersonRepository struct {
	// struct attributes
}

// InitMerchantGroupRepository ..
func InitMerchantGroupInternalContactPersonRepository() *MerchantGroupInternalContactPersonRepository {
	return &MerchantGroupInternalContactPersonRepository{}
}

func (repo *MerchantGroupInternalContactPersonRepository) Get(data dbmodels.MerchantGroupInternalContactPerson) (dbmodels.MerchantGroupInternalContactPerson, error) {
	db := GetDbCon()
	var res dbmodels.MerchantGroupInternalContactPerson

	err := db.Where(&data).First(&res).Error

	return res, err
}

func (repo *MerchantGroupInternalContactPersonRepository) GetById (id int64) (dbmodels.MerchantGroupInternalContactPerson, error) {
	db := GetDbCon()
	var res dbmodels.MerchantGroupInternalContactPerson

	err := db.Where("id = ? ", id).First(&res).Error

	return res, err
}