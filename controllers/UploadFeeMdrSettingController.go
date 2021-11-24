package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
	"os"
	"rose-be-go/constants"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	"rose-be-go/services"
	ottoutils "ottodigital.id/library/utils"

)

type UploadFeeMdrSettingController struct {

}

func (controller *UploadFeeMdrSettingController) Upload(ctx *gin.Context) {
	fmt.Println("UploadFeeMdrSettingController")

	res := models.Response{
		ErrCode:  constants.EC_TRANSACTION_FAILED,
		ErrDesc: constants.EC_TRANSACTION_FAILED_DESC,
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println("Request body error:", err)
		ctx.JSON(http.StatusBadRequest, res)
		log.Println("Body request error ", zap.Error(err))
		return
	}

	res = services.InitUploadFeeMdrSettingService().UploadFile(ctx, file)

	ctx.JSON(http.StatusOK, res)

}

func (controller *UploadFeeMdrSettingController) GetFilterPaging(ctx *gin.Context)  {
	fmt.Println("GetFilterPaging")

	var req dto.ReqUploadFeeMdrSettingDto
	res := models.Response{
		ErrCode:  constants.EC_TRANSACTION_FAILED,
		ErrDesc: constants.EC_TRANSACTION_FAILED_DESC,
	}


	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}



	res = services.InitUploadFeeMdrSettingService().GetDataUpload(req)



	ctx.JSON(http.StatusOK, res)

}

func (controller *UploadFeeMdrSettingController) HandleResultDownload(ctx *gin.Context) {

	req := models.ReqDowloadFile{}

	res := models.Response{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}


	pathDir := ottoutils.GetEnv("PATH_DOWNLOAD_RESULT_FEE_MDR","/opt/app-rose/fee-mdr-setting/result/")
	path := pathDir + req.FilePath

	fmt.Println("File Path :", path)

	w := ctx.Writer
	f, err := os.Open(path)

	if f != nil {
		defer f.Close()
	}
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Error(w,"File Not Found", http.StatusInternalServerError)
		fmt.Println("Failed to open file", err)
		return
	}

	contentDisposition := fmt.Sprintf("attachment; filename=%s", f.Name())
	w.Header().Set("Content-Disposition", contentDisposition)

	if _, err := io.Copy(w, f); err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Error(w,"File Not Found", http.StatusInternalServerError)
		fmt.Println("Failed to copy file", err)
		return
	}
}

func (controller *UploadFeeMdrSettingController) HandleDownload(ctx *gin.Context) {

	req := models.ReqDowloadFile{}

	res := models.Response{}


	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}


	pathDir := ottoutils.GetEnv("PATH_DOWNLOAD_FEE_MDR","/opt/app-rose/fee-mdr-setting/process/")
	path := pathDir + req.FilePath

	fmt.Println("File Path :", path)

	w := ctx.Writer
	f, err := os.Open(path)

	if f != nil {
		defer f.Close()
	}
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Error(w,"File Not Found", http.StatusInternalServerError)
		fmt.Println("Failed to open file", err)
		return
	}

	contentDisposition := fmt.Sprintf("attachment; filename=%s", f.Name())
	w.Header().Set("Content-Disposition", contentDisposition)

	if _, err := io.Copy(w, f); err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Error(w,"File Not Found", http.StatusInternalServerError)
		fmt.Println("Failed to copy file", err)
		return
	}
}


func (controller *UploadFeeMdrSettingController) HandleTemplateDownload(ctx *gin.Context) {


	pathDir := ottoutils.GetEnv("PATH_DOWNLOAD_TEMPLATE_FEE_MDR","/opt/app-rose/fee-mdr-setting/")
	path := pathDir + "template.xlsx"

	fmt.Println("File Path :", path)

	w := ctx.Writer
	f, err := os.Open(path)

	if f != nil {
		defer f.Close()
	}
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Error(w,"File Not Found", http.StatusInternalServerError)
		fmt.Println("Failed to open file", err)
		return
	}

	contentDisposition := fmt.Sprintf("attachment; filename=%s", f.Name())
	w.Header().Set("Content-Disposition", contentDisposition)

	if _, err := io.Copy(w, f); err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Error(w,"File Not Found", http.StatusInternalServerError)
		fmt.Println("Failed to copy file", err)
		return
	}
}