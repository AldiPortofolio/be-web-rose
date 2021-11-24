package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
	"github.com/unknwon/com"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"strconv"
	"strings"
)

type ReportMagTransactionsRepository struct {
}

func InitReportMagTransactionsRepository() *ReportMagTransactionsRepository {
	return &ReportMagTransactionsRepository{}
}

// GetAllDataMagTransaction ...
func (repo *ReportMagTransactionsRepository) GetAllDataMagTransaction(req dto.ReqReportMagTransactionsDto) ([]dbmodels.MagTransactions, int, error) {
	db := GetDbOfinCon()

	var data []dbmodels.MagTransactions
	var total int

	if req.FilterBy != "" && req.Keyword != "" {
		key := "%" + req.Keyword + "%"
		setConditions(&db, req.FilterBy, key)
	}
	setLimit(&db, req.Page, req.Limit)
	err := db.Order("merchant_pay_time desc").Find(&data).Limit(-1).Offset(0).Count(&total).Error

	return data, total, err
}

func setLimit(db **gorm.DB, requestPage int, requestLimit int) {
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

func setConditions(db **gorm.DB, field string, name string) {
	if field == "user" {
		field = `mag_transactions."user"`
	}
	if strings.TrimSpace(name) != "" {
		query := fmt.Sprintf("%s ILIKE '%s'", field, name)
		log.Info("Query >> ", query)
		*db = (*db).Where(query)
	}
}

// Export ...
func (repo *ReportMagTransactionsRepository) Export(data []dbmodels.MagTransactions) ([]dbmodels.ExportMagTransactions) {
	var file []dbmodels.ExportMagTransactions

	for i, value := range data {
		no := strconv.Itoa(i + 1)

		row := dbmodels.ExportMagTransactions{
			No:                no,
			Id:                com.ToStr(value.Id),
			User:              value.User,
			Partner:           value.Partner,
			BillingId:         value.BillingId,
			Channel:           value.Channel,
			MerchantId:        value.MerchantId,
			TerminalId:        value.TerminalId,
			Amount:            value.Amount,
			Tip:               value.Tip,
			TotalAmount:       value.TotalAmount,
			ReqReferenceNo:    value.ReqReferenceNo,
			QrCreatedAt:       value.QrCreatedAt.Format("2006-01-02 15:04:05"),
			MerchantPayStatus: value.MerchantPayStatus,
			MerchantPayRef:    value.MerchantPayRef,
			MerchantPayTime:   value.MerchantPayTime.Format("2006-01-02 15:04:05"),
			Issuer:            value.Issuer,
			IssuerCustAccount: value.IssuerCustAccount,
			IssuerRef:         value.IssuerRef,
			MagBillingId:      value.MagBillingId,
		}

		file = append(file, row)

	}

	return file
}