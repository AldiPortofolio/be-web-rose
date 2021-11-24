package services

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"log"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/constants/status_approval"
	"rose-be-go/constants/mdr_type"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"time"
)

// MdrAggregatorService struct 
type MdrAggregatorService struct {
	General 	models.GeneralModel
	MdrAggregatorRepository *db.MdrAggregatorRepository
}

// InitMdrAggregatorService ...
func InitMdrAggregatorService(gen models.GeneralModel) *MdrAggregatorService {
	return &MdrAggregatorService{
		General: gen,
		MdrAggregatorRepository: db.InitMdrAggregatorRepository(),
	}
}

// SaveTemp ...
func (service *MdrAggregatorService) SaveTemp(req dto.ReqMdrAggragtorDto) models.Response {
	fmt.Println(">>> MdrAggregatorService - SaveTemp <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MdrAggregatorService: SaveTemp",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MdrAggregatorService: SaveTemp")
	defer span.Finish()

	var res models.Response

	if req.ActionType == status_approval.EDIT || req.ActionType == status_approval.DELETE {
		cek := service.MdrAggregatorRepository.CheckQueue(req.MdrAggregatorID)
		log.Println("cek queue --> ", cek)

		if !cek {
			res.ErrCode = constants.EC_FAIL_PENDING_APPROVAL
			res.ErrDesc = constants.EC_FAIL_PENDING_APPROVAL_DESC
			return res
		}
	}

	data:= dbmodels.MdrAggregatorTemp{
		MdrAggregatorID:  req.MdrAggregatorID,
		MidMerchant:      req.MidMerchant,
		MidPartner:       req.MidPartner,
		GroupPartner:     req.GroupPartner,
		ActionType:       req.ActionType,
		Status:           req.Status,
		Notes:            req.Notes,
		Mdr:              req.Mdr,
		MdrType:          mdr_type.QR,
		MerchantCategory: req.MerchantCategory,
		TransactionType:  req.TransactionType,
		LatestSuggestion: time.Now(),
		LatestSuggestor:  auth.UserLogin.Name,
	}

	err:= service.MdrAggregatorRepository.SaveMdrAggregatorTemp(data)
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

// Filter ...
func (service *MdrAggregatorService) Filter(req dto.ReqMdrAggragtorDto) models.Response {
	fmt.Println(">>> MdrAggregatorService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MdrAggregatorService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MdrAggregatorService: Filter")
	defer span.Finish()

	var res models.Response


	list, total, err := service.MdrAggregatorRepository.FilterMdrAggregator(req)

	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.TotalData = total
		res.Contents = list
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.TotalData = total
	res.Contents = list

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}

// GetById ...
func (service *MdrAggregatorService) GetById(id int64) models.Response {
	fmt.Println(">>> MdrAggregatorService - GetById <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MdrAggregatorService: GetById",
		zap.Any("id", id))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MdrAggregatorService: GetById")
	defer span.Finish()

	var res models.Response

	data, err := service.MdrAggregatorRepository.GetMdrAggregatorByID(id)
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
