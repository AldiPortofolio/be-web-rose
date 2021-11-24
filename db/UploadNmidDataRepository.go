package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"rose-be-go/models/dbmodels"
	"strings"
	"time"
)

type UploadNmidDataRepository struct {

}

func InitUploadNmidDataRepository() *UploadNmidDataRepository {
	return &UploadNmidDataRepository{}
}

var (
	transactionDateLayout string
)

func init() {
	transactionDateLayout = "2006-01-02T15:04:05.000000Z07:00"
}

// saveNMID
func (repo *UploadNmidDataRepository) SaveNmid(upload_nmid *dbmodels.UploadNmidData) error{

	err:= DbCon.Save(&upload_nmid).Error
	return err
}

// GetDataUploadNmid ..
func (repo *UploadNmidDataRepository) GetDataUploadNmid(req dbmodels.UploadNmidData) ([]dbmodels.UploadNmidData, int, error) {
	db := GetDbCon()

	page := req.Page
	limit := req.Limit

	var dataUploadNmid []dbmodels.UploadNmidData
	var total int

	repo.setWhereStartDate(&db, req.StartDate)

	repo.setWhereEndDate(&db, req.EndDate)

	err := db.Limit(limit).Offset((page-1) * limit).Order("id desc").Order("id").Find(&dataUploadNmid).Limit(-1).Offset(0).Count(&total).Error // query

	// cara lain kalo QUEry (bukan pake paetik satu, tapi di bawah esc)
	//err := db.Exec(`query nya disni aja`).Find(&dataUploadNmid).Error

	if err != nil {
		fmt.Println("<<< Error get data upload nmid >>>")
		return dataUploadNmid, 0, err
	}

	fmt.Println("<<< Error get data upload nmid >>> {}", dataUploadNmid)

	return dataUploadNmid, total, nil
}

func (repo *UploadNmidDataRepository) setWhereStartDate(db **gorm.DB, startDate string) {
	if strings.TrimSpace(startDate) != "" {
		if date, err := time.Parse("2006-01-02", startDate); err != nil {
			fmt.Println("Failed to parse request start date to time:", err)
		} else {
			*db = (*db).Where("date >= ?", date.Format(transactionDateLayout))
		}
	}
}

func (repo *UploadNmidDataRepository) setWhereEndDate(db **gorm.DB, endDate string) {
	if strings.TrimSpace(endDate) != "" {
		if date, err := time.Parse("2006-01-02", endDate); err != nil {
			fmt.Println("Failed to parse request start date to time:", err)
		} else {
			*db = (*db).Where("date <= ?", date.Format(transactionDateLayout))
		}
	}
}

