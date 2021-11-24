package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"log"
	"net/http"
	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
	"rose-be-go/services"
	"strconv"

	"rose-be-go/constants"
	"rose-be-go/models"
)

type MerchantWipController struct {

}

func (controller *MerchantWipController) Approve(ctx *gin.Context)  {
	fmt.Println(">>> MerchantWipController - Approve <<<")

	parent := context.Background()
	defer parent.Done()

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "MerchantWipController"

	res := models.Response{
		ErrCode: "01",
		ErrDesc: "Transaction failed",
	}

	action, errAction := strconv.Atoi(ctx.Param("action"))
	if errAction != nil {
		log.Println("error", errAction)
		res.ErrCode = constants.ERR_UNMARSHAL
		res.ErrDesc = constants.ERR_UNMARSHAL_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}

	var req models.MerchantWip
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.ErrCode = constants.ERR_UNMARSHAL
		res.ErrDesc = constants.ERR_UNMARSHAL_MSG
		go sugarLogger.Error("Body request error ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, res)
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
		zap.Any("BODY", req),
		zap.Any("HEADER", ctx.Request.Header))

	gen := models.GeneralModel{
		ParentSpan: span,
		OttoZaplog: sugarLogger,
		SpanId:     spanID,
		Context:    context,
	}

	switch action {
	case 0:
		// data masuk dari Rest ottoSFA ottoMerchant
		// data save
		
	case 5:
		// REGISTERED - REGISTERED
		// 1. tarik dari topix registered
		res = services.InitMerchantWipService(gen).GetNextRegistered()
		
	case 99:

		
	}


	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)


}