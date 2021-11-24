package db

import (
	"log"
	"rose-be-go/models/dto"
	"testing"
)

func TestFilterMerchantQrisStatus(t *testing.T) {

	req := dto.ReqFilterDto{
		Limit:10,
		Page: 1,
		Status:"1",
		RequestDate: "21-04-2020",
		InstallDate: "21-04-2020",

	}


	res,total,err := InitMerchantQrisStatusRepository().Filter(req)

	log.Println(res)
	log.Println(total)
	log.Println(err)
}

