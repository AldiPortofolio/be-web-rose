package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"net/http"
	"rose-be-go/models"
	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
	"rose-be-go/services"
)

func TestController(ctx *gin.Context)  {
	fmt.Println("Test Controller")
	parent := context.Background()
	defer parent.Done()

	res := models.Response{
		ErrCode:  "01",
		ErrDesc: "Transaction failed",
	}

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "TestController"

	req := models.TestRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		ctx.JSON(http.StatusBadRequest, res)
		go sugarLogger.Error("Body request error ", zap.Error(err))
		return
	}

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

	res = services.InitTestService(gen).GetData(req.Name)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)
}