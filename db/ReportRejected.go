package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"rose-be-go/models/dbmodels"
	"strings"
	"time"
)

func init() {
	transactionDateLayout = "2006-01-02T15:04:05.000000Z07:00"
}

// GetDataReportRejected ..
func GetDataReportRejected(req dbmodels.ReportReject) ([]dbmodels.ReportReject, int, error) {
	db := GetDbCon()

	page := req.Page
	limit := req.Limit

	var dataReportReject []dbmodels.ReportReject
	var total int

	GetDateReject(&db, req.StartDate, req.EndDate)

	err := db.Limit(limit).Offset((page - 1) * limit).Order("id desc").Find(&dataReportReject).Limit(-1).Offset(0).Count(&total).Error // query


	if err != nil {
		fmt.Println("<<< Error get data Report Rejected >>>")
		return dataReportReject, 0, err
	}

	fmt.Println("<<< Error get data Report Rejected >>> {}", dataReportReject)

	return dataReportReject, total, nil
}

// GetDataPage ...
func GetDateReject(db **gorm.DB, startDate string, endDate string) {

	if strings.TrimSpace(startDate) != "" && strings.TrimSpace(endDate) != "" {
		dateStart, errStart := time.Parse("2006-01-02", startDate)
		if errStart != nil {
			fmt.Println("Failed to parse request start date to time:", errStart)
		}

		dateEnd, errEnd := time.Parse("2006-01-02", endDate)
		if errEnd != nil {
			fmt.Println("Failed to parse request start date to time:", errEnd)
		}

		*db = (*db).Where("start_date >= ? and end_date <= ?", dateStart.Format(transactionDateLayout), dateEnd.Format(transactionDateLayout))
	}
}
