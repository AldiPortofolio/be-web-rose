package db

import (
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type ReportCrmPenugasanRepository struct {

}

func InitReportCrmPenugasanRepository() *ReportCrmPenugasanRepository {
	return &ReportCrmPenugasanRepository{}
}


func (repo *ReportCrmPenugasanRepository) FilterPaging(req dto.ReqGetReporCrmPenugasanDto) ([]dbmodels.ReportCrmPenugasan, int, error) {

	db := DbCon
	page := req.Page
	limit := req.Limit

	var dataReportFinish []dbmodels.ReportCrmPenugasan
	var total int

	err := db.Limit(limit).Offset((page - 1) * limit).Order("id desc").Find(&dataReportFinish).Limit(-1).Offset(0).Count(&total).Error // query

	return dataReportFinish, total, err


}