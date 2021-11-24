package db

import (
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

// FeeMdrSettingMerchantGroupRepository
type FeeMdrSettingMerchantGroupRepository struct {
	// struct attributes
}

// InitFeeMdrSettingMerchantGroupRepository ..
func InitFeeMdrSettingMerchantGroupRepository() *FeeMdrSettingMerchantGroupRepository {
	return &FeeMdrSettingMerchantGroupRepository{}
}

// Save ..
func (repo *FeeMdrSettingMerchantGroupRepository) Save(req *dbmodels.FeeMdrSettingMerchantGroup) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

// FindByID ..
func (repo *FeeMdrSettingMerchantGroupRepository) FindByID(id int64) (dbmodels.FeeMdrSettingMerchantGroup, error) {
	db := GetDbCon()

	var data dbmodels.FeeMdrSettingMerchantGroup

	err := db.Where(dbmodels.FeeMdrSettingMerchantGroup{
		ID: id,
	}).First(&data).Error

	return data, err
}

// Filter ..
func (repo *FeeMdrSettingMerchantGroupRepository) Filter(req dto.ReqFeeMdrSettingMerchantGroupDto) ([]dbmodels.FeeMdrSettingMerchantGroup, int, error) {
	db := GetDbCon()
	var res []dbmodels.FeeMdrSettingMerchantGroup
	var total int
	page := req.Page
	limit := req.Limit

	if req.IdMerchantGroup != 0 {
		db = db.Where("id_merchant_group = ?", req.IdMerchantGroup)
	}

	if req.MidBank != "" {
		db = db.Where("mid_bank = ?", req.MidBank)
	}

	if req.SecretID != "" {
		db = db.Where("secret_id = ?", req.SecretID)
	}

	if req.Bank != "" {
		db = db.Where("bank = ?", req.Bank)
	}

	if req.Tenor != "" {
		db = db.Where("tenor = ?", req.Tenor)
	}

	err := db.Limit(limit).Offset((page - 1) * limit).Order("id asc").Find(&res).Limit(-1).Offset(0).Count(&total).Error

	return res, total, err
}

// FindAll ..
func (repo *FeeMdrSettingMerchantGroupRepository) FindAll() ([]dbmodels.FeeMdrSettingMerchantGroup, int, error) {
	db := GetDbCon()
	var res []dbmodels.FeeMdrSettingMerchantGroup
	var total int

	err := db.Order("id asc").Find(&res).Count(&total).Error

	return res, total, err
}
