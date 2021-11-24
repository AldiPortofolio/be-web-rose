package db

import (
	"encoding/json"
	"log"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"testing"
)

func TestFilterDataAggregatorDetail(t *testing.T) {
	req := dto.ReqFilterDto{
		Limit: 2,
		Page: 1,

	}
	res, total, err := InitMerchantAggregatorDetailRepository().FilterDataAggregatorDetail(req)

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println(total, err)
}

func TestFilterListAggregatorDetail(t *testing.T) {
	req := dto.ReqFilterDto{
		MidAggregator: "687687",
	}
	res, err := InitMerchantAggregatorDetailRepository().FilterListAggregatorDetail(req)

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println( err)
}

func TestSaveMerchantAggregatorDetailTemp(t *testing.T) {

	req := dbmodels.MerchantAggregatorDetailTemp{
		ID: 0,
		ActionType: 0,
		MidAggregator: "123",
		MidMerchant: "456789",
	}

	err := InitMerchantAggregatorDetailRepository().SaveMerchantAggregatorDetailTemp(&req)
	log.Println("err ->", err)
}


func TestFindMerchantAggregatorDetailTempByIdDetail(t *testing.T) {
	res, err:= InitMerchantAggregatorDetailRepository().FindMerchantAggregatorDetailTempByIdDetail(1)
	log.Println(res)
	log.Println(err)
}


func TestFindMerchantAggregatorDetailTempByMidAggregator(t *testing.T) {
	req := dto.ReqFilterDto{
		MidAggregator: "12345678",
		Page:0,
		Limit:2,
	}
	res, total, err:= InitMerchantAggregatorDetailRepository().FindMerchantAggregatorDetailTempByMidAggregator(req)
	log.Println(res)
	log.Println(total)
	log.Println(err)
}

func TestFindListMerchantAggregatorApproval(t *testing.T) {
	req := dto.ReqFilterDto{
		Name:"ac",
		Page:0,
		Limit:2,
	}
	res, total, err:= InitMerchantAggregatorDetailRepository().FindListMerchantAggregatorApproval(req)
	log.Println(res)
	log.Println(total)
	log.Println(err)
}

func TestFindMerchantApprovalByMidAggregator(t *testing.T) {
	req := dto.ReqFilterDto{
		MidAggregator:"O12345678",
		Page:0,
		Limit:10,
	}
	res, total, err:= InitMerchantAggregatorDetailRepository().FindMerchantApprovalByMidAggregator(req)
	log.Println(res)
	log.Println(total)
	log.Println(err)
}




func TestFindMerchantAggregatorDetailTempById(t *testing.T)  {
	res, err:= InitMerchantAggregatorDetailRepository().FindMerchantAggregatorDetailTempById(10)
	log.Println(res)
	log.Println(err)
}

func TestFindMerchantAggregatorDetailById(t *testing.T)  {
	res, err:= InitMerchantAggregatorDetailRepository().FindMerchantAggregatorDetailById(3)
	log.Println(res)
	log.Println(err)
}


func TestSaveMerchantAggregatorDetail(t *testing.T) {

	req := dbmodels.MerchantAggregatorDetail{
		ID: 0,
		ActionType: 0,
		MidAggregator: "123",
		MidMerchant: "456789",
	}

	err := InitMerchantAggregatorDetailRepository().SaveMerchantAggregatorDetail(&req)
	log.Println("err ->", err)
}
