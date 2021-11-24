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

type LookupGroupController struct {

}

// @Summary LookupGroup - All
// @Description LookupGroup All
// @ID LookupGroupAll
// @Param Authorization header string true "Bearer"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response{contents=[]dbmodels.LookupGroup}
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/lookup-group/all [get]
func (controller *LookupGroupController) All(ctx *gin.Context)  {
	fmt.Println(">>> LookupGroupController - All <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	var res models.Response


	services.InitLookupGroupService(logs).FindAll(&res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response LookupGroupController - All",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}

// @Summary LookupGroup - Filter
// @Description LookupGroup Filter Paging
// @ID LookupGroupFilter
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqLookupGroupDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response{contents=[]dbmodels.LookupGroup}
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/lookup-group/filter [post]
func (controller *LookupGroupController) Filter(ctx *gin.Context)  {
	fmt.Println(">>> LookupGroupController - Filter <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	var res models.Response

	var req dto.ReqLookupGroupDto

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


	services.InitLookupGroupService(logs).Filter(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response LookupGroupController - Filter",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}

// @Summary LookupGroup - Save / Edit
// @Description LookupGroup Save/EDit
// @ID LookupGroupSave/Edit
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqLookupGroupDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/lookup-group [post]
func (controller *LookupGroupController) Save(ctx *gin.Context)  {
	fmt.Println(">>> LookupGroupController - Save <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req dto.ReqLookupGroupDto

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


	services.InitLookupGroupService(logs).Save(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response LookupGroupController - Save",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)



}