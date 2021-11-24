package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"log"
	"net/http"
	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
	"rose-be-go/constants"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	"rose-be-go/services"
)

// FeeMdrSettingMerchantGroupController ..
type FeeMdrSettingMerchantGroupController struct {
	// struct attributes
}

// FindAll ..
func (controller *FeeMdrSettingMerchantGroupController) FindAll(ctx *gin.Context) {
	fmt.Println(">>> FeeMdrSettingMerchantGroupController - FindAll <<<")

	var res models.Response

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "FeeMdrSettingMerchantGroupController.FindAll"

	span := TracingEmptyFirstControllerCtx(ctx, nameCtrl)
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
	res = services.InitFeeMdrSettingMerchantGroupService(gen).FindAll()

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)
}

// GetFilterPaging ..
func (controller *FeeMdrSettingMerchantGroupController) GetFilterPaging(ctx *gin.Context) {
	fmt.Println(">>> FeeMdrSettingMerchantGroupController - GetFilterPaging <<<")

	var req dto.ReqFeeMdrSettingMerchantGroupDto
	var res models.Response

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "FeeMdrSettingMerchantGroupController.GetFilterPaging"

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
	res = services.InitFeeMdrSettingMerchantGroupService(gen).Filter(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)
}

// Save ..
func (controller *FeeMdrSettingMerchantGroupController) Save(ctx *gin.Context) {
	fmt.Println(">>> FeeMdrSettingMerchantGroupController - Save <<<")

	var req dto.ReqFeeMdrSettingMerchantGroupDto
	var res models.Response

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "FeeMdrSettingMerchantGroupController.Save"

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

	res = services.InitFeeMdrSettingMerchantGroupService(gen).Save(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)
}
