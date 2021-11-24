package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"log"
	"net/http"
	"rose-be-go/models"
	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
	"rose-be-go/services"
)
type KategoriBisnisController struct {

}

func ( controller *KategoriBisnisController ) FindAll ( ctx *gin.Context ) {
	fmt.Println(">>> KategoriBisnisController - FindAll <<<")

	var res models.Response

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "KategoriBisnisController.FindAll"

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
	res = services.InitKategoriBisnisService(gen).FindAll()

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)

	

	

}