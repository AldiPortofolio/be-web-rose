package blast_notif

import (
	"encoding/json"
	"log"
	"rose-be-go/constants"
	"rose-be-go/models/blastnotifmodels"
	"testing"
)

func TestInitBlastNotifHost(t *testing.T) {

	req := blastnotifmodels.ReqNotifAll{
		Tilte:"asdfasd",
		Desc:"fsdsfdsf",
		Target: "sdfdsfds",
	}

	res, err:= InitBlastNotifHost().Send(req,constants.BLAST_NOTIF_SEND_ALL)

	log.Println(err)
	log.Println(string(res))
	var resData []blastnotifmodels.ResNotifAll

	json.Unmarshal(res,&resData)
	log.Println(resData)

}