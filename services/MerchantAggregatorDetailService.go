package services

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"log"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/constants/status_approval"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"time"
)

type MerchantAggregatorDetailService struct {
	General 	models.GeneralModel
}

func InitMerchantAggregatorDetailService(gen models.GeneralModel) *MerchantAggregatorDetailService  {
	return &MerchantAggregatorDetailService{
		General: gen,
	}
}


func (service *MerchantAggregatorDetailService) Filter(req dto.ReqFilterDto) models.Response {
	fmt.Println(">>> MerchantAggregatorDetailService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MerchantAggregatorDetailService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MerchantAggregatorDetailService: Filter")
	defer span.Finish()

	var res models.Response

	//list, total, err := db.FilterDataAggregatorDetail(req)
	list, total, err := db.InitMerchantAggregatorDetailRepository().FilterMerchantAggList(req)
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

func (service *MerchantAggregatorDetailService) FilterTemp(req dto.ReqFilterDto) models.Response {
	fmt.Println(">>> MerchantAggregatorDetailService - FilterTemp <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MerchantAggregatorDetailService: FilterTemp",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MerchantAggregatorDetailService: FilterTemp")
	defer span.Finish()

	var res models.Response

	list, total, err := db.InitMerchantAggregatorDetailRepository().FilterMerchantAggTempList(req)
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

func (service *MerchantAggregatorDetailService) ListMerchantAggregator(req dto.ReqFilterDto) models.Response {
	fmt.Println(">>> MerchantAggregatorDetailService - ListMerchantAggregator <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MerchantAggregatorDetailService: ListMerchantAggregator",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MerchantAggregatorDetailService: ListMerchantAggregator")
	defer span.Finish()

	var res models.Response

	list, err := db.InitMerchantAggregatorDetailRepository().FilterListAggregatorDetail(req)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.Contents = list
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.Contents = list

	sugarLogger.Info("Response", zap.Any("res", res))

	return res

}



func (service *MerchantAggregatorDetailService) SaveTemp(req dto.ReqFilterDto) models.Response {
	fmt.Println(">>> MerchantAggregatorDetailService - SaveTemp <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MerchantAggregatorDetailService: SaveTemp",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MerchantAggregatorDetailService: SaveTemp")
	defer span.Finish()

	var res models.Response

	// validation merchant_aggregator_detail on approva;
	listMAGDetail,total,_ := db.InitMerchantAggregatorDetailRepository().FindMerchantAggregatorDetailTempByMidAggregator(req)
	log.Println(len(listMAGDetail))
	if total > 0 {
		res.ErrCode = constants.EC_FAIL_PENDING_APPROVAL
		res.ErrDesc = constants.EC_FAIL_PENDING_APPROVAL_DESC
		return res
	}


	// bulk insert to aggregator details
	for i:=0; i<len(req.MidMerchant); i++ {
		magdTemp := service.PackMsgTemp(req, req.MidMerchant[i])
		errSave := db.InitMerchantAggregatorDetailRepository().SaveMerchantAggregatorDetailTemp(&magdTemp)
		if errSave != nil {
			log.Println("err save --> ", errSave.Error())
		}
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}


func (service *MerchantAggregatorDetailService) FilterDetailDataApproval(req dto.ReqFilterDto) models.Response {
	fmt.Println(">>> MerchantAggregatorDetailService - FilterDataApproval <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MerchantAggregatorDetailService: FilterDataApproval",
		zap.Any("id", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MerchantAggregatorDetailService: FilterDataApproval")
	defer span.Finish()

	var res models.Response

	data, total, err := db.InitMerchantAggregatorDetailRepository().FindMerchantApprovalByMidAggregator(req)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.Data = data
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.Contents = data
	res.TotalData = total
	sugarLogger.Info("Response", zap.Any("res", res))

	return res

}

func (service *MerchantAggregatorDetailService) ListDataApproval(req dto.ReqFilterDto) models.Response {
	fmt.Println(">>> MerchantAggregatorDetailService - ListDataApproval <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MerchantAggregatorDetailService: ListDataApproval",
		zap.Any("id", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MerchantAggregatorDetailService: ListDataApproval")
	defer span.Finish()

	var res models.Response

	data, total, err := db.InitMerchantAggregatorDetailRepository().FindListMerchantAggregatorApproval(req)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.Data = data
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.Contents = data
	res.TotalData = total
	sugarLogger.Info("Response", zap.Any("res", res))

	return res

}

func (service *MerchantAggregatorDetailService) Approve(req dto.ReqFilterDto) models.Response {
	fmt.Println(">>> MerchantAggregatorDetailService - Approve <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MerchantAggregatorDetailService: Approve",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MerchantAggregatorDetailService: Approve")
	defer span.Finish()

	var res models.Response

	temps, _, err := db.InitMerchantAggregatorDetailRepository().FindMerchantAggregatorDetailTempByMidAggregator(req)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		return res
	}

	for i:=0; i<len(temps); i++ {
		var detail dbmodels.MerchantAggregatorDetail

		if temps[i].ActionType == 4 {
			req.MidFilter = temps[i].MidMerchant
			data, _, _ := db.InitMerchantAggregatorDetailRepository().FilterMerchantAggList(req)
			detail = data[0]
		}

		if errApprove := service.ApproveAggregator(temps[i], &detail); errApprove != nil {
			log.Println("err approve -->", errApprove)
		}
	}



	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	sugarLogger.Info("Response", zap.Any("res", res))

	return res

}

/*
	Approve Aggregator
	- edit action type on temp to approve
	- update data aggregator_detail from temp
*/
func (service *MerchantAggregatorDetailService) ApproveAggregator(temp dbmodels.MerchantAggregatorDetailTemp, detail *dbmodels.MerchantAggregatorDetail) error{

	service.PackMsgDetailFromTemp(temp, detail)

	fmt.Println("action Type--->", temp.ActionType)

		temp.ActionType = status_approval.APPROVE

		if errTemp := db.InitMerchantAggregatorDetailRepository().SaveMerchantAggregatorDetailTemp(&temp); errTemp!=nil {
			return errTemp
		}

		//detail.Status = 1

		if errDetail := db.InitMerchantAggregatorDetailRepository().SaveMerchantAggregatorDetail(detail); errDetail!=nil {
			return errDetail
		}

	return nil
}


/*
	Pack message detail from data temp
*/
func (service *MerchantAggregatorDetailService) PackMsgDetailFromTemp(temp dbmodels.MerchantAggregatorDetailTemp, detail *dbmodels.MerchantAggregatorDetail) {

	var status int64

	switch temp.ActionType {
	case 0 : status = status_approval.STATUS_APPROVE
	case 4 : status = status_approval.STATUS_REJECTED
	}

	detail.MidMerchant = temp.MidMerchant
	detail.MidAggregator = temp.MidAggregator
	detail.LatestSuggestor = temp.LatestSuggestor
	detail.LatestSuggestion = temp.LatestSuggestion
	detail.LatestApproval = time.Now()
	detail.LatestApprover = auth.UserLogin.Name
	detail.Status = status

}


func (service *MerchantAggregatorDetailService) PackMsgTemp(req dto.ReqFilterDto, midMerchant string) dbmodels.MerchantAggregatorDetailTemp {

	var actionType int

	switch req.Action {
	case 0 : actionType = status_approval.CREATE
	case 4 : actionType = status_approval.DELETE
	}

	return dbmodels.MerchantAggregatorDetailTemp{
		MidMerchant: midMerchant,
		MidAggregator: req.MidAggregator,
		ActionType: actionType,
		LatestSuggestion: time.Now(),
		LatestSuggestor: auth.UserLogin.Name,
	}
}

