package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"log"
	"net/http"
	"rose-be-go/constants"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	"rose-be-go/services"
	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
	"strconv"
)

type MdrAggregatorController struct {

}

func (controller *MdrAggregatorController) Save(ctx *gin.Context)  {
	fmt.Println(">>> MdrAggregatorController -- Save <<<")

	var req dto.ReqMdrAggragtorDto
	var res models.Response

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "MdrAggregatorController"

	span := TracingFirstControllerCtx(ctx, req, nameCtrl)
	defer span.Finish()
	c := ctx.Request.Context()
	context := opentracing.ContextWithSpan(c, span)


	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		res.ErrCode = constants.ERR_UNMARSHAL
		res.ErrDesc = constants.ERR_UNMARSHAL_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}

	spanID := ottoutils.GetSpanId(span)
	sugarLogger.Info("REQUEST:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("HEADER", ctx.Request.Header))

	gen := models.GeneralModel{
		ParentSpan: span,
		OttoZaplog: sugarLogger,
		SpanId:     spanID,
		Context:    context,
	}

	reqByte,_ := json.Marshal(req)
	log.Println("req -->", string(reqByte))

	res = services.InitMdrAggregatorService(gen).SaveTemp(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)

}

func (controller *MdrAggregatorController) GetFilterPaging(ctx *gin.Context) {
	fmt.Println(">>> MdrAggregatorController - GetFilterPaging <<<")

	var req dto.ReqMdrAggragtorDto
	var res models.Response

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "MdrAggregatorController"

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		res.ErrCode = constants.ERR_UNMARSHAL
		res.ErrDesc = constants.ERR_UNMARSHAL_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}

	span := TracingFirstControllerCtx(ctx, req, nameCtrl)
	defer span.Finish()
	c := ctx.Request.Context()
	context := opentracing.ContextWithSpan(c, span)

	spanID := ottoutils.GetSpanId(span)
	sugarLogger.Info("REQUEST:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("HEADER", ctx.Request.Header))

	gen := models.GeneralModel{
		ParentSpan: span,
		OttoZaplog: sugarLogger,
		SpanId:     spanID,
		Context:    context,
	}

	log.Println(gen)
	res = services.InitMdrAggregatorService(gen).Filter(req)


	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)

}

func (controller *MdrAggregatorController) GetById(ctx *gin.Context) {
	fmt.Println(">>> MdrAggregatorController - GetById <<<")

	var res models.Response

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "MdrAggregatorController"

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		res.ErrCode = constants.ERR_PARAMETER
		res.ErrDesc = constants.ERR_PARAMETER_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}

	span := TracingFirstControllerCtx(ctx, id, nameCtrl)
	defer span.Finish()
	c := ctx.Request.Context()
	context := opentracing.ContextWithSpan(c, span)

	spanID := ottoutils.GetSpanId(span)
	sugarLogger.Info("REQUEST:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("HEADER", ctx.Request.Header))

	gen := models.GeneralModel{
		ParentSpan: span,
		OttoZaplog: sugarLogger,
		SpanId:     spanID,
		Context:    context,
	}

	log.Println(gen)
	res = services.InitMdrAggregatorService(gen).GetById(id)


	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)

}