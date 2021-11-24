package services

import (
	"encoding/json"
	"fmt"
	"log"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/redis/redis_cluster"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// ConfigVaMerchantGroupService struct
type ConfigVaMerchantGroupService struct {
	General          models.GeneralModel
	ConfigVaMerchantGroupRepository *db.ConfigVaMerchantGroupRepository
	SaveRedis 					func(string, interface{}) (error)
	GetRedis 					func(string) (string,error)
}

// InitConfigVaMerchantGroupService ...
func InitConfigVaMerchantGroupService(gen models.GeneralModel) *ConfigVaMerchantGroupService {
	return &ConfigVaMerchantGroupService{
		General:          gen,
		ConfigVaMerchantGroupRepository: db.InitConfigVaMerchantGroupRepository(),
		SaveRedis: redis_cluster.SaveRedis,
		GetRedis: redis_cluster.GetRedisKey,
	}
}

// Save ...
func (service *ConfigVaMerchantGroupService) Save(req dbmodels.ConfigVaMerchantGroup) models.Response {
	fmt.Println(">>> ConfigVaMerchantGroupService - save <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("ConfigVaMerchantGroupService: save",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "ConfigVaMerchantGroupService: save")
	defer span.Finish()

	var res models.Response

	if err := service.ConfigVaMerchantGroupRepository.Save(&req); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	key := fmt.Sprintf("CONFIG_VA_MERCHANT_GROUP:%d", req.MerchantGroupId)
	_ = redis_cluster.DelRedisKey(key)

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}

// SaveToRedis ...
func (service *ConfigVaMerchantGroupService) SaveToRedis(req dbmodels.ConfigVaMerchantGroup) models.Response {
	fmt.Println(">>> ConfigVaMerchantGroupService - SaveToRedis <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("ConfigVaMerchantGroupService: SaveToRedis",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "ConfigVaMerchantGroupService: SaveToRedis")
	defer span.Finish()

	var res models.Response

	key,value := service.PackMesasageRedis(req)
	if err:=service.SaveRedis(key, value); err != nil {
		log.Println("Error Save redis --> ", key, value, err )
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}
	
	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}

// FindByBroupId ...
func (service *ConfigVaMerchantGroupService) FindByGroupId (id int64) models.Response {
	fmt.Println(">>> ConfigVaMerchantGroupService - FindByGroupId <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("ConfigVaMerchantGroupService: FindByGroupId",
		zap.Any("req", id))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "ConfigVaMerchantGroupService: FindByGroupId")
	defer span.Finish()

	var res models.Response

	data, err := service.ConfigVaMerchantGroupRepository.FindByGroupId(id)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		return res
	}

	res.Data = data
	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}

// FindByGroupIdFromRedis ...
func (service *ConfigVaMerchantGroupService) FindByGroupIdFromRedis (id int64) models.Response {
	fmt.Println(">>> ConfigVaMerchantGroupService - FindByGroupIdFromRedis <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("ConfigVaMerchantGroupService: FindByGroupIdFromRedis",
		zap.Any("req", id))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "ConfigVaMerchantGroupService: FindByGroupIdFromRedis")
	defer span.Finish()
	var data dbmodels.ConfigVaMerchantGroup
	var res models.Response

	key := fmt.Sprintf("CONFIG_VA_MERCHANT_GROUP:%d", id)
	getData, err := service.GetRedis(key)
	fmt.Println("get data redis ->", getData)
	
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		return res
	}
	json.Unmarshal([]byte(getData),&data)
	res.Data = data
	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}

// PackMesasageRedis ...
func (service *ConfigVaMerchantGroupService) PackMesasageRedis(data dbmodels.ConfigVaMerchantGroup) (string, string) {
	key := fmt.Sprintf("CONFIG_VA_MERCHANT_GROUP:%d", data.MerchantGroupId)
	log.Println("key : ", key)
	jsonByte,_ := json.Marshal(data)

	return key, string(jsonByte)

}
