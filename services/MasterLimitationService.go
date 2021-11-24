package services

import (
	"encoding/json"
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

// MasterLimitationService struct 
type MasterLimitationService struct {
	General 	models.GeneralModel
	MasterLimitationRepository *db.MasterLimitationRepository
}

// InitMasterLimitationService ...
func InitMasterLimitationService(gen models.GeneralModel) *MasterLimitationService {
	return &MasterLimitationService{
		General: gen,
		MasterLimitationRepository: db.InitMasterLimitationRepository(),
	}
}

// Filter ...
func (service *MasterLimitationService) Filter(req dto.MasterLimitationReq) models.Response {
	fmt.Println(">>> LimitationMerchantService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("LimitationMerchantService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "LimitationMerchantService: Filter")
	defer span.Finish()

	var res models.Response


	list, total, err := service.MasterLimitationRepository.FilterLimitationMerchant(req)

	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.TotalData = total
		res.Contents = list
		return res
	}

	for i:=0;i < len(list) ;i++  {
		var merchantGroup []string
		//dataGroup, _ := json.Marshal(list[i].ByGroup)
		json.Unmarshal([]byte(list[i].ByGroup), &merchantGroup)
		fmt.Println("merchantGroup-->", merchantGroup)
		merchant, _ := service.MasterLimitationRepository.GetGroupMerchant(merchantGroup)
		list[i].MerchantGroup = merchant
	}


	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.TotalData = total
	res.Contents = list

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}

// SaveTemp ...
func (service *MasterLimitationService) SaveTemp(req dto.MasterLimitationReq) models.Response {
	fmt.Println(">>> LimitationMerchantService - SaveTemp <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("LimitationMerchantService: SaveTemp",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "LimitationMerchantService: SaveTemp")
	defer span.Finish()

	var res models.Response
	var data dbmodels.LimitationMerchantTemp


	//dataByte, _ := json.Marshal(req)
	//json.Unmarshal(dataByte, data)

	if req.ActionType == status_approval.EDIT || req.ActionType == status_approval.DELETE {
		cek := service.MasterLimitationRepository.CheckQueue(req.MasterLimitationId)

		if !cek {
			res.ErrCode = constants.EC_FAIL_PENDING_APPROVAL
			res.ErrDesc = constants.EC_FAIL_PENDING_APPROVAL_DESC
			return res
		}
	}

	//dataGroup, _ := json.Marshal(req.ByGroup)
	//dataByGroup := fmt.Sprintf(string(dataGroup))
	data.ByGroup = req.ByGroup
	data.LatestSuggestion = time.Now()
	data.LatestSuggestor = auth.UserLogin.Name
	data.ActionType = req.ActionType
	data.ProductType = req.ProductType
	data.ProductName = req.ProductName
	data.LimitFreq = req.LimitFreq
	data.LimitAmt = req.LimitAmt
	data.LimitFreqMin = req.LimitFreqMin
	data.LimitAmtMin = req.LimitAmtMin
	data.ByTime = req.ByTime
	data.MasterLimitationId = req.MasterLimitationId
	data.Status = req.Status

	err := service.MasterLimitationRepository.SaveLimitationTemp(data)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}

// GetById ...
func (service *MasterLimitationService) GetById(id int64) models.Response {
	fmt.Println(">>> MasterLimitationService - GetById <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MasterLimitationService: GetById",
		zap.Any("id", id))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MasterLimitationService: GetById")
	defer span.Finish()

	var res models.Response

	data, err := service.MasterLimitationRepository.GetMasterLimitationByID(id)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.Contents = data
		return res
	}


	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.Contents = data

	sugarLogger.Info("Response", zap.Any("res", res))

	return res

}
