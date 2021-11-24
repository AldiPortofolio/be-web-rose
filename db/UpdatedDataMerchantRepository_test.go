package db

import (
	"log"
	"rose-be-go/models/dto"
	"testing"
)

func TestInitUpdatedDataMerchantRepository(t *testing.T) {

	req := dto.ReqUpdatedDataMerchantDto{
		Page:1,
		Limit: 10,
	}

	res, total, _:=InitUpdatedDataMerchantRepository().Filter(req)
	log.Println(total)
	log.Println(res)
}