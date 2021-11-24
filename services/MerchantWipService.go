package services

import (
	"encoding/json"
	"fmt"
	"log"
	"rose-be-go/auth"
	"rose-be-go/constants/status_registration"
	"rose-be-go/db"
	"rose-be-go/models"
)

type MerchantWipService struct {
	General models.GeneralModel
}

func InitMerchantWipService(gen models.GeneralModel) *MerchantWipService {
	return &MerchantWipService{
		General: gen,
	}
}

func (service *MerchantWipService) GetNextRegistered() models.Response {
	var res models.Response

	var merchantWip models.MerchantWip
	queueLevel := status_registration.NONE;


	merchantWip = service.GetMerchantWipByStatusMessaging(status_registration.VVIP_VERIFIER_START_VIEW)
	if merchantWip.ID == 0 {
		queueLevel = status_registration.VVIP_VERIFIER_START_VIEW
		merchantWip = service.GetMerchantWipByStatusMessaging(status_registration.VIP_VERIFIER_START_VIEW)
	}

	if merchantWip.ID == 0 {
		queueLevel = status_registration.VIP_VERIFIER_START_VIEW
		merchantWip = service.GetMerchantWipByStatusMessaging(status_registration.VERIFIER_START_VIEW)
	}

	if merchantWip.ID == 0 {
		queueLevel = status_registration.VERIFIER_START_VIEW
		merchantWip = service.GetMerchantWipByStatusMessaging(status_registration.VERIFIER_START_VIEW)
	}

	if queueLevel != status_registration.NONE {
		fmt.Println(queueLevel)
	}



	log.Println("merchantWip --> ",merchantWip)

	return res
}


func (service *MerchantWipService) GetMerchantWipByStatusMessaging(status string) models.MerchantWip {
	var res models.MerchantWip

	user := auth.UserLogin.Name
	log.Println("status ===> ", status)
	log.Println("user --> ", user)

	wipQueue, _ := db.InitWipQueueRepository().GetWipQueueByUserAndStatus(user, status)
	json.Unmarshal([]byte(wipQueue.Value), &res)
	//res = wipQueue.Value


	return res
}