package ottopay

import (
	"encoding/json"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"net/http"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/constants"
	httputils"rose-be-go/utils/http"
)

type OttopayHost struct {
	Ottolog logger.OttologInterface
}

func InitOttopayHost(logs logger.OttologInterface) *OttopayHost {
	return &OttopayHost{
		Ottolog:logs,
	}
}

type OttopayEnv struct {
	Name           string `envconfig:"NAME" default:"Ottopay"`
	Host           string `envconfig:"HOST" default:"http://13.228.25.85:8987"`
	EndpointPushNotif	string `envconfig:"ENDPOINT_PUSH_NOTIF" default:"/ottopay/v0.1.0/sendnotif"`
}

var (
	ottopayEnv OttopayEnv
)

func init() {
	err := envconfig.Process("OTTOPAY", &ottopayEnv)
	if err != nil {
		fmt.Println("Failed to get OTTOPAY  env:", err)
	}
}

// Send ..
func (head *OttopayHost) Send(msgReq interface{}, typeTrans string) ([]byte, error) {

	header := make(http.Header)
	header.Add("Accept", "*/*")

	urlSvr := ottopayEnv.Host + ottopayEnv.EndpointPushNotif
	method := constants.HttpMethodPost


	dataReq, _ := json.Marshal(msgReq)

	fmt.Println(string(dataReq))

	data, err := httputils.SendHttpRequest(method, urlSvr, header, msgReq)
	fmt.Println("Response Ottopay push Notif: ", string(data))


	return data, err
}