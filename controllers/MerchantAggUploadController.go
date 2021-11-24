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
	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
	"rose-be-go/constants"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"rose-be-go/services"
)

type MerchantAggUploadController struct {

}

func (controller *MerchantAggUploadController) UploadFile(ctx *gin.Context){
	fmt.Println("MerchantAggUploadController")

	parent := context.Background()
	defer parent.Done()

	res := models.Response{
		ErrCode:  "01",
		ErrDesc: "Transaction failed",
	}

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "MerchantAggUploadController"

	file, err := ctx.FormFile("file")
	mid := ctx.PostForm("mid_aggregator")
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

	go services.InitMerchantAggUpload(gen).UplaodFile(ctx, file, mid)

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.Data = "On Progress"


	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)
}


func (controller *MerchantAggUploadController) GetDataMerchantAggUpload(ctx *gin.Context) {
	fmt.Println(">>> Merchant Agg Upload Data <<<")

	req := dbmodels.MerchantAggUpload{}

	res := models.Response{
		Contents: make([]interface{}, 0),
	}

	sugarLogger := ottologger.GetLogger()
	ctrlName := "CheckMerchantAggData"

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

	res = services.GetDataMerchantAggUpload(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", ctrlName),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)
}

func (controller *MerchantAggUploadController) Download(ctx *gin.Context) {

	req := models.ReqDowloadFile{}

	res := models.Response{}

	sugarLogger := ottologger.GetLogger()
	ctrlName := "MerchantAggDownload"

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
	//path := "/opt/app-rose/nmid/proccess/" + req.FilePath

	pathDir := ottoutils.GetEnv("PATH_MERCHANT_AGG_UPLOAD","/opt/app-rose/merchant-agg/upload/")
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


func (controller *MerchantAggUploadController) ApproveMerchantAggregatorUpload(ctx *gin.Context) {
	fmt.Println(">>> Merchant Agg Approve Data <<<")

	req := dto.ReqMerchantAggregatorUpload{}

	res := models.Response{
		Contents: make([]interface{}, 0),
	}

	sugarLogger := ottologger.GetLogger()
	ctrlName := "ApproveMerchantAggregatorUpload"

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}
	span := TracingFirstControllerCtx(ctx, nil, ctrlName)
	defer span.Finish()
	c := ctx.Request.Context()
	context := opentracing.ContextWithSpan(c, span)

	spanID := ottoutils.GetSpanId(span)
	sugarLogger.Info("REQUEST:", zap.String("SPANID", spanID), zap.String("CTRL", ctrlName),
		zap.Any("HEADER", ctx.Request.Header))

	gen := models.GeneralModel{
		ParentSpan: span,
		OttoZaplog: sugarLogger,
		SpanId:     spanID,
		Context:    context,
	}

	res = services.InitMerchantAggUpload(gen).Approve(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", ctrlName),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)
}

func (controller *MerchantAggUploadController) HandleTemplateDownload(ctx *gin.Context) {


	sugarLogger := ottologger.GetLogger()
	ctrlName := "MerchantAggUploadController"

	span := TracingFirstControllerCtx(ctx, nil, ctrlName)
	defer span.Finish()



	spanID := ottoutils.GetSpanId(span)
	sugarLogger.Info("REQUEST:", zap.String("SPANID", spanID), zap.String("CTRL", ctrlName),
		zap.Any("HEADER", ctx.Request.Header))


	pathDir := ottoutils.GetEnv("PATH_DOWNLOAD_TEMPLATE_UPLOAD_MERCHANT_AGGREGATOR","/opt/app-rose/merchant-agg/")
	path := pathDir + "TemplateUploadAgg.csv"

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

