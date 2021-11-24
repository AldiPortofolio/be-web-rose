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

type MasterTagController struct {

}

// @Summary MasterTag - All
// @Description MasterTag AllData
// @ID MasterTagAllData
// @Param Authorization header string true "Bearer"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response{contents=[]dbmodels.MasterTag}
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/master-tag/all [get]
func (controller *MasterTagController) All(ctx *gin.Context)  {
	fmt.Println(">>> MasterTagController - All <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	var res models.Response


	services.InitMasterTagService(logs).GetAll(&res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response MasterTag Controller - All",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}


// @Summary MasterTag - Filter
// @Description MasterTag Filter Paging
// @ID MasterTagFilter
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqMasterTagDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response{contents=[]dbmodels.MasterTag}
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/master-tag/filter [post]
func (controller *MasterTagController) Filter(ctx *gin.Context)  {
	fmt.Println(">>> MasterTagController - Filter <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	var res models.Response

	var req dto.ReqMasterTagDto

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


	services.InitMasterTagService(logs).Filter(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response MasterTag Controller - Filter",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}

// @Summary MasterTag - Save / Edit
// @Description MasterTag Save/EDit
// @ID MasterTag - Save/Edit
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqMasterTagDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/master-tag [post]
func (controller *MasterTagController) Save(ctx *gin.Context)  {
	fmt.Println(">>> MasterTagController - Filter <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req dto.ReqMasterTagDto

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


	services.InitMasterTagService(logs).Save(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response MasterTag Controller - Filter",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)



}