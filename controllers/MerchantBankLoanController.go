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

type MerchantBankLoanController struct {

}

// @Summary MerchantBankLoan - Filter
// @Description MerchantBankLoan Filter Paging
// @ID MerchantBankLoanFilter
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqMerchantBankLoanDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response{contents=[]dbmodels.MerchantBankLoan}
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/merchant-bank-loan/filter [post]
func (controller *MerchantBankLoanController) Filter(ctx *gin.Context)  {
	fmt.Println(">>> MerchantBankLoanController - Filter <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	var res models.Response

	var req dto.ReqMerchantBankLoanDto

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


	services.InitMerchantBankLoanService().Filter(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response MerchantBankLoanController  - Filter",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}

// @Summary MerchantBankLoan - FindSubMerchantBankLoan
// @Description MerchantBankLoan FindSubMerchantBankLoan
// @ID MerchantBankLoanFindSubMerchantBankLoan
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqSubMerchantBankloanDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response{contents=[]dbmodels.SubMerchantBankLoan}
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/merchant-bank-loan/find-sub [post]
func (controller *MerchantBankLoanController) FindSubMerchantBankLoan(ctx *gin.Context)  {
	fmt.Println(">>> MerchantBankLoanController - FindSubMerchantBankLoan <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	var res models.Response

	var req dto.ReqSubMerchantBankloanDto

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


	services.InitMerchantBankLoanService().FindSubMerchantBankLoan(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response MerchantBankLoanController  - FindSubMerchantBankLoan",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}