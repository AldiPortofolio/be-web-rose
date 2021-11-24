package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/db"
	ottoutils "ottodigital.id/library/utils"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"rose-be-go/utils"
	"strings"
	"time"
)
type UploadMerchantBankLoanService struct {
	UploadMerchantBankLoanRepository *db.UploadMerchantBankLoanRepository

}

func InitUploadMerchantBankLoanService() *UploadMerchantBankLoanService {
	return &UploadMerchantBankLoanService{
		UploadMerchantBankLoanRepository: db.InitUploadMerchantBankLoanRepository(),
	}
}

var (
	pathUploadMerchantBankLoan string
)


func init()  {
	pathUploadMerchantBankLoan = ottoutils.GetEnv("ROSE_UPLOAD_MERCHANT_BANK_LOAN_PATH", "/opt/app-rose/merchant-bank-loan/upload/")

}

func (service *UploadMerchantBankLoanService) UploadFile(ctx *gin.Context, file *multipart.FileHeader, res *models.Response)  {
	fmt.Println("<< UploadMerchantBankLoanService - Upload >>")
	user := auth.UserLogin
	transactionDate := time.Now()

	fileNameTemp := utils.ConvertTime(transactionDate) + ".xlsx"
	fileNameTemp = strings.Replace(fileNameTemp, ":", "_", -1)

	log.Println("fileNameTemp -->", fileNameTemp)
	filename := filepath.Base(fileNameTemp)

	filePathErr := fileNameTemp + "-err.xlsx"

	uploadMerchant := dbmodels.UploadMerchantBankLoan{
		FilePath: fileNameTemp,
		Date:transactionDate,
		User: user.Name,
		Status: constants.ON_PROGRESS,
		FilePathErr: filePathErr,
		Notes: "",
	}

	if err:=service.UploadMerchantBankLoanRepository.Save(&uploadMerchant); err != nil {
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
	if _, err := os.Stat(pathUploadMerchantBankLoan); err != nil {
		fmt.Println("create new folder")
		os.MkdirAll(pathUploadMerchantBankLoan, os.ModePerm)
	}

	newPath := pathUploadMerchantBankLoan + fileNameTemp
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


func (service *UploadMerchantBankLoanService) GetDataUploadMerchant(req dto.ReqUploadMerchantBankLoanDto, res *models.Response) {

	fmt.Println("<< UploadMerchantBankLoanService: GetDataUploadMerchant >>")


	var total int

	list, total, err := service.UploadMerchantBankLoanRepository.GetDataUploadMerchant(req)
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