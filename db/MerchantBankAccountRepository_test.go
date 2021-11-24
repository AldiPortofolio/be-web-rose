package db

import (
	"log"
	"net/http"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/models/dto"
	"testing"
)

func TestInitMerchantBankAccountRepository2(t *testing.T) {
	logs := logger.InitLogs(&http.Request{})


	res, err:= InitMerchantBankAccountRepository(logs).GetDataByMid("2123213")
	if err != nil {
		log.Println(err)
	}

	log.Println(res)
}

func TestInitMerchantBankAccountRepository(t *testing.T) {
	logs := logger.InitLogs(&http.Request{})

	req := dto.ReqMerchantBankAccountDto{
		Limit: 10,
		Page: 1,
	}
	res, total, err:= InitMerchantBankAccountRepository(logs).FilterApproval(req)
	if err != nil {
		log.Println(err)
	}

	log.Println(total)
	log.Println(res)
}