package stringbuilder

import (
	"encoding/json"
	"log"
	"net/http"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/models/sbmodels"
	"testing"
)

func TestSbHost_GenerateQr(t *testing.T) {
	logs:= logger.InitLogs(&http.Request{})

	req := sbmodels.ReqGenerateQr{
		Mid:"OP1A00005350",
		Tid:"A01",
	}
	res, err :=InitSbHost(logs).GenerateQr(req)
	log.Println(string(res))

	var data sbmodels.ResGenerateQr
	json.Unmarshal(res, &data)
	log.Println(data)
	log.Println(err)
}