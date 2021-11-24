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

type AkuisisiSfaController struct {

}

// @Summary Akuisisi SFA - Filter
// @Description Akuisisi SFA Filter
// @ID Akuisisi SFA - Filter
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqAkuisisiSfa true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response{contents=[]dbmodels.AkuisisiSfa}
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/akuisisi-sfa/filter [post]
func (controller *AkuisisiSfaController) GetFilterPaging(ctx *gin.Context) {
	fmt.Println(">>> AkuisisiSfaController - GetFilterPaging <<<")

	var req dto.ReqAkuisisiSfa
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

	res = services.InitAkuisisiSfaService().Filter(req)


	ctx.JSON(http.StatusOK, res)

}