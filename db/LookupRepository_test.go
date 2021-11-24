package db

import (
	"log"
	"rose-be-go/models/dto"
	"testing"
)

func TestInitLookupRepository(t *testing.T) {
	req := dto.ReqLookupDto{
		Page: 1,
		Limit: 10,
		LookupGroup:"GENDER",
	}
	data, total, err:=InitLookupRepository().Filter(req)
	log.Println(err)
	log.Println(total)
	log.Println(data)
}