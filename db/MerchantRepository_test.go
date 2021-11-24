package db

import (
	"log"
	"rose-be-go/models/dto"
	"testing"
)

func TestInitMerchantRepository(t *testing.T) {
	req := dto.ReqMerchantDto{
		Page:1,
		Limit:10,
		StoreName: "SYANDI",
		MerchantGroupID:55,
	}
	res, total, err:= InitMerchantRepository().Filter(req)

	log.Println(res)
	log.Println(total)
	log.Println(err)
}

func TestInitMerchantRepository2(t *testing.T) {
	data, err :=InitMerchantRepository().FindTanggalAkuisisi("080944433330")
	log.Println(err)
	log.Println(data)
}