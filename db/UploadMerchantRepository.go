package db

import (
	"errors"
	"fmt"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"strings"
	"time"
)

type UploadMerchantRepository struct {

}

func InitUploadMerchantRepository() *UploadMerchantRepository {
	return &UploadMerchantRepository{}
}

func (repo *UploadMerchantRepository)Save(uploadMerchant *dbmodels.UploadMerchant) (error) {
	db := GetDbCon()
	if err := db.Save(&uploadMerchant).Error; err != nil {
		return  errors.New("Gagal Insert Upload Merchant " + err.Error())
	}
	return nil
}

func (repo *UploadMerchantRepository) GetDataUploadMerchant(req dto.ReqUploadMerchant) ([]dbmodels.UploadMerchant, int, error) {
	db := GetDbCon()

	page := req.Page
	limit := req.Limit

	var uploadMerchant []dbmodels.UploadMerchant
	var total int

	db.Where("date")

	//repo.setWhereStartDate(&db, req.StartDate)
	//repo.setWhereEndDate(&db, req.EndDate)

	if strings.TrimSpace(req.StartDate) != "" && strings.TrimSpace(req.EndDate) != "" {
		dateStart, errStart := time.Parse("2006-01-02", req.StartDate)
		if errStart != nil {
			fmt.Println("Failed to parse request start date to time:", errStart)
		}

		dateEnd, errEnd := time.Parse("2006-01-02", req.EndDate)
		dateEnd = dateEnd.Add(time.Hour*23).Add(time.Minute*59).Add(time.Second*59)
		if errEnd != nil {
			fmt.Println("Failed to parse request start date to time:", errEnd)
		}

		db = db.Where("date >= ? and date <= ?", dateStart.Format(transactionDateLayout), dateEnd.Format(transactionDateLayout))
	}

	err := db.Limit(limit).Offset((page-1) * limit).Order("id desc").Order("id").Find(&uploadMerchant).Limit(-1).Offset(0).Count(&total).Error // query

	if err != nil {
		fmt.Println("<<< Error get data upload merchant >>>", err)
		return uploadMerchant, 0, err
	}

	return uploadMerchant, total, nil
}

