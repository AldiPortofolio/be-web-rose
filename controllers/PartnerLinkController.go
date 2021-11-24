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
	"strconv"

	"github.com/gin-gonic/gin"
	"ottodigital.id/library/logger/v2"
)

type PartnerLinkController struct {

}


// @Summary PartnerLink - Filter
// @Description PartnerLink Filter Paging
// @ID PartnerLink - Filter
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqPartnerLinkDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response{contents=[]dbmodels.PartnerLink}
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/partner-link/filter [post]
func (controller *PartnerLinkController) Filter(ctx *gin.Context)  {
	fmt.Println(">>> PartnerLinkController - Filter <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req dto.ReqPartnerLinkDto

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


	services.InitPartnerLinkService(logs).Filter(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response PartnerLink Controller - Filter",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}

// @Summary PartnerLinkController - Save / Edit
// @Description PartnerLinkController Save/EDit
// @ID PartnerLinkController - Save/Edit
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqPartnerLinkDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/partner-Link [post]
func (controller *PartnerLinkController) Save(ctx *gin.Context)  {
	fmt.Println(">>> PartnerLinkController - Save <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req dto.ReqPartnerLinkDto

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


	services.InitPartnerLinkService(logs).Save(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response PartnerLink Controller - Save",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)



}


// @Summary PartnerLinkController - Delete
// @Description PartnerLinkController Delete
// @ID PartnerLinkController - Delete
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqPartnerLinkDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/partner-link/delete/ [post]
func (controller *PartnerLinkController) Delete(ctx *gin.Context)  {
	fmt.Println(">>> PartnerLinkController - Delete <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	id := ctx.Param("id")
	log.Println("req Id --> ", id)
	reqId, _ := strconv.Atoi(id)

	services.InitPartnerLinkService(logs).Delete(reqId, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response PartnerLink Controller - Delete",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)



}