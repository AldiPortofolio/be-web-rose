package db

import (
	"fmt"
	"rose-be-go/models/dbmodels"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

func init() {
	transactionDateLayout = "2006-01-02T15:04:05.000000Z07:00"
}

// GetDataReportRejected ..
func GetReportExportMerchant(req dbmodels.ReportExportMerchant) ([]dbmodels.ReportExportMerchant, int, error) {
	db := GetDbCon()

	page := req.Page
	limit := req.Limit

	var dataReportExportMerchant []dbmodels.ReportExportMerchant
	var total int

	GetDateExportMerchant(&db, req.StartDate, req.EndDate)

	err := db.Limit(limit).Offset((page - 1) * limit).Order("id desc").Find(&dataReportExportMerchant).Limit(-1).Offset(0).Count(&total).Error // query


	if err != nil {
		fmt.Println("<<< Error get data Report Rejected >>>")
		return dataReportExportMerchant, 0, err
	}

	fmt.Println("<<< Error get data Report Export Merchant >>> {}", dataReportExportMerchant)

	return dataReportExportMerchant, total, nil
}

// GetDateExportMerchant ...
func GetDateExportMerchant(db **gorm.DB, startDate string, endDate string) {

	if strings.TrimSpace(startDate) != "" && strings.TrimSpace(endDate) != "" {
		dateStart, errStart := time.Parse("2006-01-02 15:04:05", startDate)
		if errStart != nil {
			fmt.Println("Failed to parse request start date to time:", errStart)
		}

		dateEnd, errEnd := time.Parse("2006-01-02 15:04:05", endDate)
		if errEnd != nil {
			fmt.Println("Failed to parse request start date to time:", errEnd)
		}

		*db = (*db).Where("transaction_date >= ? and transaction_date <= ?", dateStart.Format(transactionDateLayout), dateEnd.Format(transactionDateLayout))
	}
}
