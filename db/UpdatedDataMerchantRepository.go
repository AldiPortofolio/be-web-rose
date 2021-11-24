package db

import (
	"fmt"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type UpdatedDataMerchantRepository struct {

}

func InitUpdatedDataMerchantRepository() *UpdatedDataMerchantRepository {
	return &UpdatedDataMerchantRepository{}
}

func (r *UpdatedDataMerchantRepository) Save(data *dbmodels.UpdatedDataMerchant) (err error) {
	db := GetDbCon()
	err = db.Save(&data).Error
	return
}

func (r *UpdatedDataMerchantRepository) FindByID(id int64) (dbmodels.UpdatedDataMerchant, error) {

	db := GetDbCon()
	var res dbmodels.UpdatedDataMerchant
	err := db.Where("id = ?", id).First(&res).Error
	return res, err


}

func (r *UpdatedDataMerchantRepository) FindByIDAndStatusPending(mid string) (dbmodels.UpdatedDataMerchant, error) {

	db := GetDbCon()
	var res dbmodels.UpdatedDataMerchant
	err := db.Where("mid = ? , status = ?", mid, "Pending").First(&res).Error
	return res, err


}


func (r *UpdatedDataMerchantRepository)Filter(req dto.ReqUpdatedDataMerchantDto) (res []dbmodels.UpdatedDataMerchant, total int, err error) {
	fmt.Println("<< UpdatedDataMerchantRepository - Filter >>")
	db := GetDbCon()

	limit := req.Limit
	page := req.Page


	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}

	if req.StoreName != "" {
		db = db.Where("store_name ilike ?", "%" + req.StoreName +"%")

	}

	if req.LoanBankCode != "" {
		db = db.Where("loan_bank_code ilike ?", "%" + req.LoanBankCode +"%")

	}

	if req.Mid != "" {
		db = db.Where("mid ilike ?", "%" + req.Mid +"%")

	}

	err =  db.Limit(limit).Offset((page-1)*limit).Order("updated_at asc").Find(&res).Limit(-1).Offset(0).Count(&total).Error




	return res, total, err
}