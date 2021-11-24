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
	ottoutils "ottodigital.id/library/utils"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	"rose-be-go/services"
)

type UploadMerchantMasterTagController struct {

}

func (controller *UploadMerchantMasterTagController) Upload(ctx *gin.Context)  {
	fmt.Println("UploadMerchantMasterTagController - Upload")
	var res models.Response
	logs := logger.InitLogs(ctx.Request)


	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println("Request body error:", err)
		ctx.JSON(http.StatusBadRequest, res)
		log.Println("Body request error ", err.Error())
		return
	}


	services.InitUploadMerchantMasterTagService().UploadFile(ctx, file, &res)

	bodyRes, _ := json.Marshal(res)
	logs.Info("Response MasterTag Controller - Upload",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}

func (controller *UploadMerchantMasterTagController) GetFilterPaging(ctx *gin.Context)  {
	fmt.Println("UploadMerchantMasterTagController - GetFilterPaging")
	logs := logger.InitLogs(ctx.Request)
	var req dto.ReqUploadMerchantMasterTagDto
	var res models.Response


	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}

	reqByte,_ := json.Marshal(req)
	log.Println("req --> ", string(reqByte))

	services.InitUploadMerchantMasterTagService().GetDataUploadMerchant(req,&res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response MasterTag Controller - Filter",
		logs.AddField("ResponseBody: ", string(bodyRes)))



	ctx.JSON(http.StatusOK, res)

}

func (controller *UploadMerchantMasterTagController) HandleDownload(ctx *gin.Context) {

	req := models.ReqDowloadFile{}

	res := models.Response{}


	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}


	pathDir := ottoutils.GetEnv("ROSE_DOWNLOAD_MERCHANT_MASTER_TAG_PATH","/opt/app-rose/merchant-master-tag/process/")
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


func (controller *UploadMerchantMasterTagController) HandleResultDownload(ctx *gin.Context) {

	req := models.ReqDowloadFile{}

	res := models.Response{}


	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}

	pathDir := ottoutils.GetEnv("PATH_DOWNLOAD_RESULT_MERCHANT_MASTER_TAG","/opt/app-rose/merchant-master-tag/result/")
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

func (controller *UploadMerchantMasterTagController) HandleDownloadExample(ctx *gin.Context) {

	pathDir := ottoutils.GetEnv("PATH_DOWNLOAD_EXAMPLE_MERCHANT_MASTER_TAG","/opt/app-rose/merchant-master-tag/")
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