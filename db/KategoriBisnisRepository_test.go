package db

import (
	"encoding/json"
	"log"
	
	"testing"
)

func TestKategoriBisnisRepository_GetAll(t *testing.T) {
	res, err := InitKategoriBisnisRepository().GetAll()

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println(err)
}