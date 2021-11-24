package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"os"
	ottoutils "ottodigital.id/library/utils"
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
)

type UploadMissingDataService struct {
	UploadMissingDataRepository *db.UploadMissingDataRepository
}

func InitUploadMissingDataService() *UploadMissingDataService {
	return &UploadMissingDataService{
		UploadMissingDataRepository: db.InitUploadMissingDataRepository(),
	}
}

var (
	pathUploadMissingData string
)

func init()  {
	pathUploadMissingData = ottoutils.GetEnv("ROSE_UPLOAD_DATA_MISSING_PATH", "/opt/app-rose/merchant-missing/upload/")
}

func (service *UploadMissingDataService) UploadFile(ctx *gin.Context, file *multipart.FileHeader, res *models.Response)  {
	fmt.Println("<< UploadMissingDataService - Upload >>")
	user := auth.UserLogin
	transactionDate := time.Now()

	fileNameTemp := utils.ConvertTime(transactionDate) + ".xlsx"
	fileNameTemp = strings.Replace(fileNameTemp, ":", "_", -1)

	log.Println("fileNameTemp -->", fileNameTemp)
	filename := filepath.Base(fileNameTemp)

	filePathErr := fileNameTemp + "-err.xlsx"

	uploadMerchant := dbmodels.UploadMissingData{
		FilePath: fileNameTemp,
		Date:transactionDate,
		User: user.Name,
		Status: constants.ON_PROGRESS,
		FilePathErr: filePathErr,
		Notes: "",
	}

	if err:=service.UploadMissingDataRepository.Save(&uploadMerchant); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return
	}

	log.Println("save upload File")
	fmt.Println("--SaveUploadedFile -- copy save file--")
	if err := ctx.SaveUploadedFile(file, filename); err != nil {

		log.Println("Save File Err ", err.Error())

		res.ErrCode = constants.EC_FAIL_UPLOAD_XSLX_FILE
		res.ErrDesc = constants.EC_FAIL_UPLOAD_XSLX_FILE_DESC
		return
	}

	/*Check Folder path, if err then create folder*/
	if _, err := os.Stat(pathUploadMissingData); err != nil {
		fmt.Println("create new folder")
		os.MkdirAll(pathUploadMissingData, os.ModePerm)
	}

	newPath := pathUploadMissingData + fileNameTemp
	log.Println("newPath -> ", newPath)

	fmt.Println("--Rename File--")
	fmt.Println("rename to  ", newPath)
	if err:=os.Rename(fileNameTemp, newPath); err!= nil {
		log.Println("Save File Err ", err.Error())
		res.ErrCode = constants.EC_FAIL_RENAME_XSLX_FILE
		res.ErrDesc = constants.EC_FAIL_RENAME_XSLX_FILE_DESC
		return
	}

	fmt.Println("finish")

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	
}

func (service *UploadMissingDataService) GetDataUploadMerchant(req dto.ReqUploadMissingDataDto, res *models.Response) {

	fmt.Println("<< UploadMissingDataService: GetDataUploadMerchant >>")


	var total int

	list, total, err := service.UploadMissingDataRepository.GetDataUploadMerchant(req)
	if err != nil {
		res.ErrCode = constants.EC_TRANSACTION_FAILED
		res.ErrDesc = constants.EC_TRANSACTION_FAILED_DESC
		return
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.TotalData = total
	res.Contents = list

	
}