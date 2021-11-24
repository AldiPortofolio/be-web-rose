package db

import (
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type InstructionListRepository struct {
}

func InitInstructionListRepository() *InstructionListRepository {
	return &InstructionListRepository{}
}

func (repo *InstructionListRepository) FindByID(id int64) (dbmodels.InstructionList, error) {
	db := GetDbCon()

	var data dbmodels.InstructionList

	err := db.Where(dbmodels.InstructionList{
		ID: id,
	}).First(&data).Error

	return data, err
}

func (repo *InstructionListRepository) Save(req *dbmodels.InstructionList) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

func (r *InstructionListRepository) Delete(id int) error {
	db := GetDbCon()
	var instructionList dbmodels.InstructionList
	err := db.Where("id = ?", id).Delete(&instructionList).Error

	if err != nil {
		return err
	}

	return nil
}
func (repo *InstructionListRepository) Filter(req dto.ReqInstructionListDto) ([]dbmodels.InstructionList, int, error) {
	db := GetDbCon()
	var res []dbmodels.InstructionList
	var total int
	page := req.Page
	limit := req.Limit

	err := db.Limit(limit).Offset((page - 1) * limit).Order("id DESC").Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		return res, 0, err
	}

	return res, total, nil
}
