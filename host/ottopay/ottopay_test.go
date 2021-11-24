package ottopay

import (
	"log"
	"net/http"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/constants"
	"rose-be-go/models/ottopaymodels"
	"testing"
)

func TestOttopayHost_SendNotif(t *testing.T) {
	logs := logger.InitLogs(&http.Request{})

	req := ottopaymodels.ReqPushNotif{
		Desc:"sdfsdfsdf",
		CustAccount:"087874566253",
		Target: "history_oc",
		Title: "fsdfdsf",
	}
	res, err := InitOttopayHost(logs).Send(req, constants.SEND_NOTIF)

	log.Println(err)
	log.Println(res)
}