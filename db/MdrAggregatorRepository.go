package db

import (
	"rose-be-go/constants/status_approval"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type MdrAggregatorRepository struct {
	
}

func InitMdrAggregatorRepository() *MdrAggregatorRepository {
	return &MdrAggregatorRepository{}
}


func (repo *MdrAggregatorRepository) CheckQueue(id int64) bool {
	db := GetDbCon()
	total := 0
	var data []dbmodels.MdrAggregatorTemp

	err := db.Where("action_type in (0,1,4) and mdr_aggregator_id != 0 and mdr_aggregator_id = ?", id).Find(&data).Count(&total).Error

	if err != nil {
		return false
	}

	if total > 0 {
		return false
	}

	return true
}

func (repo *MdrAggregatorRepository) SaveMdrAggregatorTemp(req dbmodels.MdrAggregatorTemp) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

func (repo *MdrAggregatorRepository) FilterMdrAggregator(req dto.ReqMdrAggragtorDto) ([]dbmodels.MdrAggregator, int, error) {
	db := GetDbCon()
	var res []dbmodels.MdrAggregator
	var total int
	page := req.Page
	limit := req.Limit

	db = db.Where("status = ?", status_approval.STATUS_ACTIVE)

	if req.GroupPartner != "" {
		db = db.Where("group_partner = ?", req.GroupPartner)
	}

	if req.MerchantCategory != "" {
		db = db.Where("merchant_category = ?", req.MerchantCategory)
	}

	if req.TransactionType != "" {
		db = db.Where("transaction_type = ?", req.TransactionType)
	}

	if req.MdrType != "" {
		db = db.Where("mdr_type = ?", req.MdrType)
	}

	if req.MidPartner != "" {
		db = db.Where("mid_partner = ?", req.MidPartner)
	}

	if req.MidMerchant != "" {
		db = db.Where("mid_merchant = ?", req.MidMerchant)
	}

	err := db.Limit(limit).Offset((page-1)*limit).Order("group_partner asc").Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		return res, 0, err
	}

	return res, total, nil
}


func (repo *MdrAggregatorRepository) GetMdrAggregatorByID(id int64) (dbmodels.MdrAggregator, error) {
	db := GetDbCon()

	var data dbmodels.MdrAggregator

	err := db.Where(dbmodels.MdrAggregator{
		ID: id,
	}).First(&data).Error

	return data, err
}

func (repo *MdrAggregatorRepository) SaveMdrAggregator(data *dbmodels.MdrAggregator) error {
	db:= GetDbCon()

	err:= db.Save(&data).Error

	return err
}
