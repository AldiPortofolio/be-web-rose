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

// MdrAggergatorTempService struct
type MdrAggergatorTempService struct {
	General 	models.GeneralModel
	MdrAggregatorTempRepository *db.MdrAggregatorTempRepository
	MdrAggregatorRepository *db.MdrAggregatorRepository
}

// InitMdrAggergatorTempService ...
func InitMdrAggergatorTempService(gen models.GeneralModel) *MdrAggergatorTempService {
	return &MdrAggergatorTempService{
		General: gen,
		MdrAggregatorTempRepository: db.InitMdrAggregatorTempRepository(),
		MdrAggregatorRepository: db.InitMdrAggregatorRepository(),
	}
}

// GetById ...
func (service *MdrAggergatorTempService) GetById(id int64) models.Response {
	fmt.Println(">>> MdrAggergatorTempService - GetById <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MdrAggergatorTempService: GetById",
		zap.Any("id", id))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MdrAggergatorTempService: GetById")
	defer span.Finish()

	var res models.Response

	data, err := service.MdrAggregatorTempRepository.GetMdrAggregatorTempByID(id)
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

// Filter ...
func (service *MdrAggergatorTempService) Filter(req dto.ReqFilterDto) models.Response {
	fmt.Println(">>> MdrAggergatorTempService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MdrAggergatorTempService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MdrAggergatorTempService: Filter")
	defer span.Finish()

	var res models.Response

	list, total, err := service.MdrAggregatorTempRepository.FilterMdrAggregatorTemp(req)
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

// Approve ...
func (service *MdrAggergatorTempService) Approve(req dto.ReqFilterDto) models.Response {
	fmt.Println(">>> MdrAggergatorTempService - Approve <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MdrAggergatorTempService: Approve",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MdrAggergatorTempService: Approve")
	defer span.Finish()

	var res models.Response

	temp, err := service.MdrAggregatorTempRepository.GetMdrAggregatorTempByID(req.Id)
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

	var data dbmodels.MdrAggregator
	data = service.PackMsgMdrAggregatorFromTemp(temp, data)
	data.ActionType = status_approval.APPROVE

	errSaveMdrAggregator := service.MdrAggregatorRepository.SaveMdrAggregator(&data)
	if errSaveMdrAggregator != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		res.Contents = data
		return res
	}

	temp.ActionType = status_approval.APPROVE
	errSaveMdrAggregatorTemp := service.MdrAggregatorTempRepository.SaveMdrAggregatorTemp(&temp)
	if errSaveMdrAggregatorTemp != nil {
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
func (service *MdrAggergatorTempService) Reject(req dto.ReqFilterDto) models.Response {
	fmt.Println(">>> MdrAggergatorTempService - Reject <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MdrAggergatorTempService: Reject",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MdrAggergatorTempService: Reject")
	defer span.Finish()

	var res models.Response

	temp, err := service.MdrAggregatorTempRepository.GetMdrAggregatorTempByID(req.Id)
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

	if temp.MdrAggregatorID != 0 {
		data, err := service.MdrAggregatorRepository.GetMdrAggregatorByID(temp.MdrAggregatorID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			res.Contents = temp
			return res
		}
		data.ActionType = status_approval.REJECT
		errSaveMdrAggregator := service.MdrAggregatorRepository.SaveMdrAggregator(&data)
		if errSaveMdrAggregator != nil {
			res.ErrCode = constants.EC_FAIL_SAVE
			res.ErrDesc = constants.EC_FAIL_SAVE_DESC
			res.Contents = data
			return res
		}
	}

	temp.ActionType = status_approval.REJECT
	errSaveMdrAggregatorTemp := service.MdrAggregatorTempRepository.SaveMdrAggregatorTemp(&temp)
	if errSaveMdrAggregatorTemp != nil {
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
func (service *MdrAggergatorTempService) ValidationApproval(actionType int) (models.Response) {
	var res models.Response
	if actionType == status_approval.APPROVE {
		res.ErrCode = constants.EC_FAIL_HAS_BEEN_APPROVED
		res.ErrDesc = constants.EC_FAIL_HAS_BEEN_APPROVED_DESC
		res.Contents = dbmodels.MdrAggregatorTemp{}
		return res
	}

	if actionType == status_approval.REJECT {
		res.ErrCode = constants.EC_FAIL_HAS_BEEN_REJECTED
		res.ErrDesc = constants.EC_FAIL_HAS_BEEN_REJECTED_DESC
		res.Contents = dbmodels.MdrAggregatorTemp{}
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	return res
}

// PackMsgMdrAggregatorFromTemp ...
func (service *MdrAggergatorTempService) PackMsgMdrAggregatorFromTemp(temp dbmodels.MdrAggregatorTemp, data dbmodels.MdrAggregator) dbmodels.MdrAggregator {
	return dbmodels.MdrAggregator{
		ID: temp.MdrAggregatorID,
		LatestSuggestion: temp.LatestSuggestion,
		LatestSuggestor: temp.LatestSuggestor,
		LatestApprover: auth.UserLogin.Name,
		LatestApproval: time.Now(),
		MidMerchant: temp.MidMerchant,
		MidPartner: temp.MidPartner,
		MdrType: temp.MdrType,
		TransactionType: temp.TransactionType,
		MerchantCategory: temp.MerchantCategory,
		GroupPartner: temp.GroupPartner,
		Mdr: temp.Mdr,
		Notes: temp.Notes,
		Status: temp.Status,
	}
}