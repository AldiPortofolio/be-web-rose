package db

import "rose-be-go/models/dbmodels"

// OwnerRepository
type OwnerRepository struct {
	// struct attributes
}

// InitMerchantGroupRepository ..
func InitOwnerRepository() *OwnerRepository {
	return &OwnerRepository{}
}

// Get ..
func (repo *OwnerRepository) Get(data dbmodels.Owner) (dbmodels.Owner, error) {
	db := GetDbCon()
	var res dbmodels.Owner

	err := db.Where(&data).First(&res).Error
	return res, err
}

// Update ..
func (repo *OwnerRepository) Update(data *dbmodels.Owner) (dbmodels.Owner, error) {
	db := GetDbCon()
	var res dbmodels.Owner

	err := db.Model(&res).Updates(&data).Error

	return res, err
}

func (r *OwnerRepository) FindByID(id int64) (dbmodels.Owner, error) {

	db := GetDbCon()
	var res dbmodels.Owner
	err := db.Where("id = ?", id).First(&res).Error
	return res, err


}

func (repo *OwnerRepository) Save(data *dbmodels.Owner) (error) {
	db := GetDbCon()
	

	err := db.Save(&data).Error

	return err
}

func (r *OwnerRepository) FindByID2(id int64) (dbmodels.OwnerDetail, error) {

	db := GetDbCon()
	var res dbmodels.OwnerDetail
	err := db.Where("id = ?", id).First(&res).Error
	return res, err


}