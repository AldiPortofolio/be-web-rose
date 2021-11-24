package db

import (
	"net/http"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/models/dbmodels"
	"testing"
	"time"
)

func TestInitMasterTagRepository(t *testing.T) {
	logs := logger.InitLogs(&http.Request{})

	req := dbmodels.MasterTag{
		Name: "sampo",
		Code: "001",
		Status: true,
		Description:"sampoooo",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		CreatedBy:"syandi",
		UpdatedBy: "syandi",
	}
	InitMasterTagRepository(logs).Save(&req)
}