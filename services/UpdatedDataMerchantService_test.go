package services

import (
	"net/http"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	"testing"
)

func TestInitUpdatedDataMerchantService(t *testing.T) {

	req := dto.ReqUpdatedDataMerchantDto{
		Page:1,
		Limit: 10,
	}
	var res models.Response
	logs := logger.InitLogs(&http.Request{})
	InitUpdatedDataMerchantService(logs).Filter(req, &res)


}

