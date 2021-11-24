package db

import (
	"log"
	"rose-be-go/models/dto"
	"testing"
)

func TestInitQrPrePrintedRepository(t *testing.T) {

	req:= dto.ReqQrPrePrintedDto{
		Page:1,
		Limit:10,
		StartDate: "2020-09-17",
		EndDate: "2020-09-17",
	}
	res, total, err:=InitQrPrePrintedRepository().GetData(req)
	log.Println(res)
	log.Println(total)
	log.Println(err)
}