package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"rose-be-go/constants"
	"rose-be-go/models/dto"
	ottoutils "ottodigital.id/library/utils"

	"rose-be-go/services"
	"rose-be-go/models"

)

type ReportCrmPenugasanController struct {

}

// Send ...
func (controller *ReportCrmPenugasanController) Send(ctx *gin.Context) {
	fmt.Println(">>> ReportCrmPenugasanController - Send <<<")


	res := models.Response{
		ErrCode: "01",
		ErrDesc: "Transaction failed",
	}
	var req dto.ReqReportCrmPenugasanDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.ErrCode = constants.ERR_UNMARSHAL
		res.ErrDesc = constants.ERR_UNMARSHAL_MSG

		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	req.Topic = "rose-report-crm-penugasan-topic"
	reqByte, _ := json.Marshal(req)
	log.Println("req --> ", string(reqByte))

	res = services.InitReportCrmPenugasanService().Send(req)



	ctx.JSON(http.StatusOK, res)

}

func (controller *ReportCrmPenugasanController) GetReport(ctx *gin.Context) {
	fmt.Println(">>> Get Report Finished Data <<<")

	req := dto.ReqGetReporCrmPenugasanDto{}

	res := models.Response{
		Contents: make([]interface{}, 0),
	}


	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}

	res = services.InitReportCrmPenugasanService().FilterPaging(req)


	ctx.JSON(http.StatusOK, res)
}


// DownloadPath ...
func (controller *ReportCrmPenugasanController) DownloadPath(ctx *gin.Context) {

	req := models.ReqDowloadFile{}
	res := models.Response{}

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		fmt.Println("Failed bind Request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}



	path := ottoutils.GetEnv("PATH_DOWNLOAD_REPORT_CRM_PENUGASAN","/opt/app-rose/report/crm-penugasan/")

	file := path + req.FilePath

	if _, err := os.Stat(path); err != nil {
		fmt.Println("create new folder")
		os.MkdirAll(path, os.ModePerm)
	}

	fmt.Println("File Path:", file)

	w := ctx.Writer
	f, err := os.Open(file)

	if f != nil {
		defer f.Close()
	}

	if err != nil {
		res.ErrCode = "01"
		res.ErrDesc = "File Not Found"
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Error(w,"File Not Found", http.StatusInternalServerError)
		//ctx.JSON(http.StatusOK, res)
		fmt.Println("Failed to open file", err)
		return
	}

	contentDisposition := fmt.Sprintf("attachment; filename=%s", f.Name())
	w.Header().Set("Content-Disposition", contentDisposition)

	if _, err := io.Copy(w, f); err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		res.ErrCode = "01"
		res.ErrDesc = "File Not Found"
		http.Error(w,"File Not Found", http.StatusInternalServerError)
		//ctx.JSON(http.StatusOK, res)
		fmt.Println("Failed to copy file", err)
		return
	}
}