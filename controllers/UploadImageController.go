package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rose-be-go/constants"
	"rose-be-go/hosts/minio"
	"rose-be-go/models"
	"rose-be-go/services"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
)

type UploadImageController struct {
}

// @Summary UploadImage
// @Description Upload Image
// @ID Upload Image
// @Param Authorization header string true "Bearer"
// @Param body body minio.UploadRequest true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/UploadImage [post]
func (controller *UploadImageController) Upload(ctx *gin.Context) {
	fmt.Println(">>> UploadImageController - Upload <<<")

	var req minio.UploadRequest
	var res models.Response

	sugarLogger := ottologger.GetLogger()
	nameCtrl := "UploadImageController"

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

	res = services.InitUploadImageService(gen).Upload(req)

	sugarLogger.Info("RESPONSE:", zap.String("SPANID", spanID), zap.String("CTRL", nameCtrl),
		zap.Any("BODY", res))

	ctx.JSON(http.StatusOK, res)

}


