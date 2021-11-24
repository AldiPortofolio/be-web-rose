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
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"rose-be-go/utils"
	"strings"
	"time"
	ottoutils "ottodigital.id/library/utils"

)

type UploadMerchantOkService struct {
	General models.GeneralModel
	UploadMerchantOkRepository *db.UploadMerchantOkRepository
}

func InitUploadMerchantOkService(gen models.GeneralModel)  *UploadMerchantOkService {
	return &UploadMerchantOkService{
		General: gen,
		UploadMerchantOkRepository: db.InitUploadMerchantOkRepository(),
	}
}

var (
	pathUploadMerchantOk string
)

func init()  {
	pathUploadMerchantOk = ottoutils.GetEnv("ROSE_UPLOAD_MERCHANT_OK_PATH", "/opt/app-rose/merchant-ok/upload/")

}

func (service *UploadMerchantOkService) UploadFile(ctx *gin.Context, file *multipart.FileHeader) models.Response {

	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("UploadMerchantOkService: UploadFile",
		zap.String("file", file.Filename))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "UploadMerchantOkService: UploadFile")
	defer span.Finish()

	var res models.Response


	user := auth.UserLogin
	transactionDate := time.Now()

	fileNameTemp := utils.ConvertTime(transactionDate) + ".xlsx"
	fileNameTemp = strings.Replace(fileNameTemp, ":", "_", -1)

	log.Println("fileNameTemp -->", fileNameTemp)
	filename := filepath.Base(fileNameTemp)

	filePathErr := fileNameTemp + "-err.xlsx"

	uploadMerchant := dbmodels.UploadMerchantOk{
		FilePath: fileNameTemp,
		Date:transactionDate,
		User: user.Name,
		Status: constants.ON_PROGRESS,
		FilePathErr: filePathErr,
		Notes: "",
	}

	if err:=service.UploadMerchantOkRepository.Save(&uploadMerchant); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}


	log.Println("save upload File")

	fmt.Println("--SaveUploadedFile -- copy save file--")
	if err := ctx.SaveUploadedFile(file, filename); err != nil {
		//res = models.Response{
		//	ErrCode:  "02",
		//	ErrDesc: "Save file err",
		//}
		go sugarLogger.Error("Save File Err ", zap.Error(err))

		res.ErrCode = constants.EC_FAIL_UPLOAD_XSLX_FILE
		res.ErrDesc = constants.EC_FAIL_UPLOAD_XSLX_FILE_DESC
		return res
	}

	/*Check Folder path, if err then create folder*/
	if _, err := os.Stat(pathUploadMerchantOk); err != nil {
		fmt.Println("create new folder")
		os.MkdirAll(pathUploadMerchantOk, os.ModePerm)
	}

	newPath := pathUploadMerchantOk + fileNameTemp
	log.Println("newPath -> ", newPath)

	fmt.Println("--Rename File--")
	fmt.Println("rename to  ", newPath)
	if err:=os.Rename(fileNameTemp, newPath); err!= nil {
		go sugarLogger.Error("Save File Err ", zap.Error(err))
		res.ErrCode = constants.EC_FAIL_RENAME_XSLX_FILE
		res.ErrDesc = constants.EC_FAIL_RENAME_XSLX_FILE_DESC
		return res
	}

	fmt.Println("finish")

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG



	return res
}

func (service *UploadMerchantOkService) GetDataUploadMerchant(req dto.ReqUploadMerchant) models.Response {

	fmt.Println("<< UploadMerchantOkService: GetDataUploadMerchant >>")

	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("UploadMerchantOkService: GetDataUploadMerchant",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "UploadMerchantOkService: GetDataUploadMerchant")
	defer span.Finish()

	var res models.Response

	var total int

	list, total, err := service.UploadMerchantOkRepository.GetDataUploadMerchant(req)
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