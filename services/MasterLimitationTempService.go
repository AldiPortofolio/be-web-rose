package services

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/constants/status_approval"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"time"
)

// MasterLimitationTempService struct 
type MasterLimitationTempService struct {
	General 	models.GeneralModel
	MasterLimitationRepository *db.MasterLimitationRepository
	MasterLimitationTempRepository *db.MasterLimitationTempRepository
}

// InitMasterLimitationTempService ...
func InitMasterLimitationTempService(gen models.GeneralModel) *MasterLimitationTempService {
	return &MasterLimitationTempService{
		General: gen,
		MasterLimitationRepository: db.InitMasterLimitationRepository(),
		MasterLimitationTempRepository: db.InitMasterLimitationTempRepository(),
	}
}

// GetById ...
func (service *MasterLimitationTempService) GetById(id int64) models.Response {
	fmt.Println(">>> MasterLimitationTempService - GetById <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MasterLimitationTempService: GetById",
		zap.Any("id", id))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MasterLimitationTempService: GetById")
	defer span.Finish()

	var res models.Response

	data, err := service.MasterLimitationTempRepository.GetMasterLimitationTempByID(id)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.Contents = data
		return res
	}

	switch data.ActionType {
	case status_approval.CREATE:
		data.ActionTypeDesc = status_approval.CREATE_DESC
		
	case status_approval.EDIT:
		data.ActionTypeDesc = status_approval.EDIT_DESC
		
	case status_approval.REJECT:
		data.ActionTypeDesc = status_approval.REJECT_DESC
		
	case status_approval.APPROVE:
		data.ActionTypeDesc = status_approval.APPROVE_DESC
		
	case status_approval.DELETE:
		data.ActionTypeDesc = status_approval.DELETE_DESC
		

	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.Contents = data

	sugarLogger.Info("Response", zap.Any("res", res))

	return res

}

// Approve ...
func (service *MasterLimitationTempService) Approve(req dto.ReqFilterDto) models.Response {
	fmt.Println(">>> MasterLimitationTempService - Approve <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MasterLimitationTempService: Approve",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MasterLimitationTempService: Approve")
	defer span.Finish()

	var res models.Response

	temp, err := service.MasterLimitationTempRepository.GetMasterLimitationTempByID(req.Id)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.Contents = temp
		return res
	}

	res = service.ValidationApproval(temp.ActionType)
	if res.ErrCode != constants.ERR_SUCCESS {
		return res
	}

	var data dbmodels.LimitationMerchant
	data = service.PackMsgMasterLimitationFromTemp(temp, data)
	data.ActionType = status_approval.APPROVE

	errSaveMasterLimitation := service.MasterLimitationRepository.SaveMasterLimitation(&data)
	if errSaveMasterLimitation != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		res.Contents = data
		return res
	}

	temp.ActionType = status_approval.APPROVE
	errSaveMasterLimitationTemp := service.MasterLimitationTempRepository.SaveMasterLimitationTemp(&temp)
	if errSaveMasterLimitationTemp != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		res.Contents = data
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.Contents = data

	sugarLogger.Info("Response", zap.Any("res", res))

	return res

}

// Reject ...
func (service *MasterLimitationTempService) Reject(req dto.ReqFilterDto) models.Response {
	fmt.Println(">>> MasterLimitationTempService - Reject <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MasterLimitationTempService: Reject",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MasterLimitationTempService: Reject")
	defer span.Finish()

	var res models.Response

	temp, err := service.MasterLimitationTempRepository.GetMasterLimitationTempByID(req.Id)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.Contents = temp
		return res
	}

	res = service.ValidationApproval(temp.ActionType)
	if res.ErrCode != constants.ERR_SUCCESS {
		return res
	}

	if temp.MasterLimitationId != 0 {
		data, err := service.MasterLimitationRepository.GetMasterLimitationByID(temp.MasterLimitationId)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			res.Contents = temp
			return res
		}
		data.ActionType = status_approval.REJECT
		errSaveMasterLimitation := service.MasterLimitationRepository.SaveMasterLimitation(&data)
		if errSaveMasterLimitation != nil {
			res.ErrCode = constants.EC_FAIL_SAVE
			res.ErrDesc = constants.EC_FAIL_SAVE_DESC
			res.Contents = data
			return res
		}
	}

	temp.ActionType = status_approval.REJECT
	errSaveMasterLimitationTemp := service.MasterLimitationTempRepository.SaveMasterLimitationTemp(&temp)
	if errSaveMasterLimitationTemp != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		res.Contents = temp
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.Contents = temp

	sugarLogger.Info("Response", zap.Any("res", res))

	return res

}

// ValidationApproval ...
func (service *MasterLimitationTempService) ValidationApproval(actionType int) (models.Response) {
	var res models.Response
	if actionType == status_approval.APPROVE {
		res.ErrCode = constants.EC_FAIL_HAS_BEEN_APPROVED
		res.ErrDesc = constants.EC_FAIL_HAS_BEEN_APPROVED_DESC
		res.Contents = dbmodels.LimitationMerchantTemp{}
		return res
	}

	if actionType == status_approval.REJECT {
		res.ErrCode = constants.EC_FAIL_HAS_BEEN_REJECTED
		res.ErrDesc = constants.EC_FAIL_HAS_BEEN_REJECTED_DESC
		res.Contents = dbmodels.LimitationMerchantTemp{}
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	return res
}

// PackMsgMasterLimitationFromTemp ...
func (service *MasterLimitationTempService) PackMsgMasterLimitationFromTemp(temp dbmodels.LimitationMerchantTemp, data dbmodels.LimitationMerchant) dbmodels.LimitationMerchant {
	return dbmodels.LimitationMerchant{
		ID: temp.MasterLimitationId,
		LatestSuggestion: temp.LatestSuggestion,
		LatestSuggestor: temp.LatestSuggestor,
		LatestApprover: auth.UserLogin.Name,
		LatestApproval: time.Now(),
		ByGroup: temp.ByGroup,
		ByTime: temp.ByTime,
		LimitAmt: temp.LimitAmt,
		LimitAmtMin: temp.LimitAmtMin,
		LimitFreq: temp.LimitFreq,
		LimitFreqMin: temp.LimitFreqMin,
		ProductName: temp.ProductName,
		ProductType: temp.ProductType,
		Status: temp.Status,
	}
}

// Filter ...
func (service *MasterLimitationTempService) Filter(req dto.ReqFilterDto) models.Response {
	fmt.Println(">>> MasterLimitationTempService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MasterLimitationTempService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MasterLimitationTempService: Filter")
	defer span.Finish()

	var res models.Response

	list, total, err := service.MasterLimitationTempRepository.FilterMasterLimitationTemp(req)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.TotalData = total
		res.Contents = list
		return res
	}

	for i:=0; i<len(list); i++ {
		switch list[i].ActionType {
		case status_approval.CREATE:
			list[i].ActionTypeDesc = status_approval.CREATE_DESC
			
		case status_approval.EDIT:
			list[i].ActionTypeDesc = status_approval.EDIT_DESC
			
		case status_approval.REJECT:
			list[i].ActionTypeDesc = status_approval.REJECT_DESC
			
		case status_approval.APPROVE:
			list[i].ActionTypeDesc = status_approval.APPROVE_DESC
			
		case status_approval.DELETE:
			list[i].ActionTypeDesc = status_approval.DELETE_DESC
			

		}

	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.TotalData = total
	res.Contents = list

	sugarLogger.Info("Response", zap.Any("res", res))

	return res

}