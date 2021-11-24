package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rose-be-go/constants"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	"rose-be-go/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
)

type AcquititionsController struct {
}

// @Summary Acquititons - Filter
// @Description Acquititons Filter
// @ID Acquititons - Filter
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqFilterAcquititionsDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response{contents=[]dto.ReqAcquititionsDto}
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/acquititions/filter [post]
func (controller *AcquititionsController) GetFilterPaging(ctx *gin.Context) {
	fmt.Println(">>> AcquititionsController - GetFilterPaging <<<")

	var req dto.ReqFilterAcquititionsDto
	var res models.Response

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "AcquititionsController"

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		res.ErrCode = constants.ERR_UNMARSHAL
		res.ErrDesc = constants.ERR_UNMARSHAL_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}
	reqByte, _ := json.Marshal(req)
	log.Println("req --> ", string(reqByte))

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
	res = services.InitAcquititionsService(gen).Filter(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)

}

// @Summary Acquititons - Save
// @Description Acquititons Save
// @ID Acquititons - Save
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqAcquititionsDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response{contents=[]dto.ResAcquititionsDto}
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/acquititions [post]
func (controller *AcquititionsController) Save(ctx *gin.Context) {
	fmt.Println(">>> AcquititionsController - Save <<<")

	var req dto.ReqAcquititionsDto
	var res models.Response

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "AcquititionsController"

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		res.ErrCode = constants.ERR_UNMARSHAL
		res.ErrDesc = constants.ERR_UNMARSHAL_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}
	reqByte, _ := json.Marshal(req)
	log.Println("req --> ", string(reqByte))

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

	res = services.InitAcquititionsService(gen).Save(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)

}

// @Summary Acquititons - Delete
// @Description Acquititons Delete
// @ID Acquititons - Delete
// @Param Authorization header string true "Bearer"
// @Param id path int true "ID of the order to be deleted"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/acquititions/delete/{id} [delete]
func (controller *AcquititionsController) Delete(ctx *gin.Context) {
	fmt.Println(">>> AcquititionsController - Delete <<<")
	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "AcquititionsController"

	id := ctx.Param("id")
	log.Println("req Id --> ", id)
	reqId, _ := strconv.Atoi(id)

	span := TracingFirstControllerCtx(ctx, reqId, nameCtrl)
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

	res = services.InitAcquititionsService(gen).Delete(reqId)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)

}
