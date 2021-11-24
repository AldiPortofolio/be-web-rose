package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/kafka"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"rose-be-go/utils"
	"strings"
	"time"
	ottoutils "ottodigital.id/library/utils"
)

type MerchantAggUpload struct {
	General models.GeneralModel
	SendKafka   func(req kafka.PublishReq) ([]byte, error)
}

func InitMerchantAggUpload(gen models.GeneralModel) *MerchantAggUpload {
	return &MerchantAggUpload{
		General: gen,
		SendKafka: kafka.SendPublishKafka,
	}
}

func (service *MerchantAggUpload) UplaodFile(ctx *gin.Context, file *multipart.FileHeader, mid string) models.Response {

	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MerchantAggUploadService: UploadFile",
		zap.String("file", file.Filename))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MerchantAggUploadService: UploadFile")
	defer span.Finish()
	var res models.Response

	time := time.Now()
	name := utils.ConvertTime(time)
	fileNameTemp := name + ".csv"

	fileNameTemp = strings.Replace(fileNameTemp, ":", "_", -1)

	filename := filepath.Base(fileNameTemp)
	log.Println("fileNameTemp -->", fileNameTemp)

	if err := ctx.SaveUploadedFile(file, filename); err != nil {
		go sugarLogger.Error("Save File Err ", zap.Error(err))
		res.ErrCode = "01"
		res.ErrDesc = err.Error()
		return res
	}

	path:= ottoutils.GetEnv("PATH_MERCHANT_AGG_UPLOAD","/opt/app-rose/merchant-agg/upload/")
	if _, err := os.Stat(path); err != nil {
		fmt.Println("create new folder")
		os.MkdirAll(path, os.ModePerm)
	}

	newPath := path + fileNameTemp

	errRename := os.Rename(fileNameTemp, newPath)
	if errRename != nil {
		go sugarLogger.Error("Save File Err ", zap.Error(errRename))
		res.ErrCode = "01"
		res.ErrDesc = errRename.Error()
		return res
	}

	log.Println("newPath --> ", newPath)

	go service.InserDb(fileNameTemp, mid)

	var req dto.ReqMerchantAggregatorUpload
	req.FilePath = fileNameTemp
	req.MidAggregator = mid

	res = service.PushToKafkaMerchantAggregator(req)

	return res
}


func (service *MerchantAggUpload) InserDb(filePath string, mid string)  {

	log.Println("path InsertoDb ===> ", filePath)
	filePathErr := filePath + "-err.csv"
	user := auth.UserLogin
	transactionDate := utils.ConvertTime(time.Now())
	status := constants.ON_PROGRESS

	merchantAggUpload := dbmodels.MerchantAggUpload{
		FilePath: filePath,
		FilePathSuccess: filePath,
		FilePathErr: filePathErr,
		Date: transactionDate,
		Status: status,
		User: user.Name,
		MidAggregator: mid,
	}
	db.InitMerchantAggUploadRepository().SaveMerchantAggUpload(&merchantAggUpload)

}

func GetDataMerchantAggUpload(req dbmodels.MerchantAggUpload) models.Response {
	var res models.Response
	var total int

	list, total, err := db.InitMerchantAggUploadRepository().GetDataMerchantAggUpload(req)
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


func (service *MerchantAggUpload) Approve(req dto.ReqMerchantAggregatorUpload) models.Response {
	fmt.Println(">>> ReportQRService - Send <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("ReportQRService: Send",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "ReportQRService: Send")
	defer span.Finish()

	var res models.Response = service.PushToKafkaMerchantAggregator(req)

	return res

}


func (service *MerchantAggUpload) PushToKafkaMerchantAggregator(req dto.ReqMerchantAggregatorUpload) models.Response  {
	topic := ottoutils.GetEnv("ROSE_WORKER_APPROVE_AGG_UPLOAD_KAFKA_TOPICS", "rose-worker-approve-agg-upload-topic")

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
