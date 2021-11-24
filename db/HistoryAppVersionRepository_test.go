package db

import (
	"rose-be-go/models/dbmodels"
	"testing"
	"time"
)

func TestInitHistoryAppVersionRepository_Save(t *testing.T) {
	req := dbmodels.HistoryAppVersion{
		Version:"2",
		CreatedBy:"testing",
		AppName:"tes App",
		CreatedAt: time.Now(),
	}
	InitHistoryAppVersionRepository().Save(req);
}