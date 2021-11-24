package db

import (
	"ottodigital.id/library/logger/v2"
	"rose-be-go/models/dbmodels"
	
)

type MerchantSettlementConfigRepository struct {
	Ottolog logger.OttologInterface

}

func InitMerchantSettlementConfigRepository(logs logger.OttologInterface) *MerchantSettlementConfigRepository {
	return &MerchantSettlementConfigRepository{
		Ottolog:logs,
	}
}

func (r *MerchantSettlementConfigRepository)Save(req *dbmodels.MerchantSettlementConfig) error {
	db := GetDbCon()

	if err:= db.Save(&req).Error; err !=nil{
		r.Ottolog.Error("Error save to db " + err.Error())
		return err
	}

	return nil
}

func (r *MerchantSettlementConfigRepository)Update(req *dbmodels.MerchantSettlementConfig) error {
	db := GetDbCon()

	var res dbmodels.MerchantSettlementConfig

	if err:= db.Model(&res).Where("mid = ?", req.Mid).Updates(&req).Error; err !=nil{
		r.Ottolog.Error("Error save to db " + err.Error())
		return err
	}

	return nil
}

func (r *MerchantSettlementConfigRepository)FindByMid(mid string) (dbmodels.MerchantSettlementConfig, error) {
	db := GetDbCon()

	var res dbmodels.MerchantSettlementConfig


	if err:= db.Where("mid = ?", mid).First(&res).Error; err !=nil{
		r.Ottolog.Error("Error find data " + err.Error())
		return res, err
	}

	return res, nil
}

func (r *MerchantSettlementConfigRepository)FindById(id int64) (dbmodels.SettlementDetail, error) {
	db := GetDbCon()

	var res dbmodels.SettlementDetail


	if err:= db.Where("id = ?", id).First(&res).Error; err !=nil{
		r.Ottolog.Error("Error find data " + err.Error())
		return res, err
	}

	return res, nil
}
