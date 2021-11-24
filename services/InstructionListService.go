package services

import (
	"fmt"
	"log"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// InstructionList struct
type InstructionListService struct {
	General                   models.GeneralModel
	InstructionListRepository *db.InstructionListRepository
}

// InitInstructionList ...
func InitInstructionListService(gen models.GeneralModel) *InstructionListService {
	return &InstructionListService{
		General:          gen,
		InstructionListRepository: db.InitInstructionListRepository(),
	}
}

// Filter ...
func (service *InstructionListService) Filter(req dto.ReqInstructionListDto) models.Response {
	fmt.Println(">>> InstructionList - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("InstructionList: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "InstructionList: Filter")
	defer span.Finish()

	var res models.Response
	list, total, err := service.InstructionListRepository.Filter(req)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.TotalData = total
		res.Contents = list
		return res
	}

	log.Println("total -->", total)

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.TotalData = total
	res.Contents = list

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}

// Save ...
func (service *InstructionListService) Save(req dto.ReqInstructionListDto) models.Response {
	fmt.Println(">>> InstructionList - save <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("InstructionList: save",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "InstructionList: save")
	defer span.Finish()

	var res models.Response

	var data dbmodels.InstructionList
	var err error

	if req.ID > 0 {
		data, err = service.InstructionListRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return res
		}
	}

	data = dbmodels.InstructionList{
		ID:              req.ID,
		Title: 			 req.Title,
		Logo: 			 req.Logo,
		Description:     req.Description,
		Sequence:        req.Sequence,
		
	}

	if err := service.InstructionListRepository.Save(&data); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}

// Save ...
func (service *InstructionListService) Delete(id int) models.Response {
	fmt.Println(">>> InstructionList - Delete <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("InstructionList: save",
		zap.Any("req", id))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "InstructionList: delete")
	defer span.Finish()

	var res models.Response

	if err := service.InstructionListRepository.Delete(id); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}
