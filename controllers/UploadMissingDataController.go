package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	"rose-be-go/services"
	ottoutils "ottodigital.id/library/utils"

)

type UploadMissingDataController struct {

}

func (controller *UploadMissingDataController) Upload(ctx *gin.Context)  {
	fmt.Println("UploadMissingDataController - Upload")
	var res models.Response
	logs := logger.InitLogs(ctx.Request)


	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println("Request body error:", err)
		ctx.JSON(http.StatusBadRequest, res)
		log.Println("Body request error ", err.Error())
		return
	}


	services.InitUploadMissingDataService().UploadFile(ctx, file, &res)

	bodyRes, _ := json.Marshal(res)
	logs.Info("Response UploadMissingDataController - Upload",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}

func (controller *UploadMissingDataController) GetFilterPaging(ctx *gin.Context)  {
	fmt.Println("UploadMerchantBankLoanController - GetFilterPaging")
	logs := logger.InitLogs(ctx.Request)
	var req dto.ReqUploadMissingDataDto
	var res models.Response


	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}

	reqByte,_ := json.Marshal(req)
	log.Println("req --> ", string(reqByte))

	services.InitUploadMissingDataService().GetDataUploadMerchant(req,&res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response UploadMerchantBankLoanController - Filter",
		logs.AddField("ResponseBody: ", string(bodyRes)))



	ctx.JSON(http.StatusOK, res)

}

func (controller *UploadMissingDataController) HandleDownload(ctx *gin.Context) {

	req := models.ReqDowloadFile{}

	res := models.Response{}


	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}


	pathDir := ottoutils.GetEnv("ROSE_DOWNLOAD_MISSING_DATA_PATH","/opt/app-rose/merchant-missing/process/")
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


func (controller *UploadMissingDataController) HandleResultDownload(ctx *gin.Context) {

	req := models.ReqDowloadFile{}

	res := models.Response{}


	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}

	pathDir := ottoutils.GetEnv("PATH_DOWNLOAD_RESULT_DATA_MISSING","/opt/app-rose/merchant-missing/result/")
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

func (controller *UploadMissingDataController) HandleDownloadExample(ctx *gin.Context) {

	pathDir := ottoutils.GetEnv("PATH_DOWNLOAD_EXAMPLE_MERCHANT_MISSING","/opt/app-rose/merchant-missing/")
	path := pathDir + "example.xlsx"

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
