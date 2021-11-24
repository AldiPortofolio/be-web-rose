package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rose-be-go/constants"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	"rose-be-go/services"

	"github.com/gin-gonic/gin"
	"ottodigital.id/library/logger/v2"
)


type AppUserController struct {

}

func (controller *AppUserController) Update(ctx *gin.Context)  {
	fmt.Println(">>> AppUserController - Update <<<")

	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req dto.ReqAppUserDto

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


	services.InitAppUserService(logs).Save(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response MasterTag Controller - Filter",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}
