package db

import (
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type ReportUpdatedDataMerchantRepository struct {

}

func InitReportUpdatedDataMerchantRepository() *ReportUpdatedDataMerchantRepository {
	return &ReportUpdatedDataMerchantRepository{}
}


func (repo *ReportUpdatedDataMerchantRepository) FilterPaging(req dto.ReqGetReportUpdatedDataMerchantDto) ([]dbmodels.ReportUpdatedDataMerchant, int, error) {

	db := DbCon
	page := req.Page
	limit := req.Limit

	var dataReportFinish []dbmodels.ReportUpdatedDataMerchant
	var total int

	err := db.Limit(limit).Offset((page - 1) * limit).Order("id desc").Find(&dataReportFinish).Limit(-1).Offset(0).Count(&total).Error // query

	return dataReportFinish, total, err


}