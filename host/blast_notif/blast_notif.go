package blast_notif

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/kelseyhightower/envconfig"
	"net/http"
	"rose-be-go/constants"
	httputils "rose-be-go/utils/http"

)

type BlastNotifHost struct {

}

func InitBlastNotifHost() *BlastNotifHost {
	return &BlastNotifHost{}
}

type BlastNotifEnv struct {
	Name              string `envconfig:"NAME" default:"BLAST NOTIF"`
	Host              string `envconfig:"HOST" default:"http://13.228.25.85:8987"`
	SendAll    string `envconfig:"SEND_ALL" default:"/ottopay/v0.1.0/sendnotif/all"`
}

var (
	blastNotifEnv BlastNotifEnv
)

func init() {
	err := envconfig.Process("BLAST_NOTIF", &blastNotifEnv)
	if err != nil {
		fmt.Println("Failed to get BLAST_NOTIF env:", err)
	}
}


// Send ..
func (head *BlastNotifHost) Send(msgReq interface{}, typeTrans string) ([]byte, error) {
	header := make(http.Header)
	header.Add("Accept", "*/*")

	method := constants.HttpMethodPost

	dataReq, _ := json.Marshal(msgReq)

	urlSvr := ""
	switch typeTrans {
	case constants.BLAST_NOTIF_SEND_ALL:
		urlSvr = blastNotifEnv.Host + blastNotifEnv.SendAll
		break
	}

	datareq, _ := json.Marshal(msgReq)

	fmt.Println(string(dataReq))


	logs.Info(fmt.Sprintf("[Request %s]", typeTrans), fmt.Sprintf("[%s]", string(datareq)))

	fmt.Println("Headerrrrrr======>", header)

	data, err := httputils.SendHttpRequest(method, urlSvr, header, msgReq)

	fmt.Println("xxxx-----------xxxx")
	fmt.Println("urlSvr", urlSvr)
	fmt.Println("msgreq", msgReq)
	fmt.Println("header", header)
	fmt.Println("datareq", string(datareq))
	fmt.Println("err", err)
	fmt.Println("xxxx-----------xxxx")

	logs.Info(fmt.Sprintf("[Response %s]", typeTrans), fmt.Sprintf("[%s]", string(data)))

	return data, err
}
