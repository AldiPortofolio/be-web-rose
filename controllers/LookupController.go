package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rose-be-go/constants"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	"rose-be-go/services"
)

type LookupController struct {

}


// @Summary Lookup - Filter
// @Description Lookup Filter Paging
// @ID LookupFilter
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqLookupDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response{contents=[]dbmodels.Lookup}
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/lookup/filter [post]
func (controller *LookupController) Filter(ctx *gin.Context)  {
	fmt.Println(">>> LookupController - Filter <<<")
	// initiate logs

	var res models.Response

	var req dto.ReqLookupDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.ErrCode = constants.ERR_UNMARSHAL
		res.ErrDesc = constants.ERR_UNMARSHAL_MSG
		fmt.Println("err unmarshal "+err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	reqByte,_ := json.Marshal(req)
	log.Println("req --> ", string(reqByte))


	services.InitLookupService().Filter(req, &res)

	bodyRes, _ := json.Marshal(res)

	fmt.Println("Response LookupController - Filter",
		"ResponseBody: ", string(bodyRes))


	ctx.JSON(http.StatusOK, res)

}


// @Summary Lookup - Save / Edit
// @Description Lookup Save/EDit
// @ID LookupSave/Edit
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqLookupDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/lookup [post]
func (controller *LookupController) Save(ctx *gin.Context)  {
	fmt.Println(">>> LookupController - Save <<<")
	// initiate logs

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req dto.ReqLookupDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.ErrCode = constants.ERR_UNMARSHAL
		res.ErrDesc = constants.ERR_UNMARSHAL_MSG
		fmt.Println("err unmarshal "+err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	reqByte,_ := json.Marshal(req)
	log.Println("req --> ", string(reqByte))


	services.InitLookupService().Save(req, &res)

	bodyRes, _ := json.Marshal(res)

	fmt.Println("Response LookupGroupController - Save", "ResponseBody: ", string(bodyRes))


	ctx.JSON(http.StatusOK, res)



}