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
	"rose-be-go/models/dto"
	"rose-be-go/services"
)

type BlastNotifController struct {

}


// @Summary BlastNotif Send All
// @Description BlastNotif Send All
// @ID BlastNotifSendAll
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqBlastNotificationSendAllDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/blast-notif/send-all [post]
func (controller *BlastNotifController) SendAll(ctx *gin.Context)  {
	fmt.Println(">>> BlastNotifController - Send All <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	var res models.Response
	var req dto.ReqBlastNotificationSendAllDto

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


	services.InitBlastNotifService().SendAll(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response BlastNotifController - Send All",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}