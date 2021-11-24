package db

import (
	"log"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"testing"
	"time"
)

func TestMdrBankRepository_Save(t *testing.T) {
	req := dbmodels.MdrBank{
		UpdatedBy:"syandi",
		UpdatedAt: time.Now(),
		Status: "y",
		BankCode: "bni",
		BankName: "Bank Negara Indonesia 46",


	}
	InitMdrBankRepository().Save(&req)
}

func TestMdrBankRepository_Filter(t *testing.T) {
	req := dto.ReqMdrDto{
		Limit: 10,
		Page: 1,
	}
	res, total, err := InitMdrBankRepository().Filter(req)
	log.Println(res)
	log.Println(total)
	log.Println(err)
}

func TestMdrBankRepository_FindAll(t *testing.T) {
	res, err := InitMdrBankRepository().FindAll()
	log.Println(res)
	log.Println(err)

}

func TestMdrBankRepository_FindByID(t *testing.T) {
	res, err := InitMdrBankRepository().FindByID(10)
	log.Println(res)
	log.Println(err)
}