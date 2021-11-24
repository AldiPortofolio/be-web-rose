package stringbuilder

import (
	"encoding/json"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"net/http"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/constants"
	httputils"rose-be-go/utils/http"
)

type SbHost struct {
	Ottolog logger.OttologInterface
}

func InitSbHost(logs logger.OttologInterface) *SbHost {
	return &SbHost{
		Ottolog:logs,
	}
}

type SbEnv struct {
	Name           string `envconfig:"NAME" default:"String Builder"`
	Host           string `envconfig:"HOST" default:"http://13.228.25.85:8995"`
	GenerateQr	string `envconfig:"GENERATE_QR" default:"/merchant/qr/generate"`
}

var (
	sbEnv SbEnv
)

func init() {
	err := envconfig.Process("STRING_BUILDER", &sbEnv)
	if err != nil {
		fmt.Println("Failed to get OTTOPAY  env:", err)
	}
}

// Send ..
func (head *SbHost) GenerateQr(msgReq interface{}) ([]byte, error) {

	header := make(http.Header)
	header.Add("Accept", "*/*")

	urlSvr := sbEnv.Host + sbEnv.GenerateQr
	method := constants.HttpMethodPost


	dataReq, _ := json.Marshal(msgReq)

	fmt.Println(string(dataReq))

	data, err := httputils.SendHttpRequest(method, urlSvr, header, msgReq)
	fmt.Println("Response STRING BUILDER: ", string(data))


	return data, err
}