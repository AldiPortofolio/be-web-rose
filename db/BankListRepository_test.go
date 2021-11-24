package db

import (
	"log"
	"net/http"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"testing"
)

func TestInitBankListRepository(t *testing.T) {

	logs := logger.InitLogs(&http.Request{})

	req := dto.ReqBankListDto{
		Page: 1,
		Limit: 10,
		ShortName: "c",
	}

	res, total, err := InitBankListRepository(logs).Filter(req)

	log.Println(err)
	log.Println(total)
	log.Println(res)
}

func TestInitBankListRepository2(t *testing.T) {

	logs := logger.InitLogs(&http.Request{})

	req := dbmodels.BankList{
		ShortName:"dsfsdf",
		Code:"2344",
		Status: "y",
		Seq: 10,
	}

	err := InitBankListRepository(logs).Save(&req)

	log.Println(err)
}