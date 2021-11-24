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

type UpdatedDataMerchantController struct {

}


// @Summary UpdatedDataMerchant - Filter
// @Description UpdatedDataMerchant - Filter
// @ID UpdatedDataMerchantFilter
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqUpdatedDataMerchantDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response{contents=[]dbmodels.UpdatedDataMerchant}
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/updated-data/filter [post]
func (controller *UpdatedDataMerchantController) Filter(ctx *gin.Context)  {
	fmt.Println(">>> UpdatedDataMerchantController - Filter <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req dto.ReqUpdatedDataMerchantDto

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


	services.InitUpdatedDataMerchantService(logs).Filter(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response UpdatedDataMerchantController - Filter",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}

// @Summary UpdatedDataMerchant - Approve
// @Description UpdatedDataMerchant Approve
// @ID UpdatedDataMerchantApprove
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqUpdateDataMerchantDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/updated-data/approve [post]
func (controller *UpdatedDataMerchantController) Approve(ctx *gin.Context)  {
	fmt.Println(">>> UpdatedDataMerchantController - Approve <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req dto.ReqUpdateDataMerchantDto

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


	services.InitUpdatedDataMerchantService(logs).Approve(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response UpdatedDataMerchantController - Approve",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}


// @Summary UpdatedDataMerchant - Reject
// @Description UpdatedDataMerchant Reject
// @ID UpdatedDataMerchantReject
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqUpdateDataMerchantDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/updated-data/reject [post]
func (controller *UpdatedDataMerchantController) Reject(ctx *gin.Context)  {
	fmt.Println(">>> UpdatedDataMerchantController - Reject <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req dto.ReqUpdateDataMerchantDto

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


	services.InitUpdatedDataMerchantService(logs).Reject(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response UpdatedDataMerchantController - Reject",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}