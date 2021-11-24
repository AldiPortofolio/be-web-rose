package db

import (
	"log"
	"rose-be-go/models/dto"
	"testing"
)

func TestInitAkuisisiSfaRepository(t *testing.T) {
	req := dto.ReqAkuisisiSfa{
		Key: "",
		Page: 1,
		Limit: 10,
	}

	res, total, err:=InitAkuisisiSfaRepository().Filter(req)
	log.Println(err)
	log.Println(total)
	log.Println(res)
}