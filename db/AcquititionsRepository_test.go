package db

import (
	"encoding/json"
	"log"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"testing"
)

func TestAcquititionsRepository_FindByID(t *testing.T) {
	res, err := InitBannerRepository().FindByID(1)

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println(err)
}

func TestAcquititionsRepository_Filter(t *testing.T) {
	req := dto.ReqFilterAcquititionsDto{
		Limit:           10,
		Page:            1,
	}
	res, total, err := InitAcquititionsRepository().Filter(req)

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println(total, err)
}

func TestAcquititionsRepository_Save(t *testing.T) {
	req := dbmodels.Acquititions{
		ID:              0,
		MerchantType     : "test merchant type",
		MerchantGroupId  : 55,
		MerchantCategory : "test merchant category",
		Name             : "test",
		LogoUrl          : "http://test",
		RegisterUsingId  : true,
		Sequence         : 5,
		ShowInApp        : "Active",
		SalesRetails     : "10",
		BusinessType     : "test business type",

		
	}
	err := InitAcquititionsRepository().Save(&req)

	log.Println(err)
}