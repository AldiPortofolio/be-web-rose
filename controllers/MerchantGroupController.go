package controllers

import (
	"fmt"
	"log"
	"net/http"
	"rose-be-go/constants"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	"rose-be-go/services"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
)

type MerchantGroupController struct {

}

func (controller *MerchantGroupController) GetById(ctx *gin.Context) {
	fmt.Println(">>> MerchantGroupController - GetById <<<")

	var res dto.MerchantGroupDtoRes

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "MerchantGroupController"

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		errCode := constants.ERR_PARAMETER
		errDesc := constants.ERR_PARAMETER_MSG
		res.ErrCode = &errCode
		res.ErrDesc = &errDesc
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
	res = services.InitMerchantGroupService(gen).GetDetail(id)


	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)

}

func (controller *MerchantGroupController) MerchantGroupActivationController(ctx *gin.Context) {
	fmt.Println(">>> MerchantGroupController - MerchantGroupActivationController <<<")

	var req dto.MerchantGroupListIdDto
	var res models.Response

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "MerchantGroupController"

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.ErrCode = constants.ERR_UNMARSHAL
		res.ErrDesc = constants.ERR_UNMARSHAL_MSG
		logs.Error("err unmarshal "+err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	span := TracingFirstControllerCtx(ctx, req.Id, nameCtrl)
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
	res = services.InitMerchantGroupService(gen).MerchantGroupActivationService(req.Id)


	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)

}
