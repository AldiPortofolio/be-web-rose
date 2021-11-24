package op_bank

import (
	"encoding/json"
	"log"
	"rose-be-go/constants"
	"rose-be-go/models/opbankmodels"
	"testing"
)

func TestOpBankHost_Send(t *testing.T) {

	req := opbankmodels.ReqInquiry{
		AccountNo:"1190080001041",
		BankCode: "014",
	}
	res, err :=InitOpBankHost().Send(req, constants.INQUIRY_INTERNAL)

	var resData opbankmodels.ResInquiryInternal
	json.Unmarshal(res, &resData)
	log.Println("data --->", resData)
	log.Println(err)
	log.Println(res)
}