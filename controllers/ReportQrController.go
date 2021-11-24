package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
	"rose-be-go/models"
	"rose-be-go/services"
)

type ReportQrController struct {

}

// GetReportFinished ...
func (controller *ReportQrController) GetReportQr(ctx *gin.Context) {
	fmt.Println(">>> Get Report Qr Data <<<")

	req := models.Pagination{}

	res := models.Response{
		Contents: make([]interface{}, 0),
	}

	sugarLogger := ottologger.GetLogger()
	ctrlName := "Check Data Report Qr"

	span := TracingFirstControllerCtx(ctx, nil, ctrlName)
	defer span.Finish()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}

	spanID := ottoutils.GetSpanId(span)
	sugarLogger.Info("REQUEST:", zap.String("SPANID", spanID), zap.String("CTRL", ctrlName),
		zap.Any("HEADER", ctx.Request.Header))

	res = services.GetDataReportQr(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", ctrlName),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)
}

// DownloadPath ...
func (controller *ReportQrController) DownloadPath(ctx *gin.Context) {

	req := models.ReqDowloadFile{}
	res := models.Response{}

	sugarLogger := ottologger.GetLogger()
	ctrlName := "CheckFileReportQr"

	span := TracingFirstControllerCtx(ctx, nil, ctrlName)
	defer span.Finish()

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		fmt.Println("Failed bind Request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}

	spanID := ottoutils.GetSpanId(span)
	sugarLogger.Info("REQUEST:", zap.String("SPANID", spanID), zap.String("CTRL", ctrlName),
		zap.Any("HEADER", ctx.Request.Header))

	//path := "/opt/app-rose/nmid/proccess/"
	path := ottoutils.GetEnv("PATH_DOWNLOAD_NMID","/opt/app-rose/nmid/proccess/")

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

// DownloadPath ...
func (controller *ReportQrController) DownloadResultPath(ctx *gin.Context) {

	req := models.ReqDowloadFile{}
	res := models.Response{}

	sugarLogger := ottologger.GetLogger()
	ctrlName := "CheckFileResultReportQr"

	span := TracingFirstControllerCtx(ctx, nil, ctrlName)
	defer span.Finish()

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		fmt.Println("Failed bind Request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}

	spanID := ottoutils.GetSpanId(span)
	sugarLogger.Info("REQUEST:", zap.String("SPANID", spanID), zap.String("CTRL", ctrlName),
		zap.Any("HEADER", ctx.Request.Header))

	//path := "/opt/app-rose/nmid/result/qr/"
	path := ottoutils.GetEnv("PATH_DOWNLOAD_RESULT_QR","/opt/app-rose/nmid/result/qr/")

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
		http.Error(w, err.Error(), http.StatusInternalServerError)
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