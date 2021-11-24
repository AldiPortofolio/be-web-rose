package redis_cluster

import (
	"errors"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/redis.v5"
	"time"
)

// Env ..
type Env struct {
	AddressCluster1 string `envconfig:"HOST_CLUSTER1" default:"34.101.208.23:8077"`
	AddressCluster2 string `envconfig:"HOST_CLUSTER2" default:"34.101.208.23:8078"`
	AddressCluster3 string `envconfig:"HOST_CLUSTER3" default:"34.101.208.23:8079"`
	AddressSlave1 string `envconfig:"HOST_SLAVE1" default:"34.101.208.23:8074"`
	AddressSlave2 string `envconfig:"HOST_SLAVE2" default:"34.101.208.23:8075"`
	AddressSlave3 string `envconfig:"HOST_SLAVE3" default:"34.101.208.23:8076"`
	Address         string `envconfig:"HOST" default:"34.101.222.33"`
	Port            string `envconfig:"PORT" default:"6077"`
}

var (
	ClientRed *redis.ClusterClient
	redisEnv  Env
)

func init() {
	// get environment variables using envconfig
	err := envconfig.Process("REDIS", &redisEnv)
	if err != nil {
		fmt.Println("Failed to get REDIS env:", err)
	}

	ClientRed = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{redisEnv.AddressCluster1, redisEnv.AddressSlave1, redisEnv.AddressCluster2,redisEnv.AddressSlave2 ,redisEnv.AddressCluster3, redisEnv.AddressSlave3},
		RouteByLatency: true,
	})
	pong, err := ClientRed.Ping().Result()

	fmt.Println("Redis Cluster Ping ", pong)
	fmt.Println("Redis Cluster Ping ", err)
}

// GetRedisConnection ..
func GetRedisConnection() *redis.ClusterClient {
	return ClientRed
}

// GetRedisUri ..
func GetRedisUri() string {
	return "redis://" + redisEnv.Address + ":" + redisEnv.Port + "/"
}


/*
 Redis Standard Set
*/
func SaveRedis(key string, val interface{}) error {
	var err error
	for i := 0; i < 3; i++ {
		err = ClientRed.Set(key, val, 0).Err()
		if err == nil {
			break
		}
	}
	return err
}

/*
 Redis Standard Get
*/
func GetRedisKey(Key string) (string, error) {
	val2, err := ClientRed.Get(Key).Result()
	if err == redis.Nil {
		err = errors.New("Key Does Not Exists")
		fmt.Println("keystruct does not exists")
	} else if err != nil {
		fmt.Println("Error : ", err.Error())
	} //else {
	//fmt.Println("keystruct", val2)
	//}
	return val2, err
}

// DelRedisKey ..
func DelRedisKey(key string) error {
	return ClientRed.Del(key).Err()
}

/*
delayto * max = total timeout
*/
func GetDataRedis(key string, delayto, max int) (bool, string) {
	for i := 0; i < max; i++ {
		data, err := GetRedisKey(key)
		fmt.Println(" Err : ", err)
		fmt.Println(" data : ", data)
		if err == nil {
			return true, data
		}
		time.Sleep(time.Duration(delayto) * time.Second)
	}
	return false, ""
}

/*
 Redis Standard Set Expired
*/
func SaveRedisExp(key string, menit string, val interface{}) error {
	var err error
	for i := 0; i < 3; i++ {
		duration, _ := time.ParseDuration(menit)
		err = ClientRed.Set(key, val, duration).Err()
		if err == nil {
			break
		}
		fmt.Println("Error : ", err)
	}
	return err
}

// SaveRedisCounter ..
func SaveRedisCounter(key string) (int64, error) {
	incr := ClientRed.Incr(key)
	return incr.Val(), incr.Err()
}

// SaveRedisCounterAuto ..
func SaveRedisCounterAuto(key string, autonom int64) (int64, error) {
	incr := ClientRed.IncrBy(key, autonom)
	return incr.Val(), incr.Err()
}

// GetRedisCounter ..
func GetRedisCounter(key string) (int64, error) {
	decr := ClientRed.Decr(key)
	return decr.Val(), decr.Err()
}

// GetRedisCounterIncr ..
func GetRedisCounterIncr(key string) (int64, error) {
	decr := ClientRed.Incr(key)
	return decr.Val(), decr.Err()
}

// GetRedisByPattern ..
func GetRedisByPattern(key string) ([]string, error)  {
	data, err := ClientRed.Keys(key).Result()
	return data, err
}
