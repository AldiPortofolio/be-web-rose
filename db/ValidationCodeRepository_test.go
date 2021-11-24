package db

import (
	"log"
	"net/http"
	"ottodigital.id/library/logger/v2"
	"testing"
)

func TestInitValidationRepository(t *testing.T) {
	logs := logger.InitLogs(&http.Request{})

	count, err:=InitValidationRepository(logs).CountValidationCode("DAFTARSEKARANGJUGA")
	log.Println(count)
	log.Println(err)
}