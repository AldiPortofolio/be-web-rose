package db

import (
	"log"
	"net/http"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"testing"
	"time"

	"ottodigital.id/library/logger/v2"
)

func TestInitSettlementConfigRepository(t *testing.T) {

	logs := logger.InitLogs(&http.Request{})

	req := dto.ReqSettlementConfigDto{
		Page: 1,
		Limit: 10,
		Code: "014",
	}

	res, total, err := InitSettlementConfigRepository(logs).Filter(req)

	log.Println(err)
	log.Println(total)
	log.Println(res)
}

func TestInitSettlementConfigRepository2(t *testing.T) {

	logs := logger.InitLogs(&http.Request{})

	req := dbmodels.SettlementConfig{
		BankCode:"014",
		Status:"Active",
		SettlementType: "BCA",
		UpdatedBy: "abdul",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

		
	}

	err := InitSettlementConfigRepository(logs).Save(&req)

	log.Println(err)
}