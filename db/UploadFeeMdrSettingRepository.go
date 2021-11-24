package db

import (
	"errors"
	"fmt"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type UploadFeeMdrSettingRepository struct {

}

func InitUploadFeeMdrSettingRepository() *UploadFeeMdrSettingRepository {
	return &UploadFeeMdrSettingRepository{}
}

func (repo *UploadFeeMdrSettingRepository)Save(data *dbmodels.UploadFeeMdrSetting) (error) {
	db := GetDbCon()
	if err := db.Save(&data).Error; err != nil {
		return  errors.New("Gagal Insert Upload Fee Mdr Setting " + err.Error())
	}
	return nil
}

func (repo *UploadFeeMdrSettingRepository) GetDataUpload(req dto.ReqUploadFeeMdrSettingDto) ([]dbmodels.UploadFeeMdrSetting, int, error) {
	db := GetDbCon()

	page := req.Page
	limit := req.Limit

	var uploadMerchant []dbmodels.UploadFeeMdrSetting
	var total int


	err := db.Limit(limit).Offset((page-1) * limit).Order("id desc").Order("id").Find(&uploadMerchant).Limit(-1).Offset(0).Count(&total).Error // query

	if err != nil {
		fmt.Println("<<< Error get data upload fee mdr setting >>>", err)
		return uploadMerchant, 0, err
	}

	return uploadMerchant, total, nil
}