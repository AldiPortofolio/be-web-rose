package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	"rose-be-go/services"
)

type PortalActivationController struct {

}

func (controller *PortalActivationController) Activation(ctx *gin.Context){
	fmt.Println(">>> PortalActivationController - Activation <<<")

	res := models.Response{
		ErrCode: "01",
		ErrDesc: "Failed",
	}

	var req dto.ReqPortalActivation

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "PortalActivation"

	span := TracingFirstControllerCtx(ctx, req, nameCtrl)
	defer span.Finish()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}

	spanID := ottoutils.GetSpanId(span)
	sugarLogger.Info("REQUEST:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("HEADER", ctx.Request.Header))

	res = services.GetPortalActivation(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)
}

func (controller *PortalActivationController) BpActivation(ctx *gin.Context){
	fmt.Println(">>> PortalActivationController - BpActivation <<<")

	res := models.Response{
		ErrCode: "01",
		ErrDesc: "Failed",
	}

	var req dto.BpActivationReq

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "PortalActivation"

	span := TracingFirstControllerCtx(ctx, req, nameCtrl)
	defer span.Finish()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}

	spanID := ottoutils.GetSpanId(span)
	sugarLogger.Info("REQUEST:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("HEADER", ctx.Request.Header))

	res, _ = services.BpActivation(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)
}

func (controller *PortalActivationController) Filter(ctx *gin.Context) {

	res := models.Response{
		ErrCode: "01",
		ErrDesc: "Failed",
	}

	var req dto.ReqPortalListAccountFilter

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "PortalListActivationFilter"

	span := TracingFirstControllerCtx(ctx, req, nameCtrl)
	defer span.Finish()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}

	spanID := ottoutils.GetSpanId(span)
	sugarLogger.Info("REQUEST:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("HEADER", ctx.Request.Header))

	res = services.GetPortalListAccountFilter(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)
}

func (controller *PortalActivationController) CallBack(ctx *gin.Context){
	fmt.Println(">>> PortalActivationController - CallBack <<<")

	res := models.Response{
		ErrCode: "01",
		ErrDesc: "Failed",
	}
	var req dto.ReqPortalCallback

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "PortalCallback"

	span := TracingFirstControllerCtx(ctx, req, nameCtrl)
	defer span.Finish()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}

	spanID := ottoutils.GetSpanId(span)
	sugarLogger.Info("REQUEST:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("HEADER", ctx.Request.Header))

	res = services.GetPortalCallback(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)
}

func (controller *PortalActivationController) Reset(ctx *gin.Context){
	fmt.Println(">>> PortalActivationController - Reset Password <<<")

	res := models.Response{
		ErrCode: "01",
		ErrDesc: "Failed",
	}
	var req dto.ReqPortalActivation

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "PortalResetPasssword"

	span := TracingFirstControllerCtx(ctx, req, nameCtrl)
	defer span.Finish()



	spanID := ottoutils.GetSpanId(span)
	sugarLogger.Info("REQUEST:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("HEADER", ctx.Request.Header))

	res = services.GetPortalReset(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)
}

func (controller *PortalActivationController) ResetPassword(ctx *gin.Context){
	fmt.Println(">>> PortalActivationController - Reset Password BP <<<")

	res := models.Response{
		ErrCode: "01",
		ErrDesc: "Failed",
	}

	var req dto.ReqPortalActivation

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "PortalResetPasssword"

	span := TracingFirstControllerCtx(ctx, req, nameCtrl)
	defer span.Finish()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}

	spanID := ottoutils.GetSpanId(span)
	sugarLogger.Info("REQUEST:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("HEADER", ctx.Request.Header))

	res = services.ResetPassword(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)
}

func (controller *PortalActivationController) FilterOutlet(ctx *gin.Context){
	fmt.Println(">>> PortalActivationController - Filter Outlet <<<")

	res := models.Response{}
	var req dto.ReqFilterOutlet

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "FilterOutlet"

	span := TracingFirstControllerCtx(ctx, req, nameCtrl)
	defer span.Finish()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		ctx.JSON(http.StatusOK, res)
		return
	}

	spanID := ottoutils.GetSpanId(span)
	sugarLogger.Info("REQUEST:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("HEADER", ctx.Request.Header))

	res = services.FilterOutletPortal(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)
}
