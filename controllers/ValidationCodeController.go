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

type ValidationCodeController struct {

}


// @Summary ValidationCode - Filter
// @Description ValidationCode Filter Paging
// @ID ValidationCode - Filter
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqValidationCodeDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response{contents=[]dto.ResValidationCodeDto}
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/validation-code/filter [post]
func (controller *ValidationCodeController) Filter(ctx *gin.Context)  {
	fmt.Println(">>> ValidationCodeController - Filter <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req dto.ReqValidationCodeDto

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


	services.InitValidationCodeService(logs).Filter(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response ValidationCodeController - Filter",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}

// @Summary ValidationCode - Save / Edit
// @Description ValidationCode Save/EDit
// @ID ValidationCode - Save/Edit
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqValidationCodeDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/validation-code [post]
func (controller *ValidationCodeController) Save(ctx *gin.Context)  {
	fmt.Println(">>> ValidationCodeController - Filter <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req dto.ReqValidationCodeDto

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


	services.InitValidationCodeService(logs).Save(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response ValidationCodeController - Filter",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)



}