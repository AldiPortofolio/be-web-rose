package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
	"rose-be-go/models"
	"rose-be-go/constants"
	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
	"rose-be-go/models/dto"
	"rose-be-go/services"
)

type UploadMerchantActivatedController struct {

}

func (controller *UploadMerchantActivatedController) Upload(ctx *gin.Context)  {
	fmt.Println("UploadMerchantActivatedController")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{
		ErrCode:  constants.EC_TRANSACTION_FAILED,
		ErrDesc: constants.EC_TRANSACTION_FAILED_DESC,
	}

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "UploadMerchantActivatedController"

	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println("Request body error:", err)
		ctx.JSON(http.StatusBadRequest, res)
		go sugarLogger.Error("Body request error ", zap.Error(err))
		return
	}

	span := TracingFirstControllerCtx(ctx, file.Filename, nameCtrl)
	defer span.Finish()
	c := ctx.Request.Context()
	context := opentracing.ContextWithSpan(c, span)

	spanID := ottoutils.GetSpanId(span)
	sugarLogger.Info("REQUEST:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", file.Filename),
		zap.Any("HEADER", ctx.Request.Header))

	gen := models.GeneralModel{
		ParentSpan: span,
		OttoZaplog: sugarLogger,
		SpanId:     spanID,
		Context:    context,
	}

	res = services.InitUploadMerchantActivatedService(gen).UploadFile(ctx, file)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)

}


func (controller *UploadMerchantActivatedController) GetFilterPaging(ctx *gin.Context)  {
	fmt.Println("GetFilterPaging")
	parent := context.Background()
	defer parent.Done()

	var req dto.ReqUploadMerchant
	res := models.Response{
		ErrCode:  constants.EC_TRANSACTION_FAILED,
		ErrDesc: constants.EC_TRANSACTION_FAILED_DESC,
	}

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "UploadMerchantController"

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}

	span := TracingFirstControllerCtx(ctx, req, nameCtrl)
	defer span.Finish()
	c := ctx.Request.Context()
	context := opentracing.ContextWithSpan(c, span)

	spanID := ottoutils.GetSpanId(span)
	sugarLogger.Info("REQUEST:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", req),
		zap.Any("HEADER", ctx.Request.Header))

	gen := models.GeneralModel{
		ParentSpan: span,
		OttoZaplog: sugarLogger,
		SpanId:     spanID,
		Context:    context,
	}

	res = services.InitUploadMerchantActivatedService(gen).GetDataUploadMerchant(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)

}


func (controller *UploadMerchantActivatedController) HandleDownload(ctx *gin.Context) {

	req := models.ReqDowloadFile{}

	res := models.Response{}

	sugarLogger := ottologger.GetLogger()
	ctrlName := "UploadMerchantActivatedController"

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


	pathDir := ottoutils.GetEnv("PATH_DOWNLOAD_MERCHANT_ACTIVATED","/opt/app-rose/merchant-activated/process/")
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


func (controller *UploadMerchantActivatedController) HandleTemplateDownload(ctx *gin.Context) {


	sugarLogger := ottologger.GetLogger()
	ctrlName := "UploadMerchantActivatedController"

	span := TracingFirstControllerCtx(ctx, nil, ctrlName)
	defer span.Finish()



	spanID := ottoutils.GetSpanId(span)
	sugarLogger.Info("REQUEST:", zap.String("SPANID", spanID), zap.String("CTRL", ctrlName),
		zap.Any("HEADER", ctx.Request.Header))


	pathDir := ottoutils.GetEnv("PATH_DOWNLOAD_TEMPLATE_MERCHANT_ACTIVATED","/opt/app-rose/merchant-activated/")
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


func (controller *UploadMerchantActivatedController) HandleResultDownload(ctx *gin.Context) {

	req := models.ReqDowloadFile{}

	res := models.Response{}

	sugarLogger := ottologger.GetLogger()
	ctrlName := "UploadMerchantActivatedController"

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

	//name := strings.Replace(req.FilePath, ":", "_", -1)

	//path := "/apps/merchant/nmid/" + req.FilePath
	//path := "/opt/app-rose/nmid/result/nmid/" + req.FilePath
	pathDir := ottoutils.GetEnv("PATH_DOWNLOAD_RESULT_MERCHANT_ACTIVATED","/opt/app-rose/merchant-activated/result/")
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