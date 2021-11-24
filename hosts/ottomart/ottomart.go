package ottomart

import (
	"encoding/json"
	"fmt"
	"net/http"
	ottoutils "ottodigital.id/library/utils"
	"rose-be-go/constants"
	"rose-be-go/models"
	httputils "rose-be-go/utils/http"
)

type OttomartHost struct {
	Authorization string
	DeviceID      string
	models.GeneralModel
}

func InitOttomartHost()  *OttomartHost{
	return &OttomartHost{
		Authorization:"",
		DeviceID:"",
	}
}


var (
	host               string
	endpointClearSession string
)

func init() {
	host = ottoutils.GetEnv("OTTOMART_HOST", "http://13.228.25.85:8999/")
	endpointClearSession = ottoutils.GetEnv("OTTOMART_ENDPOINT_CLEAR_SESSION", "ottopay/v0.1.0/ottopay-mart/forcelogout")
}


// PackMessageHeader ..
func (head *OttomartHost) PackMessageHeader(authorization string, deviceID string) {
	head.Authorization = "Bearer " + authorization
	head.DeviceID = deviceID

	return
}

// Send ..
func (head *OttomartHost) Send(msgReq interface{}, typeTrans string) ([]byte, error) {

	header := make(http.Header)
	header.Add("Accept", "*/*")
	header.Add("Authorization", head.Authorization)
	header.Add("Device-Id", head.DeviceID)

	fmt.Println("header:", header)

	urlSvr := ""
	method := constants.HttpMethodPost
	switch typeTrans {

	case constants.OttomartClearSession:
		fmt.Println(constants.OttomartClearSession)
		urlSvr = host + endpointClearSession
		method = constants.HttpMethodGet
		header.Add("Content-Type", "application/json")
		break

	}

	dataReq, _ := json.Marshal(msgReq)

	fmt.Println("req --> ", string(dataReq))

	data, err := httputils.SendHttpRequest(method, urlSvr, header, msgReq)
	fmt.Println()

	fmt.Println("res -->", string(data))

	return data, err
}