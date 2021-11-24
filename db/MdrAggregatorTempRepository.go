package db

import (
	"rose-be-go/constants/status_approval"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type MdrAggregatorTempRepository struct {

}

func InitMdrAggregatorTempRepository() *MdrAggregatorTempRepository {
	return &MdrAggregatorTempRepository{}
}

func (repo *MdrAggregatorTempRepository) SaveMdrAggregatorTemp(data *dbmodels.MdrAggregatorTemp) error {
	db:= GetDbCon()

	err:= db.Save(&data).Error

	return err
}

func (repo *MdrAggregatorTempRepository) GetMdrAggregatorTempByID(id int64) (dbmodels.MdrAggregatorTemp, error) {
	db := GetDbCon()

	var data dbmodels.MdrAggregatorTemp

	err := db.Where(dbmodels.MdrAggregatorTemp{
		ID: id,
	}).First(&data).Error

	return data, err
}

func (repo *MdrAggregatorTempRepository) FilterMdrAggregatorTemp(req dto.ReqFilterDto) ([]dbmodels.MdrAggregatorTemp, int, error) {
	db := GetDbCon()

	var data []dbmodels.MdrAggregatorTemp
	var total int

	err := db.Where("action_type in (?,?,?)",status_approval.CREATE, status_approval.EDIT, status_approval.DELETE).Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Order("id desc").Find(&data).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		return data, 0, err
	}

	return data, total, nil
}