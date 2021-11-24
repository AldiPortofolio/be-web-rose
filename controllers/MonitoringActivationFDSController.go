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
	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
	"rose-be-go/services"
)

type MonitoringActivationFDSController struct {

}

// @Summary Monitoring Activation FDS - Filter
// @Description Monitoring Activation FDS Filter
// @ID Monitoring Activation FDS - Filter
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqMonitoringActivationFDSDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response{contents=[]dbmodels.LimitTransaction}
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/lmonitoring-activation-fds/filter [post]
func (controller *MonitoringActivationFDSController) GetFilterPaging(ctx *gin.Context)  {
	fmt.Println(">>> MonitoringActivationFDSController - GetFilterPaging <<<")

	var req dto.ReqMonitoringActivationFDSDto
	var res models.Response

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "MonitoringActivationFDSController"

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		res.ErrCode = constants.ERR_UNMARSHAL
		res.ErrDesc = constants.ERR_UNMARSHAL_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}
	reqByte,_ := json.Marshal(req)
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
	res = services.InitMonitoringActivationFDSService(gen).Filter(req)


	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)

}
