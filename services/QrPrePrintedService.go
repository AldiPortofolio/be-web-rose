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
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	ottoutils "ottodigital.id/library/utils"
	"rose-be-go/utils"
	"time"
)

type QrPrePrintedService struct {
	General 	models.GeneralModel
	SendKafka         	func(req kafka.PublishReq) ([]byte, error)
	QrPrePrintedRepository *db.QrPrePrintedRepository
}

func InitQrPrePrintedService(gen models.GeneralModel) *QrPrePrintedService {
	return &QrPrePrintedService{
		General: gen,
		SendKafka: kafka.SendPublishKafka,
		QrPrePrintedRepository: db.InitQrPrePrintedRepository(),

	}
}

func (service *QrPrePrintedService) Send (req dto.ReqQrPrePrintedSendDto) models.Response {
	fmt.Println(">>> QrPrePrintedService - Send <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("QrPrePrintedService: Send",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "QrPrePrintedService: Send")
	defer span.Finish()

	var res models.Response

	req.Key = utils.GenerateRandom(10)

	reqDb := dbmodels.QrPrePrinted{
		Mcc: req.Mcc,
		City: req.City,
		PostalCode: req.PostalCode,
		TotalReq: req.TotalReq,
		Key: req.Key,
		User: req.User,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status: constants.ON_PROGRESS,

	}

	if err:=service.QrPrePrintedRepository.Save(&reqDb); err!=nil{

		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}


	res = service.PushToKafka(req)




	return res
}


func (service *QrPrePrintedService) PushToKafka(req dto.ReqQrPrePrintedSendDto) models.Response  {
	topic := ottoutils.GetEnv("ROSE_BE_GO_QR_PREPRINTED_TOPIC", "qr-preprinted-topic")

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

func (service *QrPrePrintedService) GetDataPaging(req dto.ReqQrPrePrintedDto) models.Response {

	fmt.Println("<< QrPrePrintedService: GetDataPaging >>")

	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("QrPrePrintedService: GetDataPaging",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "QrPrePrintedService: GetDataPaging")
	defer span.Finish()

	var res models.Response

	var total int

	list, total, err := service.QrPrePrintedRepository.GetData(req)
	if err != nil {
		res.ErrCode = constants.EC_TRANSACTION_FAILED
		res.ErrDesc = constants.EC_TRANSACTION_FAILED_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.TotalData = total
	res.Contents = list

	return res
}