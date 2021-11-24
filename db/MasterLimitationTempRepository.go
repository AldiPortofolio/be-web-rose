package db

import (
	"rose-be-go/constants/status_approval"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type MasterLimitationTempRepository struct {

}

func InitMasterLimitationTempRepository() *MasterLimitationTempRepository {
	return &MasterLimitationTempRepository{}
}

func (repo *MasterLimitationTempRepository) SaveMasterLimitationTemp(data *dbmodels.LimitationMerchantTemp) error {
	db:= GetDbCon()

	err:= db.Save(&data).Error

	return err
}

func (repo *MasterLimitationTempRepository) GetMasterLimitationTempByID(id int64) (dbmodels.LimitationMerchantTemp, error) {
	db := GetDbCon()

	var data dbmodels.LimitationMerchantTemp

	err := db.Where(dbmodels.LimitationMerchantTemp{
		ID: id,
	}).First(&data).Error

	return data, err
}

func (repo *MasterLimitationTempRepository) FilterMasterLimitationTemp(req dto.ReqFilterDto) ([]dbmodels.LimitationMerchantTemp, int, error) {
	db := GetDbCon()

	var data []dbmodels.LimitationMerchantTemp
	var total int

	err := db.Where("action_type in (?,?,?)",status_approval.CREATE, status_approval.EDIT, status_approval.DELETE).Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Order("id asc").Find(&data).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		return data, 0, err
	}

	return data, total, nil
}
