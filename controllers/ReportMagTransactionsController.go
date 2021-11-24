package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"net/http"
	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
	"rose-be-go/constants"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	"rose-be-go/services"
)

// ReportMagTransactionsController ...
type ReportMagTransactionsController struct {
}

// GetAll ...
func (controller *ReportMagTransactionsController) GetAll(ctx *gin.Context) {
	fmt.Println(">>> ReportMagTransactionsController - GetAll <<<")

	var res models.Response
	var req dto.ReqReportMagTransactionsDto
	parent := context.Background()
	defer parent.Done()

	sugarLogger := ottologger.GetLogger()
	ctrlName := "ReportMagTransactionsController"
	span := TracingFirstControllerCtx(ctx, req, ctrlName)
	defer span.Finish()
	c := ctx.Request.Context()
	context := opentracing.ContextWithSpan(c, span)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		res.ErrCode = constants.ERR_UNMARSHAL
		res.ErrDesc = constants.ERR_UNMARSHAL_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	spanID := ottoutils.GetSpanId(span)
	sugarLogger.Info("REQUEST:", zap.String("SPANID", spanID), zap.String("CTRL", ctrlName),
		zap.Any("HEADER", ctx.Request.Header))

	gen := models.GeneralModel{
		ParentSpan: span,
		OttoZaplog: sugarLogger,
		SpanId:     spanID,
		Context:    context,
	}

	res = services.InitReportMagTransactionsService(gen).GetAllReportMagTransactions(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", ctrlName),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)
}

// GetAll ...
func (controller *ReportMagTransactionsController) Export(ctx *gin.Context) {
	fmt.Println(">>> ReportMagTransactionsController - Export <<<")

	var res models.Response
	var req dto.ReqReportMagTransactionsDto
	parent := context.Background()
	defer parent.Done()

	sugarLogger := ottologger.GetLogger()
	ctrlName := "ReportMagTransactionsController"
	span := TracingFirstControllerCtx(ctx, req, ctrlName)
	defer span.Finish()
	c := ctx.Request.Context()
	context := opentracing.ContextWithSpan(c, span)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		res.ErrCode = constants.ERR_UNMARSHAL
		res.ErrDesc = constants.ERR_UNMARSHAL_MSG
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	spanID := ottoutils.GetSpanId(span)
	sugarLogger.Info("REQUEST:", zap.String("SPANID", spanID), zap.String("CTRL", ctrlName),
		zap.Any("HEADER", ctx.Request.Header))

	gen := models.GeneralModel{
		ParentSpan: span,
		OttoZaplog: sugarLogger,
		SpanId:     spanID,
		Context:    context,
	}

	res = services.InitReportMagTransactionsService(gen).Export(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", ctrlName),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)
}




