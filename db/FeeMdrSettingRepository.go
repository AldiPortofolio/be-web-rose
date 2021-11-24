package db

import (
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

// FeeMdrSettingRepository
type FeeMdrSettingRepository struct {
	// struct attributes
}

// InitFeeMdrSettingRepository ..
func InitFeeMdrSettingRepository() *FeeMdrSettingRepository {
	return &FeeMdrSettingRepository{}
}

// Save ..
func (repo *FeeMdrSettingRepository) Save(req *dbmodels.FeeMdrSetting) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

// FindByID ..
func (repo *FeeMdrSettingRepository) FindByID(id int64) (dbmodels.FeeMdrSetting, error) {
	db := GetDbCon()

	var data dbmodels.FeeMdrSetting

	err := db.Where(dbmodels.FeeMdrSetting{
		ID: id,
	}).First(&data).Error

	return data, err
}

// Filter ..
func (repo *FeeMdrSettingRepository) Filter(req dto.ReqFeeMdrSettingDto) ([]dto.ResFeeMdrSettingDto, int, error) {
	db := GetDbCon()
	var res []dto.ResFeeMdrSettingDto
	var total int
	page := req.Page
	limit := req.Limit

	if req.MidMerchant != "" {
		db = db.Where("a.mid_merchant = ?", req.MidMerchant)
	}

	if req.MidBank != "" {
		db = db.Where("a.mid_bank = ?", req.MidBank)
	}

	if req.SecretID != "" {
		db = db.Where("a.secret_id = ?", req.SecretID)
	}

	if req.Bank != "" {
		db = db.Where("a.bank = ?", req.Bank)
	}

	if req.Tenor != "" {
		db = db.Where("a.tenor = ?", req.Tenor)
	}
	if req.StoreName != "" {
		db = db.Where("b.store_name ilike ?","%"+ req.StoreName+"%")
	}

	db = db.Table("fee_mdr_setting a").Select("a.*, b.store_name").
		Joins("left join merchant b on a.mid_merchant = b.merchant_outlet_id")
	err := db.Limit(limit).Offset((page - 1) * limit).Order("id asc").Scan(&res).Limit(-1).Offset(0).Count(&total).Error

	return res, total, err
}

// FindAll ..
func (repo *FeeMdrSettingRepository) FindAll() ([]dbmodels.FeeMdrSetting, int, error) {
	db := GetDbCon()
	var res []dbmodels.FeeMdrSetting
	var total int

	err := db.Order("id asc").Find(&res).Count(&total).Error

	return res, total, err
}
