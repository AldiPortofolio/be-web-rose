package db

import (
	"log"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"testing"
)

func TestFilterMasterLimitationTemp(t *testing.T) {

	req := dto.ReqFilterDto{
		Limit:10,
		Page: 1,
	}
	res,total,err := InitMasterLimitationTempRepository().FilterMasterLimitationTemp(req)

	log.Println(res)
	log.Println(total)
	log.Println(err)
}

func TestGetMasterLimitationTempByID(t *testing.T) {
	res, err := InitMasterLimitationTempRepository().GetMasterLimitationTempByID(10)
	log.Println(res, err)
}

func TestSaveMasterLimitationTemp(t *testing.T) {
	data := dbmodels.LimitationMerchantTemp{

	}
	err := InitMasterLimitationTempRepository().SaveMasterLimitationTemp(&data)
	log.Println(err)
}