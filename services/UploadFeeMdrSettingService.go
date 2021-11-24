package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

type UploadFeeMdrSettingService struct {
	UploadFeeMdrSettingRepository *db.UploadFeeMdrSettingRepository
}

func InitUploadFeeMdrSettingService() *UploadFeeMdrSettingService {
	return &UploadFeeMdrSettingService{
		UploadFeeMdrSettingRepository: db.InitUploadFeeMdrSettingRepository(),
	}
}

var (
	pathUploadFeeMdrSetting string
)

func init()  {
	pathUploadFeeMdrSetting = ottoutils.GetEnv("ROSE_UPLOAD_FEE_MDR_SETTING_PATH", "/opt/app-rose/fee-mdr-setting/upload/")
}

func (service *UploadFeeMdrSettingService) UploadFile(ctx *gin.Context, file *multipart.FileHeader) models.Response {
	var res models.Response

	user := auth.UserLogin
	transactionDate := time.Now()
	fileNameTemp := utils.ConvertTime(transactionDate) + ".xlsx"
	fileNameTemp = strings.Replace(fileNameTemp, ":", "_", -1)


	log.Println("fileNameTemp -->", fileNameTemp)
	filename := filepath.Base(fileNameTemp)

	filePathErr := fileNameTemp + "-err.xlsx"

	uploadMerchant := dbmodels.UploadFeeMdrSetting{
		FilePath: fileNameTemp,
		Date:transactionDate,
		User: user.Name,
		Status: constants.ON_PROGRESS,
		FilePathErr: filePathErr,
		Notes: "",
	}

	if err:=service.UploadFeeMdrSettingRepository.Save(&uploadMerchant); err != nil {
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
		log.Println("Save File Err ", zap.Error(err))

		res.ErrCode = constants.EC_FAIL_UPLOAD_XSLX_FILE
		res.ErrDesc = constants.EC_FAIL_UPLOAD_XSLX_FILE_DESC
		return res
	}

	/*Check Folder path, if err then create folder*/
	if _, err := os.Stat(pathUploadFeeMdrSetting); err != nil {
		fmt.Println("create new folder")
		os.MkdirAll(pathUploadFeeMdrSetting, os.ModePerm)
	}

	newPath := pathUploadFeeMdrSetting + fileNameTemp
	log.Println("newPath -> ", newPath)

	fmt.Println("--Rename File--")
	fmt.Println("rename to  ", newPath)
	if err:=os.Rename(fileNameTemp, newPath); err!= nil {
		log.Println("Save File Err ",err)
		res.ErrCode = constants.EC_FAIL_RENAME_XSLX_FILE
		res.ErrDesc = constants.EC_FAIL_RENAME_XSLX_FILE_DESC
		return res
	}

	fmt.Println("finish")

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG



	return res
}

func (service *UploadFeeMdrSettingService) GetDataUpload(req dto.ReqUploadFeeMdrSettingDto) models.Response {

	fmt.Println("<< UploadFeeMdrSettingService: GetDataUpload >>")

	var res models.Response

	var total int

	list, total, err := service.UploadFeeMdrSettingRepository.GetDataUpload(req)
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