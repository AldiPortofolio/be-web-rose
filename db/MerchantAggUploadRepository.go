package db

import (
	"fmt"
	"log"
	"rose-be-go/models/dbmodels"
)

type MerchantAggUploadRepository struct {

}

func InitMerchantAggUploadRepository() *MerchantAggUploadRepository {
	return &MerchantAggUploadRepository{}
}

func (repo *MerchantAggUploadRepository) SaveMerchantAggUpload(upload *dbmodels.MerchantAggUpload) error {
	err := DbCon.Save(&upload).Error
	return err
}

func (repo *MerchantAggUploadRepository) GetDataMerchantAggUpload(req dbmodels.MerchantAggUpload) ([]dbmodels.MerchantAggUpload, int, error) {
	db := GetDbCon()

	page := req.Page
	limit := req.Limit

	var merchantAggupload []dbmodels.MerchantAggUpload
	var total int

	uploadNmidDataRepository := InitUploadNmidDataRepository()

	uploadNmidDataRepository.setWhereStartDate(&db, req.StartDate)
	uploadNmidDataRepository.setWhereEndDate(&db, req.EndDate)

	err := db.Limit(limit).Offset((page-1) * limit).Order("id desc").Order("id").Find(&merchantAggupload).Limit(-1).Offset(0).Count(&total).Error // query

	if err != nil {
		fmt.Println("<<< Error get data upload nmid >>>")
		return merchantAggupload, 0, err
	}

	fmt.Println("<<< Error get data upload nmid >>> {}", merchantAggupload)

	return merchantAggupload, total, nil
}

func (repo *MerchantAggUploadRepository) CheckApproval(mid string) (int, error) {
	db := GetDbCon()

	var temp dbmodels.MerchantAggregatorDetailTemp
	var total int

	err := db.Where("mid_aggregator = ? and action_type in (1,3)", mid).Find(&temp).Count(&total).Error

	if err != nil {
		log.Println("Error ->", err)
		return 0, err
	}

	return total, nil
}
