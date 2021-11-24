package db

import (
	"log"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"testing"
	"time"
)

func TestInitMdrTenorRepository_save(t *testing.T) {

	req := dbmodels.MdrTenor{
		Seq: 1,
		TenorCode:"fullpayment",
		TenorName: "Full Payment",
		UpdatedAt: time.Now(),
		Status: "y",
		UpdatedBy:"syandi",


	}
	err := InitMdrTenorRepository().Save(&req)
	log.Println(err)
	log.Println(req	)
}

func TestMdrTenorRepository_FindByID(t *testing.T) {

	res, err := InitMdrTenorRepository().FindByID(3)
	log.Println(err)
	log.Println(res)
}

func TestMdrTenorRepository_Filter(t *testing.T) {
	req := dto.ReqMdrTenorDto{
		Limit: 10,
		Page: 1,
	}
	res, total, err := InitMdrTenorRepository().Filter(req)
	log.Println(res)
	log.Println(total)
	log.Println(err)
}

func TestMdrTenorRepository_FindAll(t *testing.T) {
	res, err :=InitMdrTenorRepository().FindAll()
	log.Println(res)
	log.Println(err)
}