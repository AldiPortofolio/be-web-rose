package db

import (
	"log"
	"net/http"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"testing"
)

func TestInitLoanProductMaintenanceRepository_save(t *testing.T) {

	logs := logger.InitLogs(&http.Request{})
	req:= dbmodels.LoanProductMaintenance{
		BankCode: "012",
		Status: "y",
		BankName:"BCA",
		AdminFeeType: "percentage",
		AdminFeeValue: 1.3,


	}
	err :=InitLoanProductMaintenanceRepository(logs).Save(&req)
	log.Println(err)
}

func TestInitLoanProductMaintenanceRepository(t *testing.T) {
	logs := logger.InitLogs(&http.Request{})
	req := dto.ReqLoanProductMaintenanceDto{
		Page:2,
		Limit: 1,

	}

	res, total, err:= InitLoanProductMaintenanceRepository(logs).Filter(req)
	log.Println(res)
	log.Println(total)
	log.Println(err)

}