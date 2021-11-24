package db

import (
	"fmt"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models"
)

type ReportQrDataRepository struct {

}

func InitReportQrDataRepository() *ReportQrDataRepository {
	return &ReportQrDataRepository{}
}
// GetDataReportQr ..
func (repo *ReportQrDataRepository) GetDataReportQr(req models.Pagination) ([]dbmodels.ReportQr, int, error) {
	db := GetDbCon()
	db.Debug().LogMode(true)

	page := req.Page
	limit := req.Limit

	var dataReportQr []dbmodels.ReportQr
	var total int

	err := db.Limit(limit).Offset((page - 1) * limit).Order("id desc").Find(&dataReportQr).Limit(-1).Offset(0).Count(&total).Error // query

	if err != nil {
		fmt.Println("<<< Error get data Report Qr >>>")
		return dataReportQr, 0, err
	}

	fmt.Println("<<< Error get data Report Qr >>> {}", dataReportQr)

	return dataReportQr, total, nil
}


func (repo *ReportQrDataRepository) SaveReportQR(reportQr *dbmodels.ReportQr) error{

	err:= DbCon.Save(&reportQr).Error
	return err
}