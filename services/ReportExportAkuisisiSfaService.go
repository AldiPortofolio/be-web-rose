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
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"time"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	ottoutils "ottodigital.id/library/utils"
)



type ReportExportAkuisisiSfaService struct {
	General      models.GeneralModel
	SendKafka         	func(req kafka.PublishReq) ([]byte, error)

}

func InitReportExportAkuisisiSfaService(gen models.GeneralModel) *ReportExportAkuisisiSfaService {
	return &ReportExportAkuisisiSfaService{
		General: gen,
		SendKafka: kafka.SendPublishKafka,

	}
}

func GetDataReportExportAkuisisiSfa(req dbmodels.ReportExportAkuisisiSfa) models.Response {
	var res models.Response

	list, total, err := db.GetReportExportAkuisisiSfa(req)
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

func (service *ReportExportAkuisisiSfaService) Send (req dto.ReqReportExportAkuisisiSfaDto) models.Response {
	fmt.Println(">>> ReportExportAkuisisiSfaService - Send <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("UploadNmidService: UploadFile",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "ReportExportAkuisisiSfaService: Send")
	defer span.Finish()

	var res models.Response
	
	res = service.PushToKafka(req)

	return res
}

func (service *ReportExportAkuisisiSfaService) PushToKafka(req dto.ReqReportExportAkuisisiSfaDto) models.Response  {
	topic := ottoutils.GetEnv("ROSE_BE_GO_REPORT_EXPORT_AKUISISI_SFA_TOPIC", "rose-report-export-akuisisi-sfa-topic")

	var res models.Response
	req.User = auth.UserLogin.Name
	log.Println("Username --->" , req.User)
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
