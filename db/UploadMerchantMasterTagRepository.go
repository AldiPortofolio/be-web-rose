package db

import (
	"errors"
	"fmt"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type UploadMerchantMasterTagRepository struct {

}

func InitUploadMerchantMasterTagRepository() *UploadMerchantMasterTagRepository {
	return &UploadMerchantMasterTagRepository{}
}

func (repo *UploadMerchantMasterTagRepository) Save(uploadMerchant *dbmodels.UploadMerchantMasterTag) (error) {
	db := GetDbCon()
	if err := db.Save(&uploadMerchant).Error; err != nil {
		return  errors.New("Gagal Insert Upload Merchant Master Tag " + err.Error())
	}
	return nil
}

func (repo *UploadMerchantMasterTagRepository) GetDataUploadMerchant(req dto.ReqUploadMerchantMasterTagDto) ([]dbmodels.UploadMerchantMasterTag, int, error) {
	db := GetDbCon()

	page := req.Page
	limit := req.Limit

	var uploadMerchant []dbmodels.UploadMerchantMasterTag
	var total int


	err := db.Limit(limit).Offset((page-1) * limit).Order("id desc").Order("id").Find(&uploadMerchant).Limit(-1).Offset(0).Count(&total).Error // query

	if err != nil {
		fmt.Println("<<< Error get data upload merchant Master Tag >>>", err)
		return uploadMerchant, 0, err
	}

	return uploadMerchant, total, nil
}
