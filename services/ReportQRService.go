package services

import (
	"encoding/json"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
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

type ReportQRService struct {
	General				models.GeneralModel
	SendKafka         	func(req kafka.PublishReq) ([]byte, error)
}

func InitReportQRService(gen models.GeneralModel) *ReportQRService  {
	return &ReportQRService{
		General: gen,
		SendKafka: kafka.SendPublishKafka,
	}
}

func (service *ReportQRService) Send(req dto.ReqReportSendDto) models.Response {
	fmt.Println(">>> ReportQRService - Send <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("ReportQRService: Send",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "ReportQRService: Send")
	defer span.Finish()

	var res models.Response = service.PushToKafka(req)


	return res

}

func (service *ReportQRService) PushToKafka(req dto.ReqReportSendDto) models.Response  {
	topic := ottoutils.GetEnv("ROSE_BE_GO_REPORT_QR_TOPIC", "rose-report-qr-topic")

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

// GetDataReportQr ...
func GetDataReportQr(req models.Pagination) models.Response {
	var res models.Response
	var total int

	list, total, err := db.InitReportQrDataRepository().GetDataReportQr(req)
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