package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	"rose-be-go/services"
	"rose-be-go/constants"
)

type MerchantBankAccountController struct {

}




// @Summary MerchantBankAccount - Save / Edit
// @Description MerchantBankAccount Save/EDit
// @ID MerchantBankAccount - Save/Edit
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqMerchantBankAccountDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/merchant-bank-account [post]
func (controller *MerchantBankAccountController) Save(ctx *gin.Context)  {
	fmt.Println(">>> MerchantBankAccountController - Save <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req dto.ReqMerchantBankAccountDto

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


	services.InitMerchantBankAccountService(logs).Save(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response MerchantBankAccount Controller - Filter",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}


// @Summary MerchantBankAccount - get by mid
// @Description Get MerchantBankAccount by mid
// @ID MerchantBankAccount - FindAllByMid
// @Param Authorization header string true "Bearer"
// @Param mid path string true "mid"
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} models.Response{contents=[]dbmodels.MerchantBankAccount}
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/merchant-bank-account/mid/{mid} [get]
func (controller *MerchantBankAccountController) FindAllByMid(ctx *gin.Context)  {
	fmt.Println(">>> MerchantBankAccountController - FindAllByMid <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	mid:= ctx.Param("mid")


	services.InitMerchantBankAccountService(logs).FindAllAccount(mid, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response MerchantBankAccountController - FindAllByMid",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)


}

// @Summary MerchantBankAccount - Get Data pending MerchantBankAccount
// @Description MerchantBankAccount filter approval
// @ID MerchantBankAccount - filter approval
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqMerchantBankAccountDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response{contents=[]dbmodels.MerchantBankAccount}
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/merchant-bank-account/approval [post]
func (controller *MerchantBankAccountController) FilterApproval(ctx *gin.Context)  {
	fmt.Println(">>> MerchantBankAccountController - FilterApproval <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req dto.ReqMerchantBankAccountDto

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


	services.InitMerchantBankAccountService(logs).FilterApproval(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response MerchantBankAccount Controller - FilterApproval",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}

// @Summary MerchantBankAccount - Approve
// @Description MerchantBankAccount Approve
// @ID MerchantBankAccount - Approve
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqApprovalMerchantBankAccountDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/merchant-bank-account/approval/approve [post]
func (controller *MerchantBankAccountController) Approve(ctx *gin.Context)  {
	fmt.Println(">>> MerchantBankAccountController - Approve <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req dto.ReqApprovalMerchantBankAccountDto

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


	services.InitMerchantBankAccountService(logs).Approve(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response MerchantBankAccount Controller - Approve",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}

// @Summary MerchantBankAccount - Reject
// @Description MerchantBankAccount Reject
// @ID MerchantBankAccount - Reject
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqApprovalMerchantBankAccountDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/merchant-bank-account/approval/reject [post]
func (controller *MerchantBankAccountController) Reject(ctx *gin.Context)  {
	fmt.Println(">>> MerchantBankAccountController - Reject <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req dto.ReqApprovalMerchantBankAccountDto

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


	services.InitMerchantBankAccountService(logs).Reject(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response MerchantBankAccount Controller - Reject",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}

func (controller *MerchantBankAccountController) ResendPushNotif(ctx *gin.Context)  {
	fmt.Println(">>> MerchantBankAccountController - ResendPushNotif <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req dto.ReqApprovalMerchantBankAccountDto

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


	services.InitMerchantBankAccountService(logs).ResendPushNotif(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response MerchantBankAccount Controller - ResendPushNotif",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}


func (controller *MerchantBankAccountController) ValidationBankAccount(ctx *gin.Context)  {
	fmt.Println(">>> MerchantBankAccountController - ValidationBankAccount <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	res := models.Response{
		ErrCode: constants.ERR_EMPTY_INPUT,
		ErrDesc: constants.ERR_EMPTY_INPUT_MSG,
	}

	var req dto.ReqValidationBankAccount

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


	services.InitMerchantBankAccountService(logs).ValidationBankAccount(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response MerchantBankAccount Controller - ValidationBankAccount",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}