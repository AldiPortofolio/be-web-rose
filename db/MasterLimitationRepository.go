package db

import (
	"rose-be-go/constants/status_approval"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type MasterLimitationRepository struct {

}

func InitMasterLimitationRepository() *MasterLimitationRepository {
	return &MasterLimitationRepository{}
}


func (repo *MasterLimitationRepository) FilterLimitationMerchant(req dto.MasterLimitationReq) ([]dbmodels.LimitationMerchant, int, error) {
	db := GetDbCon()
	var res []dbmodels.LimitationMerchant
	var total int
	page := req.Page
	limit := req.Limit

	db = db.Where("status = ?", status_approval.STATUS_ACTIVE).Limit(req.Limit).Offset((page-1)*limit)

	if req.ByTime != "" {
		db = db.Where("by_time = ?", req.ByTime)
	}

	if req.ByGroupFilter != "" {
		db = db.Where("by_group ilike = ?", "%" + req.ByGroupFilter + "%")
	}

	if req.ProductType != "" {
       db = db.Where("product_type = ?", req.ProductType)
	}

	if req.ProductName != "" {
		db = db.Where("product_name ilike ?", "%" + req.ProductName + "%")
	}

	//err := db.Find(&res).Count(&total).Limit(limit).Offset((page-1)*limit).Limit(-1).Offset(0).Error
	err := db.Order("id desc").Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		return res, 0, err
	}

	return res, total, nil
}

func (repo *MasterLimitationRepository) SaveLimitationTemp(req dbmodels.LimitationMerchantTemp) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

func (repo *MasterLimitationRepository) GetGroupMerchant(req []string) ([]dbmodels.Merchant, error) {
	db := GetDbCon()

	var res []dbmodels.Merchant

	err := db.Where("merchant_outlet_id in (?)", req).Find(&res).Order("id asc").Error

	if err != nil {
		return res, err
	}

	return res, nil

}

func (repo *MasterLimitationRepository) SaveMasterLimitation(data *dbmodels.LimitationMerchant) error {
	db:= GetDbCon()

	err:= db.Save(&data).Error

	return err
}

func (repo *MasterLimitationRepository) GetMasterLimitationByID(id int64) (dbmodels.LimitationMerchant, error) {
	db := GetDbCon()

	var data dbmodels.LimitationMerchant

	err := db.Where(dbmodels.LimitationMerchant{
		ID: id,
	}).First(&data).Error

	return data, err
}

func (repo *MasterLimitationRepository) CheckQueue(id int64) bool {
	db := GetDbCon()
	total := 0
	var data []dbmodels.LimitationMerchantTemp

	err := db.Where("action_type in (0,1,4) and master_limitation_id != 0 and master_limitation_id = ?", id).Find(&data).Count(&total).Error

	if err != nil {
		return false
	}

	if total > 0 {
		return false
	}

	return true
}
