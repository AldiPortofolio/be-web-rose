package db

import (
	"log"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"testing"
	"time"
)

func TestInitQrisConfigRepository_save(t *testing.T) {
	req := dbmodels.QrisConfig{
		IssuerName:"testing",
		TransactionType: "on_us",
		InstitutionID: "12312321",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err:= InitQrisConfigRepository().Save(&req)
	log.Println(err)
}

func TestInitQrisConfigRepository_FindByID(t *testing.T) {

	data,err:= InitQrisConfigRepository().FindByID(8)
	log.Println(data)
	log.Println(err)
}

func TestInitQrisConfigRepository_Filter(t *testing.T) {

	req := dto.ReqQrisConfigDto{
		Limit: 10,
		Page:2,
		//InstitutionID: "1",
	}

	data, total, err:= InitQrisConfigRepository().Filter(req)
	log.Println(data)
	log.Println(total)
	log.Println(err)
}