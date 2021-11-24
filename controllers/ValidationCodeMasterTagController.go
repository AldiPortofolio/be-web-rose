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

type ValidationCodeMasterTagController struct {

}

// @Summary ValidationCodeMasterTag - Filter
// @Description ValidationCodeMasterTag Filter Paging
// @ID ValidationCodeMasterTagFilter
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqValidationCodeMasterTagDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response{contents=[]dto.ResValidationCodeMasterTagDto}
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/validation-code-master-tag/filter [post]
func (controller *ValidationCodeMasterTagController) Filter(ctx *gin.Context)  {
	fmt.Println(">>> ValidationCodeMasterTagController - Filter <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	var res models.Response

	var req dto.ReqValidationCodeMasterTagDto

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


	services.InitValidationCodeMasterTagService(logs).Filter(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response ValidationCodeMasterTagController - Filter",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}

// @Summary ValidationCodeMasterTag - Save / Edit
// @Description ValidationCodeMasterTag Save/EDit
// @ID ValidationCodeMasterTagSaveEdit
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqSaveValidationCodeMasterTagDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/validation-code-master-tag [post]
func (controller *ValidationCodeMasterTagController) Save(ctx *gin.Context)  {
	fmt.Println(">>> ValidationCodeMasterTagController - Filter <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req dto.ReqSaveValidationCodeMasterTagDto

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


	services.InitValidationCodeMasterTagService(logs).Save(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response ValidationCodeMasterTagController - Filter",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)



}