package controllers

import (
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
	"rose-be-go/services"
	"strconv"
)

type LevelMerchantController struct {
}

func (controller *LevelMerchantController) GetAll(ctx *gin.Context) {
	fmt.Println(">>> LevelMerchantController - GetAll <<<")

	var res models.Response

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "LevelMerchantController"

	limit, err := strconv.ParseInt(ctx.Param("limit"), 10, 64)
	if err != nil {
		res.ErrCode = constants.ERR_PARAMETER
		res.ErrDesc = constants.ERR_PARAMETER_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}

	span := TracingFirstControllerCtx(ctx, limit, nameCtrl)
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
	res = services.InitLevelMerchantService(gen).GetAll(limit)


	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)

}
