package db

import (
	"log"
	"rose-be-go/models/dto"
	"testing"
)

func TestInitUploadMerchantRepository(t *testing.T) {
	req := dto.ReqUploadMerchant{
		Limit: 2,
		Page: 2,
		StartDate: "2020-06-08",
		EndDate: "2020-06-08",
	}
	res, total, err:= InitUploadMerchantRepository().GetDataUploadMerchant(req)
	log.Println(res)
	log.Println(total)
	log.Println(err)
}
