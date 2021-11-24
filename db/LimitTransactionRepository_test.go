package db

import (
	"log"
	"rose-be-go/models/dbmodels"
	"testing"
)

func TestInitLimitTransactionRepository(t *testing.T) {

	req := dbmodels.LimitTransaction{
		UserCategory:"op",
		LevelMerchant:"silver",
	}

	err := InitLimitTransactionRepository().Save(&req)
	log.Println(req)
	log.Println(err)
}

func TestLimitTransactionRepository_FindByID(t *testing.T) {
	data, err :=InitLimitTransactionRepository().FindByID(1)
	log.Println(data)
	log.Println(err)
}