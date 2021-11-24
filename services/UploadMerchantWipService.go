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

type UploadMerchantWipService struct {
	General models.GeneralModel
	UploadMerchantWipRepository *db.UploadMerchantWipRepository
}

func InitUploadMerchantWipService(gen models.GeneralModel)  *UploadMerchantWipService {
	return &UploadMerchantWipService{
		General: gen,
		UploadMerchantWipRepository: db.InitUploadMerchantWipRepository(),
	}
}

var (
	pathUploadMerchantWip string
)

func init()  {
	pathUploadMerchantWip = ottoutils.GetEnv("ROSE_UPLOAD_MERCHANT_WIP_PATH", "/opt/app-rose/merchant-wip/upload/")

}

func (service *UploadMerchantWipService) UploadFile(ctx *gin.Context, file *multipart.FileHeader) models.Response {

	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("UploadMerchantWipService: UploadFile",
		zap.String("file", file.Filename))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "UploadMerchantWipService: UploadFile")
	defer span.Finish()

	var res models.Response


	user := auth.UserLogin
	transactionDate := time.Now()

	fileNameTemp := utils.ConvertTime(transactionDate) + ".xlsx"
	fileNameTemp = strings.Replace(fileNameTemp, ":", "_", -1)

	log.Println("fileNameTemp -->", fileNameTemp)
	filename := filepath.Base(fileNameTemp)

	filePathErr := fileNameTemp + "-err.xlsx"

	uploadMerchant := dbmodels.UploadMerchantWip{
		FilePath: fileNameTemp,
		Date:transactionDate,
		User: user.Name,
		Status: constants.ON_PROGRESS,
		FilePathErr: filePathErr,
		Notes: "",
	}

	if err:=service.UploadMerchantWipRepository.Save(&uploadMerchant); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}


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
	if _, err := os.Stat(pathUploadMerchantWip); err != nil {
		fmt.Println("create new folder")
		os.MkdirAll(pathUploadMerchantWip, os.ModePerm)
	}

	newPath := pathUploadMerchantWip + fileNameTemp

	if err:=os.Rename(fileNameTemp, newPath); err!= nil {
		go sugarLogger.Error("Save File Err ", zap.Error(err))
		res.ErrCode = constants.EC_FAIL_RENAME_XSLX_FILE
		res.ErrDesc = constants.EC_FAIL_RENAME_XSLX_FILE_DESC
		return res
	}



	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG



	return res
}

func (service *UploadMerchantWipService) GetDataUploadMerchant(req dto.ReqUploadMerchant) models.Response {

	fmt.Println("<< UploadMerchantWipService: GetDataUploadMerchant >>")

	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("UploadMerchantWipService: GetDataUploadMerchant",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "UploadMerchantWipService: GetDataUploadMerchant")
	defer span.Finish()

	var res models.Response

	var total int

	list, total, err := service.UploadMerchantWipRepository.GetDataUploadMerchant(req)
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