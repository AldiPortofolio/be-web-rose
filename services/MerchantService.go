package services

import (
	"encoding/json"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"log"
	ottoutils "ottodigital.id/library/utils"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/kafka"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	"time"
)

type MerchantService struct {
	General models.GeneralModel
	SendKafka         	func(req kafka.PublishReq) ([]byte, error)
	MerchantRepository	*db.MerchantRepository
}

func InitMerchantService(gen models.GeneralModel) *MerchantService  {
	return &MerchantService{
		General: gen,
		SendKafka: kafka.SendPublishKafka,
		MerchantRepository: db.InitMerchantRepository(),
	}
}

func (service *MerchantService) Upgrade (req dto.ReqUpgradeMerchantFdsDto) models.Response {
	fmt.Println(">>> MerchantService - Upgrade <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MerchantService: Upgrade",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "ReportFinishService: Send")
	defer span.Finish()

	var res models.Response = service.PushToKafka(req)


	return res
}

func (service *MerchantService) PushToKafka(req dto.ReqUpgradeMerchantFdsDto) models.Response  {
	topic := ottoutils.GetEnv("ROSE_BE_GO_UPGRADE_MERCHANT_TOPIC", "rose-worker-upgrade-fds-topic")

	var res models.Response
	reqByte,_ := json.Marshal(req)

	kafkaReq := kafka.PublishReq{
		Topic: topic,
		Bytes: reqByte,
		Timestamp: time.Now().Format("2006-01-02"),
	}

	kafkaRes, err := service.SendKafka(kafkaReq)
	if err != nil {
		res.ErrCode = constants.ERR_CODE_04
		res.ErrDesc = constants.ERR_CODE_04_MSG
		return res
	}
	log.Println("kafkaRes--> ", string(kafkaRes))


	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	return res
}

func (service *MerchantService) Filter(req dto.ReqMerchantDto) models.Response {
	fmt.Println(">>> MerchantService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MerchantService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MerchantService: Filter")
	defer span.Finish()

	var res models.Response
	list, total, err := service.MerchantRepository.Filter(req)
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


//GetListByGroupId...
func (service *MerchantService) GetListByGroupId(req dto.ReqMerchantGroupDto) models.Response {
	fmt.Println(">>> MerchantService - GetListByGroupId <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MerchantService: GetListByGroupId",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MerchantService: GetListByGroupId")
	defer span.Finish()

	var res models.Response
	list, total, err := service.MerchantRepository.GetListByGroupId(req)
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

func (service *MerchantService)FindTanggalAkuisisi(req models.ReqPhoneNumber, res *models.Response)  {
	tanggalAkuisisi,_ := service.MerchantRepository.FindTanggalAkuisisi(req.StorePhoneNumber)


	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.Contents = tanggalAkuisisi


}

func (service *MerchantService) FilterDashboard(req dto.ReqDashboardMerchantDto) models.Response {
	fmt.Println(">>> MerchantService - FilterDashboard <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("MerchantService: FilterDashboard",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MerchantService: FilterDashboard")
	defer span.Finish()

	var res models.Response
	list, total, err := service.MerchantRepository.FilterDashboard(req)
	if err != nil {
		log.Println("err : ", err)
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
