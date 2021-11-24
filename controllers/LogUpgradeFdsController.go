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
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"rose-be-go/services"
)

type LogUpgradeFdsController struct {

}



// @Summary Upgrade FDS - Filter
// @Description Upgrade FDS Filter Paging
// @ID Upgrade FDS - Filter
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqLogUpgradeFdsDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response{contents=[][]dbmodels.LogUpgradeFds}
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/upgrade-fds/filter [post]
func (controller *LogUpgradeFdsController) Filter(ctx *gin.Context)  {
	fmt.Println(">>> LogUpgradeFdsController - Filter <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req dto.ReqLogUpgradeFdsDto

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


	services.InitLogUpgradeFdsService(logs).Filter(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response LogUpgradeFdsController - Filter",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}

// @Summary  Upgrade FDS - Save / Edit
// @Description  Upgrade FDS Save/EDit
// @ID  Upgrade FDS - Save/Edit
// @Param Authorization header string true "Bearer"
// @Param body body dbmodels.LogUpgradeFds true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/upgrade-fds [post]
func (controller *LogUpgradeFdsController) Save(ctx *gin.Context)  {
	fmt.Println(">>> LogUpgradeFdsController - Save <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req dbmodels.LogUpgradeFds

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


	services.InitLogUpgradeFdsService(logs).Save(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response LogUpgradeFdsController - Save",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)



}