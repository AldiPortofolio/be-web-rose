package db

import (
	"log"
	"testing"
)

func TestInitHistoryClearSessionRepository_GetLastUpdated(t *testing.T) {
	data, err := InitHistoryClearSessionRepository().GetLastUpdated()
	log.Println(data)
	log.Println(err)
}