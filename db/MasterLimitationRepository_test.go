package db

import (
	"log"
	"rose-be-go/models/dbmodels"
	"testing"
)

func TestSaveMasterLimitation(t *testing.T) {
	data := dbmodels.LimitationMerchant{

	}
	err := InitMasterLimitationRepository().SaveMasterLimitation(&data)
	log.Println(err)
}
