package services

import (
	"encoding/json"
	"log"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	"testing"
)

func TestMerchantAggregatorDetailService_PackMsgTemp(t *testing.T) {
	req := dto.ReqFilterDto {
		MidAggregator: "12312",
	}
	res := InitMerchantAggregatorDetailService(models.GeneralModel{}).PackMsgTemp(req, "")
	byteRes,_ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))

}


//func TestMerchantAggregatorDetailService_ApproveAggregator(t *testing.T) {
//	temp := dbmodels.MerchantAggregatorDetailTemp{
//		MpanAggregator: "0005",
//		MpanMerchant:"006",
//	}
//	detail := dbmodels.MerchantAggregatorDetail {
//		MpanAggregator: "10005",
//		MpanMerchant:"1006",
//	}
//	res := InitMerchantAggregatorDetailService(models.GeneralModel{}).ApproveAggregator(temp, &detail)
//	byteRes,_ := json.Marshal(res)
//	log.Println("res --> ", string(byteRes))
//
//}
