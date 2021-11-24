package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rose-be-go/constants"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
)

type ConfigVaMerchantGroupController struct {
}

// Save ...
// @Summary Config VA Merchant Group - Save
// @Description Config VA Merchant Group - Save
// @ID Config VA Merchant Group - Save
// @Param Authorization header string true "Bearer"
// @Param body body dbmodels.ConfigVaMerchantGroup true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/config-va-merchant-group/save [post]
func (controller *ConfigVaMerchantGroupController) Save(ctx *gin.Context) {
	fmt.Println(">>> ConfigVaMerchantGroupController - Save <<<")

	var req dbmodels.ConfigVaMerchantGroup
	var res models.Response

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "ConfigVaMerchantGroupController"

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

	res = services.InitConfigVaMerchantGroupService(gen).Save(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)

}

// SaveToRedis ...
// @Summary Config VA Merchant Group - SaveToRedis
// @Description Config VA Merchant Group - SaveToRedis
// @ID Config VA Merchant Group - SaveToRedis
// @Param Authorization header string true "Bearer"
// @Param body body dbmodels.ConfigVaMerchantGroup true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/config-va-merchant-group/save-to-redis [post]
func (controller *ConfigVaMerchantGroupController) SaveToRedis(ctx *gin.Context) {
	fmt.Println(">>> ConfigVaMerchantGroupController - SaveToRedis <<<")

	var req dbmodels.ConfigVaMerchantGroup
	var res models.Response

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "ConfigVaMerchantGroupController"

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

	res = services.InitConfigVaMerchantGroupService(gen).SaveToRedis(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)

}

// FindByGroupId ...
// @Summary Config VA Merchant Group - FindByGroupId
// @Description Config VA Merchant Group - FindByGroupId
// @ID Config VA Merchant Group - FindByGroupId
// @Param Authorization header string true "Bearer"
// @Param id path string true "merchant group id"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/config-va-merchant-group/find-by-group-id/{id} [get]
func (controller *ConfigVaMerchantGroupController) FindByGroupId(ctx *gin.Context) {
	fmt.Println(">>> ConfigVaMerchantGroupController - FindByGroupId <<<")
	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "ConfigVaMerchantGroupController"

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

	res = services.InitConfigVaMerchantGroupService(gen).FindByGroupId(int64(reqId))

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)

}

// FindByGroupIdFromRedis ...
// @Summary Config VA Merchant Group - FindByGroupIdFromRedis
// @Description Config VA Merchant Group - FindByGroupIdFromRedis
// @ID Config VA Merchant Group - FindByGroupIdFromRedis
// @Param Authorization header string true "Bearer"
// @Param id path string true "merchant group id"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/config-va-merchant-group/find-by-group-id-from-redis/{id} [get]
func (controller *ConfigVaMerchantGroupController) FindByGroupIdFromRedis(ctx *gin.Context) {
	fmt.Println(">>> ConfigVaMerchantGroupController - FindByGroupIdFromRedis <<<")
	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "ConfigVaMerchantGroupController"

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

	res = services.InitConfigVaMerchantGroupService(gen).FindByGroupIdFromRedis(int64(reqId))

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)

}
