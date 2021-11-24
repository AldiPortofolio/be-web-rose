package db

import (
	"log"
	"net/http"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"testing"
)

func TestInitValidationCodeMasterTagRepository(t *testing.T) {
	logs := logger.InitLogs(&http.Request{})

	req := dbmodels.ValidationCodeMasterTag{
		MasterTagID:1,
		ValidationCodeID: 1,

	}

	err:=InitValidationCodeMasterTagRepository(logs).Save(&req)
	log.Println(err)
}

func TestInitValidationCodeMasterTagRepository2(t *testing.T) {
	logs := logger.InitLogs(&http.Request{})

	req := dto.ReqValidationCodeMasterTagDto{
		Page:1,
		Limit:10,
		ValidationCodeID: 1,
	}
	res, total, err:=InitValidationCodeMasterTagRepository(logs).Filter(req)
	log.Println(err)
	log.Println(total)
	log.Println(res)
}