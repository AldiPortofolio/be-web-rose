package db

import (
	
	"rose-be-go/models/dbmodels"
	// "strconv"
)

// MerchantGroupRepository
type MerchantGroupRepository struct {
	// struct attributes
}

// InitMerchantGroupRepository ..
func InitMerchantGroupRepository() *MerchantGroupRepository {
	return &MerchantGroupRepository{}
}

// Get ..
func (repo *MerchantGroupRepository) Get(data dbmodels.MerchantGroup) (dbmodels.MerchantGroup, error) {
	db := GetDbCon()
	var res dbmodels.MerchantGroup

	err := db.Where(&data).First(&res).Error

	return res, err
}

//GetDetail...
func (repo *MerchantGroupRepository) GetDetail(data dbmodels.MerchantGroup) (dbmodels.MerchantGroup, error) {
	db := GetDbCon()
	var res dbmodels.MerchantGroup

	err := db.Where(&data).First(&res).Error

	return res, err
}

//GetDetailById...
func (repo *MerchantGroupRepository) GetDetailById(id int64) (dbmodels.MerchantGroupDetail, error) {
	db := GetDbCon()
	var res dbmodels.MerchantGroupDetail

	err := db.Where("id = ?", id).First(&res).Error

	return res, err
}

// Update ..
func (repo *MerchantGroupRepository) ActivationBp(data dbmodels.MerchantGroup) (dbmodels.MerchantGroup, error) {
	db := GetDbCon()
	var res dbmodels.MerchantGroup

	// err := db.Model(&res).Where("id = ?", data.ID).Update(dbmodels.MerchantGroup{EmailPortal:data.EmailPortal, PortalStatus:data.PortalStatus}).Error
	err := db.Model(&res).Where("id = ?", data.ID).Update(dbmodels.MerchantGroup{EmailPortal:data.EmailPortal}).Error
	return res, err
}

func (repo *MerchantGroupRepository) MerchantGroupActivationBp(id []int)  error {
	db := GetDbCon()
	// var idString string
	// for i := 0; i <= len(id); i++ {
	// 	if i == 0 {
	// 		str := strconv.Itoa(id[i])
	// 		idString = "(" + str
	// 	} else if i == len(id) {
	// 		idString = "("
	// 	}	else {
	// 		str := strconv.Itoa(id[i])
	// 		idString = "(" + str
	// 	}
	// }
	err := db.Table("merchant").Where("id IN (?)", id).Updates(map[string]interface{}{"portal_status": 1}).Error
	// err := db.Raw("UPDATE "+"merchant"+" SET "+"portal_status"+" = 1  WHERE id IN "+idString).Error

	return err
}
