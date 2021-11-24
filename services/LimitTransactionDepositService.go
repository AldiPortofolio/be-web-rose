package services

import (
	"encoding/json"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"log"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	redisModels "rose-be-go/models/redis"
	"rose-be-go/redis/redis_cluster"
)

// LimitTransactionDepositService struct
type LimitTransactionDepositService struct {
	General 					models.GeneralModel
	LimitTransactionDepositRepository 	*db.LimitTransactionDepositRepository
	SaveRedis 					func(string, interface{}) (error)
	GetRedis 					func(string) (string,error)
}

// InitLimitTransactionDepositService ...
func InitLimitTransactionDepositService(gen models.GeneralModel) *LimitTransactionDepositService {
	return &LimitTransactionDepositService{
		General:gen,
		LimitTransactionDepositRepository: db.InitLimitTransactionDepositRepository(),
		SaveRedis: redis_cluster.SaveRedis,
		GetRedis: redis_cluster.GetRedisKey,
	}
}

// Filter ...
func (service *LimitTransactionDepositService) Filter(req dto.ReqLimitTransactionDepositDto) models.Response {
	fmt.Println(">>> LimitTransactionDepositService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("LimitTransactionDepositService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "LimitTransactionDepositService: Filter")
	defer span.Finish()

	var res models.Response
	dataUserCategory, err := service.LimitTransactionDepositRepository.GetUserCategory()
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		return res
	}
	dataLevelMerchant, err := service.LimitTransactionDepositRepository.GetLevelMerchant()
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		return res
	}

	var dataRedis redisModels.LimitTransactionDeposit
	var data []redisModels.LimitTransactionDeposit

	if req.MemberType != "" && req.Category == "" {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC + " - Category cannot be empty"
		return res
	}else if req.MemberType == "" && req.Category != "" {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC + " - Member Type cannot be empty"
		return res
	}else if req.MemberType != "" && req.Category != "" {
		fmt.Println("3")
		getData,_ := service.GetRedis(fmt.Sprintf("LIMIT-TRANSACTION:DEPOSIT:%s:%s", req.Category, req.MemberType))
		json.Unmarshal([]byte(getData), &dataRedis)
		data = append(data, dataRedis)
	}else{
		fmt.Println("4")
		for _, val := range dataUserCategory {
			for _, val2 := range dataLevelMerchant {
				fmt.Println(fmt.Printf("LIMIT-TRANSACTION:DEPOSIT:%s:%s", val.Code, val2.Code))
				getData,_ := service.GetRedis(fmt.Sprintf("LIMIT-TRANSACTION:DEPOSIT:%s:%s", val.Code, val2.Code))
				if getData != "" {
					json.Unmarshal([]byte(getData), &dataRedis)
					data = append(data, dataRedis)
				}
			}
		}
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.Contents = data
	res.TotalData = len(data)

	sugarLogger.Info("Response", zap.Any("res", res))

	return res

}

// Save ...
func (service *LimitTransactionDepositService) Save(req dto.ReqLimitTransactionDepositDto) models.Response {
	fmt.Println(">>> LimitTransactionDepositService - save <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("LimitTransactionDepositService: save",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "LimitTransactionDepositService: save")
	defer span.Finish()

	var res models.Response

	var data redisModels.LimitTransactionDeposit

	data.Category = req.Category
	data.MemberType = req.MemberType
	data.MinLimit = req.MinLimit
	data.MaxLimit = req.MaxLimit

	key,value := service.PackMesasageRedis(data)
	if err:=service.SaveRedis(key, value); err != nil {
		log.Println("Error Save redis --> ", key, value, err )
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	getData,_ := service.GetRedis("LIMIT-TRANSACTION:DEPOSIT:op:gold")
	fmt.Println("get data redis ->", getData)

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res

}

// PackMesasageRedis ...
func (service *LimitTransactionDepositService) PackMesasageRedis(data redisModels.LimitTransactionDeposit) (string, string) {
	key := fmt.Sprintf("LIMIT-TRANSACTION:DEPOSIT:%s:%s", data.Category, data.MemberType)
	log.Println("key : ", key)
	jsonByte,_ := json.Marshal(data)

	return key, string(jsonByte)

}