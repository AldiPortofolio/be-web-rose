package services

import (
	"encoding/json"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"log"
	ottoutils "ottodigital.id/library/utils"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/kafka"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"rose-be-go/utils"
	"time"
)


type ReportRejectedService struct {
	General      models.GeneralModel
	SendKafka         	func(req kafka.PublishReq) ([]byte, error)

}

func InitReportRejectedService(gen models.GeneralModel) *ReportRejectedService {
	return &ReportRejectedService{
		General: gen,
		SendKafka: kafka.SendPublishKafka,

	}
}

func (service *ReportRejectedService) Send (req dto.ReqReportSendDto) models.Response {
	fmt.Println(">>> ReportFinishService - Send <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("UploadNmidService: UploadFile",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "ReportFinishService: Send")
	defer span.Finish()

	var res models.Response = service.PushToKafka(req)

	return res
}

func (service *ReportRejectedService) PushToKafka(req dto.ReqReportSendDto) models.Response  {
	topic := ottoutils.GetEnv("ROSE_BE_GO_REPORT_REJECTED_TOPIC", "rose-report-rejected-topic")

	var res models.Response
	req.User = auth.UserLogin.Name
	reqByte,_ := json.Marshal(req)

	start, err := utils.ShortDateFromString(req.StartDate)
	if err != nil {
		fmt.Println("error -->", err)
	}
	end, err := utils.ShortDateFromString(req.EndDate)
	if err != nil {
		fmt.Println("error -->", err)
	}
	fmt.Println("Start, End ->", start, end)

	if start.After(end){
		res.ErrCode = constants.EC_FAIL_DATE
		res.ErrDesc = constants.EC_FAIL_DATE_MSG
		return res
	}

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

func GetDataReportRejected(req dbmodels.ReportReject) models.Response {
	var res models.Response
	var total int

	list, total, err := db.GetDataReportRejected(req)
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

func GetReasonMerchantWip(wipid int) models.Response {
	var res models.Response
	var total int

	list, err := db.GetReasonWipQueue(wipid)
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

