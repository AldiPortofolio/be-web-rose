package db

import (
	"log"
	"testing"
)

func TestGetWipQueueByUserAndStatus(t *testing.T) {
	res, err := InitWipQueueRepository().GetWipQueueByUserAndStatus("syandi", "REGISTERED")
	log.Println(res)
	log.Println(err)
}
