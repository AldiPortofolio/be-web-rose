package services

import (
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
	"rose-be-go/models/dbmodels"
	"rose-be-go/kafka"
	"rose-be-go/utils"
	"strings"

	ottoutils "ottodigital.id/library/utils"
	"rose-be-go/models"
	"time"
)

type UploadNmidService struct {
	General models.GeneralModel
	Send    func(req kafka.PublishReq) ([]byte, error)
}

func InitUploadNmidService(gen models.GeneralModel) *UploadNmidService {
	return &UploadNmidService{
		General: gen,
		Send:    kafka.SendPublishKafka,
	}
}

func (service *UploadNmidService) UploadFile(ctx *gin.Context, file *multipart.FileHeader) error {

	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("UploadNmidService: UploadFile",
		zap.String("file", file.Filename))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "UploadNmidService: UploadFile")
	defer span.Finish()

	//res := models.Response{
	//	ErrCode:  "01",
	//	ErrDesc: "Transaction failed",
	//}

	time := time.Now()
	name := utils.ConvertTime(time)
	fileNameTemp := name + ".csv"

	fileNameTemp = strings.Replace(fileNameTemp, ":", "_", -1)

	filename := filepath.Base(fileNameTemp)
	log.Println("fileNameTemp -->", fileNameTemp)

	if err := ctx.SaveUploadedFile(file, filename); err != nil {
		//res = models.Response{
		//	ErrCode:  "02",
		//	ErrDesc: "Save file err",
		//}
		go sugarLogger.Error("Save File Err ", zap.Error(err))
		return err
	}

	//path := "/apps/merchant/nmid/"
	//path := "/opt/app-rose/nmid/upload/"
	path:= ottoutils.GetEnv("PATH_UPLOAD_NMID","/opt/app-rose/nmid/upload/")
	if _, err := os.Stat(path); err != nil {
		fmt.Println("create new folder")
		os.MkdirAll(path, os.ModePerm)
	}

	newPath := path + fileNameTemp

	errRename := os.Rename(fileNameTemp, newPath)
	if errRename != nil {
		//res = models.Response{
		//	ErrCode:  "02",
		//	ErrDesc: "Save file err",
		//}
		go sugarLogger.Error("Save File Err ", zap.Error(errRename))
		return errRename
	}

	log.Println("newPath --> ", newPath)

	//errKafka := service.SendToKafka(fileNameTemp)

	go service.InsertoDb(fileNameTemp)


	return nil
}

func (service *UploadNmidService) InsertoDb(filePath string)  {

	log.Println("path InsertoDb ===> ", filePath)
	filePathErr := filePath + "-err.csv"
	user := auth.UserLogin
	transactionDate := utils.ConvertTime(time.Now())
	status := constants.ON_PROGRESS

	//pack msg reportQR
	reportQR := dbmodels.ReportQr{
		User: user.Name,
		FilePath: filePath,
		FilePathErr: filePathErr,
		FilePathSuccess: filePath,
		Status: status,
		TransactionDate: transactionDate,
	}
	log.Println(reportQR)

	db.InitReportQrDataRepository().SaveReportQR(&reportQR)

	uploadNmid := dbmodels.UploadNmidData{
		FilePath: filePath,
		FilePathErr: filePathErr,
		Date: transactionDate,
		Status: status,
		User: user.Name,
	}

	db.InitUploadNmidDataRepository().SaveNmid(&uploadNmid)



}



/*

func (service *UploadNmidService) SendToKafka(path string) error {

	topic := ottoutils.GetEnv("ROSE_BE_GO_QRIS_UPLOAD_NMID_TOPIC", "qris-upload-nmid-topic")

	user := auth.UserLogin

	req := models.ReqKafkaUploadNmid{
		FilePath:  path,
		DateTime:  time.Now(),
		Requestor: user.Name,
	}
	//err := service.Send(topic, req)

	////////=====///////
	nmidReq, _ := json.Marshal(req)

	kafkaReq := kafka.PublishReq{
		Topic:     topic,
		Bytes:     nmidReq,
		Timestamp: time.Now().Format("2006-01-02"),
	}

	kafkaRes, err := service.Send(kafkaReq)

	log.Println("kafkaRes--> ", string(kafkaRes))

	return err
}

*/


