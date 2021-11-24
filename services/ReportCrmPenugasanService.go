package services

import (
	"encoding/json"
	"fmt"
	"log"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/kafka"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	ottoutils "ottodigital.id/library/utils"

	"time"
)

type ReportCrmPenugasanService struct {
	SendKafka         	func(req kafka.PublishReq) ([]byte, error)
	ReportCrmPenugasanRepository *db.ReportCrmPenugasanRepository
}



func InitReportCrmPenugasanService() *ReportCrmPenugasanService  {
	return &ReportCrmPenugasanService{
		SendKafka: kafka.SendPublishKafka,
		ReportCrmPenugasanRepository: db.InitReportCrmPenugasanRepository(),
	}
}

func (service *ReportCrmPenugasanService) Send (req dto.ReqReportCrmPenugasanDto) models.Response {
	fmt.Println(">>> ReportCrmPenugasanService - Send <<<")


	var res models.Response = service.PushToKafka(req)


	return res
}

func (service *ReportCrmPenugasanService) PushToKafka(req dto.ReqReportCrmPenugasanDto) models.Response  {
	topic := ottoutils.GetEnv("ROSE_BE_GO_REPORT_CRM_PENUGASAN_TOPIC", "rose-report-crm-penugasan-topic")

	var res models.Response
	req.User = auth.UserLogin.Name
	reqByte,_ := json.Marshal(req)

	kafkaReq := kafka.PublishReq{
		Topic: topic,
		Bytes: reqByte,
		Timestamp: time.Now().Format("2006-01-02"),
	}

	kafkaRes, err := service.SendKafka(kafkaReq)
	if err != nil {
		res.ErrCode = constants.ERR_CODE_04
		res.ErrDesc = constants.ERR_CODE_04_MSG
		return res
	}
	log.Println("kafkaRes--> ", string(kafkaRes))


	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	return res
}


func (service *ReportCrmPenugasanService)FilterPaging(req dto.ReqGetReporCrmPenugasanDto) models.Response{
	var res models.Response

	list, total, err := service.ReportCrmPenugasanRepository.FilterPaging(req)
	if err != nil {
		res.ErrCode = "05"
		return res
	}

	res.ErrCode = "00"
	res.ErrDesc = "Success"
	res.TotalData = total
	res.Contents = list
	return res
}
