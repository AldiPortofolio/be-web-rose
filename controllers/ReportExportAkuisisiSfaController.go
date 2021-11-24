package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"rose-be-go/constants"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"rose-be-go/services"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
)


type ReportExportAkuisisiSfaController struct {

}

func (controller *ReportExportAkuisisiSfaController) GetReportExportAkuisisiSfa(ctx *gin.Context) {
	fmt.Println(">>> Get Report Export Akuisisi Sfa Data <<<")

	req := dbmodels.ReportExportAkuisisiSfa{}

	res := models.Response{
		Contents: make([]interface{}, 0),
	}

	sugarLogger := ottologger.GetLogger()
	ctrlName := "Check Data Report Export Akuisisi Sfa"

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

	res = services.GetDataReportExportAkuisisiSfa(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", ctrlName),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)
}

// Send ...
func (controller *ReportExportAkuisisiSfaController) Send(ctx *gin.Context) {
	fmt.Println(">>> ReportExportAkuisisiSfaController - Send <<<")

	parent := context.Background()
	defer parent.Done()

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "ReportExportAkuisisiSfaController"

	res := models.Response{
		ErrCode: "01",
		ErrDesc: "Transaction failed",
	}
	var req dto.ReqReportExportAkuisisiSfaDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.ErrCode = constants.ERR_UNMARSHAL
		res.ErrDesc = constants.ERR_UNMARSHAL_MSG
		go sugarLogger.Error("Body request error ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	req.Topic = "rose-report-export-akuisisi-sfa-topic"
	reqByte, _ := json.Marshal(req)
	log.Println("req --> ", string(reqByte))

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

	//log.Println(gen)
	res = services.InitReportExportAkuisisiSfaService(gen).Send(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)

}

// DownloadPath ...
func (controller *ReportExportAkuisisiSfaController) DownloadPath(ctx *gin.Context) {

	req := models.ReqDowloadFile{}
	res := models.Response{}

	sugarLogger := ottologger.GetLogger()
	ctrlName := "CheckFileReportExportAkuisisiSfa"

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

	path := ottoutils.GetEnv("PATH_DOWNLOAD_REPORT_EXPORT_AKUISISI_SFA","/opt/app-rose/report/export_akuisisi_sfa/")
	// path := ottoutils.GetEnv("PATH_DOWNLOAD_REPORT_EXPORT_AKUISISI_SFA","/Users/abdulah/Documents/testing/report/export_akuisisi_sfa/")

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
