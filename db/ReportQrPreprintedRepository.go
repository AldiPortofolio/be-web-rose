package db

import (
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type ReportQrPreprintedRepository struct {

}

func InitReportQrPreprintedRepository() *ReportQrPreprintedRepository {
	return &ReportQrPreprintedRepository{}
}

func (repo *ReportQrPreprintedRepository) FilterPaging(req dto.ReqGetReportQrPreprintedDto) ([]dbmodels.ReportQrPreprinted, int, error) {

	db := DbCon
	page := req.Page
	limit := req.Limit

	var dataReportFinish []dbmodels.ReportQrPreprinted
	var total int

	err := db.Limit(limit).Offset((page - 1) * limit).Order("id desc").Find(&dataReportFinish).Limit(-1).Offset(0).Count(&total).Error // query

	return dataReportFinish, total, err


}
