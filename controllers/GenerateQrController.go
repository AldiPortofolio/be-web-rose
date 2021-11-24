package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/constants"
	"rose-be-go/models"
	"rose-be-go/models/sbmodels"
	"rose-be-go/services"
)

type GenerateQrController struct {

}


func (controller *GenerateQrController) Generate(ctx *gin.Context)  {
	fmt.Println(">>> GenerateQrController - Generate <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req sbmodels.ReqGenerateQr

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.ErrCode = constants.ERR_UNMARSHAL
		res.ErrDesc = constants.ERR_UNMARSHAL_MSG
		logs.Error("err unmarshal "+err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	reqByte,_ := json.Marshal(req)
	log.Println("req --> ", string(reqByte))


	services.InitGenerateQrService(logs).GenerateQr(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response  GenerateQrController - Generate ",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}
