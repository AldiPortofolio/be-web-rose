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

type AkuisisiSfaFailedController struct {

}

// @Summary Akuisisi SFA Failed- Filter
// @Description Akuisisi SFA Failed Filter
// @ID Akuisisi SFA Failed - Filter
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqAkuisisiSfaFailed true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response{contents=[]dbmodels.AkuisisiSfaFailed}
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/akuisisi-sfa-failed/filter [post]
func (controller *AkuisisiSfaFailedController) GetFilterPaging(ctx *gin.Context) {
	fmt.Println(">>> AkuisisiSfaController - GetFilterPaging <<<")

	var req dto.ReqAkuisisiSfaFailed
	var res models.Response

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind request:", err)
		res.ErrCode = constants.ERR_UNMARSHAL
		res.ErrDesc = constants.ERR_UNMARSHAL_MSG
		ctx.JSON(http.StatusOK, res)
		return
	}
	reqByte,_ := json.Marshal(req)
	log.Println("req --> ", string(reqByte))

	res = services.InitAkuisisiSfaFailedService().Filter(req)


	ctx.JSON(http.StatusOK, res)

}