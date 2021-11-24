package services

import "rose-be-go/redis/redis_cluster"


// LimitDepositService struct
type LimitDepositService struct {
	GetRedisKey  func(string) (string, error)
	SaveRedisKey  func(string, interface{}) (error)
}

// InitLimitDepositService ...
func InitLimitDepositService() *LimitDepositService {
	return &LimitDepositService{
		GetRedisKey: redis_cluster.GetRedisKey,
		SaveRedisKey: redis_cluster.SaveRedis,
	}
}

var (
	
)

// Update ...
func (service *LimitDepositService) Update()  {
	
}