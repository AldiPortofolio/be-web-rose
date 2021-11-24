package op_bank

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/kelseyhightower/envconfig"
	"net/http"
	httputils "rose-be-go/utils/http"

	"rose-be-go/constants"
)

type OpBankHost struct {
}

func InitOpBankHost() *OpBankHost {
	return &OpBankHost{}
}

type OpBankEnv struct {
	Name              string `envconfig:"NAME" default:"OP BANK"`
	Host              string `envconfig:"HOST" default:"http://13.228.25.85:8989"`
	InquiryInternal    string `envconfig:"INQUIRY_INTERNAL" default:"/v1.0/mandiri/account/inquiry"`
	InquiryExternal string `envconfig:"INQUIRY_EXTERNAL" default:"/v1.0/mandiri/account/inquiry/external"`
}

var (
	opBankEnv OpBankEnv
)

func init() {
	err := envconfig.Process("OP_BANK", &opBankEnv)
	if err != nil {
		fmt.Println("Failed to get OP_BANK env:", err)
	}
}

// Send ..
func (head *OpBankHost) Send(msgReq interface{}, typeTrans string) ([]byte, error) {
	header := make(http.Header)
	header.Add("Accept", "*/*")

	method := constants.HttpMethodPost

	dataReq, _ := json.Marshal(msgReq)

	urlSvr := ""
	switch typeTrans {
	case constants.INQUIRY_INTERNAL:
		urlSvr = opBankEnv.Host + opBankEnv.InquiryInternal
		break
	case constants.INQUIRY_EXTERNAL:
		urlSvr = opBankEnv.Host + opBankEnv.InquiryExternal

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

