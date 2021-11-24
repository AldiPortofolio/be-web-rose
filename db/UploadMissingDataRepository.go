package db

import (
	"errors"
	"fmt"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type UploadMissingDataRepository struct {

}

func InitUploadMissingDataRepository() *UploadMissingDataRepository  {
	return &UploadMissingDataRepository{}
}

func (repo *UploadMissingDataRepository) Save(uploadMerchant *dbmodels.UploadMissingData) (error) {
	db := GetDbCon()
	if err := db.Save(&uploadMerchant).Error; err != nil {
		return  errors.New("Gagal Insert Upload Missing Data " + err.Error())
	}
	return nil
}



func (repo *UploadMissingDataRepository) GetDataUploadMerchant(req dto.ReqUploadMissingDataDto) ([]dbmodels.UploadMissingData, int, error) {
	db := GetDbCon()

	page := req.Page
	limit := req.Limit

	var uploadMerchant []dbmodels.UploadMissingData
	var total int


	err := db.Limit(limit).Offset((page-1) * limit).Order("id desc").Order("id").Find(&uploadMerchant).Limit(-1).Offset(0).Count(&total).Error // query

	if err != nil {
		fmt.Println("<<< Error get data upload missing data >>>", err)
		return uploadMerchant, 0, err
	}

	return uploadMerchant, total, nil
}
