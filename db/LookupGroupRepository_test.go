package db

import (
	"log"
	"net/http"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"testing"
)

func TestInitLookupGroupRepository(t *testing.T) {
	logs := logger.InitLogs(&http.Request{})
	req := dbmodels.LookupGroup{
		Name: "lookupgroupnamebaru",
	}
	InitLookupGroupRepository(logs).Save(&req)
}

func TestInitLookupGroupRepository2(t *testing.T) {
	logs := logger.InitLogs(&http.Request{})
	res, err:=InitLookupGroupRepository(logs).FindAll()
	log.Println(err)
	log.Println(res)
}

func TestInitLookupGroupRepository3(t *testing.T) {
	logs := logger.InitLogs(&http.Request{})
	req := dbmodels.LookupGroup{
		ID: 89,
		//Name: "lookupgroupnamebaru",
	}
	InitLookupGroupRepository(logs).Delete(req)

}

func TestInitLookupGroupRepository4(t *testing.T) {
	logs := logger.InitLogs(&http.Request{})

	req := dto.ReqLookupGroupDto{
		//Name:"ACCOUNT_TYPE",
		Page:1,
		Limit:10,
	}
	data, total,_ :=InitLookupGroupRepository(logs).Filter(req)
	log.Println(total)
	log.Println(data)
}
