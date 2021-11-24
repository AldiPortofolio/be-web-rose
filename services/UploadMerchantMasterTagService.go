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

type UploadMerchantMasterTagService struct {
	UploadMerchantMasterTagRepository *db.UploadMerchantMasterTagRepository
}

func InitUploadMerchantMasterTagService() *UploadMerchantMasterTagService {
	return &UploadMerchantMasterTagService{
		UploadMerchantMasterTagRepository: db.InitUploadMerchantMasterTagRepository(),
	}
}

var (
	pathUploadMerchantMasterTag string
)

func init()  {
	pathUploadMerchantMasterTag = ottoutils.GetEnv("ROSE_UPLOAD_MERCHANT_MASTER_TAG_PATH", "/opt/app-rose/merchant-master-tag/upload/")

}

func (service *UploadMerchantMasterTagService) UploadFile(ctx *gin.Context, file *multipart.FileHeader, res *models.Response)  {
	fmt.Println("<< UploadMerchantMasterTagService - Upload >>")
	user := auth.UserLogin
	transactionDate := time.Now()

	fileNameTemp := utils.ConvertTime(transactionDate) + ".xlsx"
	fileNameTemp = strings.Replace(fileNameTemp, ":", "_", -1)

	log.Println("fileNameTemp -->", fileNameTemp)
	filename := filepath.Base(fileNameTemp)

	filePathErr := fileNameTemp + "-err.xlsx"

	uploadMerchant := dbmodels.UploadMerchantMasterTag{
		FilePath: fileNameTemp,
		Date:transactionDate,
		User: user.Name,
		Status: constants.ON_PROGRESS,
		FilePathErr: filePathErr,
		Notes: "",
	}

	if err:=service.UploadMerchantMasterTagRepository.Save(&uploadMerchant); err != nil {
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
	if _, err := os.Stat(pathUploadMerchantMasterTag); err != nil {
		fmt.Println("create new folder")
		os.MkdirAll(pathUploadMerchantMasterTag, os.ModePerm)
	}

	newPath := pathUploadMerchantMasterTag + fileNameTemp
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

func (service *UploadMerchantMasterTagService) GetDataUploadMerchant(req dto.ReqUploadMerchantMasterTagDto, res *models.Response) {

	fmt.Println("<< UploadMerchantMasterTagService: GetDataUploadMerchant >>")


	var total int

	list, total, err := service.UploadMerchantMasterTagRepository.GetDataUploadMerchant(req)
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