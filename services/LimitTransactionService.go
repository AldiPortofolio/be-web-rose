package services

import (
	"encoding/json"
	"fmt"
	"log"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"rose-be-go/redis/redis_cluster"
	"time"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// LimitTransactionService struct
type LimitTransactionService struct {
	General 					models.GeneralModel
	LimitTransactionRepository 	*db.LimitTransactionRepository
	SaveRedis 					func(string, interface{}) (error)
}

// InitLimitTransactionService ...
func InitLimitTransactionService(gen models.GeneralModel) *LimitTransactionService {
	return &LimitTransactionService{
		General:gen,
		LimitTransactionRepository: db.InitLimitTransactionRepository(),
		SaveRedis: redis_cluster.SaveRedis,
	}
}

// Filter ...
func (service *LimitTransactionService) Filter(req dto.ReqLimitTransactionDto) models.Response {
	fmt.Println(">>> LimitTransactionService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("LimitTransactionService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "LimitTransactionService: Filter")
	defer span.Finish()

	var res models.Response
	list, total, err := service.LimitTransactionRepository.Filter(req)
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
func (service *LimitTransactionService) Save(req dto.ReqLimitTransactionDto) models.Response {
	fmt.Println(">>> LimitTransactionService - save <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("LimitTransactionService: save",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "LimitTransactionService: save")
	defer span.Finish()

	var res models.Response

	var data dbmodels.LimitTransaction
	var err error


	data.CreatedAt = time.Now()
	data.CreatedBy = auth.UserLogin.Name

	if req.ID > 0 {
		data, err = service.LimitTransactionRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return res
		}
	}
	data.UserCategory = req.UserCategory
	data.LevelMerchant = req.LevelMerchant
	data.LimitAmount = req.LimitAmount
	data.MinLimitAmount = req.MinLimitAmount
	data.LimitFreq = req.LimitFreq
	data.TimeFrame = req.TimeFrame
	data.FeatureProduct = req.FeatureProduct
	data.UpdatedAt = time.Now()
	data.UpdatedBy = auth.UserLogin.Name

	key,value := service.PackMesasageRedis(data)
	if err:=service.SaveRedis(key, value); err != nil {
		log.Println("Error Save redis --> ", key, value, err )
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	log.Println("key",key)
	log.Println("val",value)

	if err:=service.LimitTransactionRepository.Save(&data); err!=nil{
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}


	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res

}

// PackMesasageRedis ...
func (service *LimitTransactionService) PackMesasageRedis(data dbmodels.LimitTransaction) (string, string) {
	key := fmt.Sprintf("LIMIT-TRANSACTION:%s:%s:%s:%s", data.UserCategory, data.LevelMerchant, data.FeatureProduct, data.TimeFrame)
	log.Println("key : ", key)
	jsonByte,_ := json.Marshal(data)

	return key, string(jsonByte)

}