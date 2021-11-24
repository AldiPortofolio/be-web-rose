package db

import (
	"github.com/jinzhu/gorm"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"strings"
)

type DepositBankRepository struct {

}

func InitDepositBankRepository() *DepositBankRepository {
	return &DepositBankRepository{}
}

func (repo *DepositBankRepository) GetAllInfo(req dto.ReqDepositBank) ([]dbmodels.DepositBank, int, error) {
	db := GetDbCon()

	var total int
	var data []dbmodels.DepositBank
	search := "%" + req.AccountName + "%"
	setWhereAccountName(&db, search)
	setLimitOffset(&db, req.Page, req.Limit)

	err := db.Order("updated_at desc").Find(&data).Limit(-1).Offset(0).Count(&total).Error // query

	return data, total, err
}

func setLimitOffset(db **gorm.DB, requestPage int, requestLimit int) {
	count := 10
	if requestLimit > 0 {
		count = requestLimit
	}

	offset := 0
	if requestPage > 0 {
		offset = (requestPage - 1) * count
	}

	*db = (*db).Limit(count).Offset(offset)
}

func setWhereAccountName(db **gorm.DB, name string) {
	if strings.TrimSpace(name) != "" {
		*db = (*db).Where("account_name ILIKE ?", name)
	}
}
