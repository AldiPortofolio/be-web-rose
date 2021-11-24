package kafka

import (
	"fmt"
	"github.com/json-iterator/go"
	"net/http"
	hcmodels "ottodigital.id/library/healthcheck/models"
	hcutils "ottodigital.id/library/healthcheck/utils"
	ottoutils "ottodigital.id/library/utils"
	"rose-be-go/redis/redis_single"

	utilhttp "rose-be-go/utils/http"
)

type PublishReq struct {
	Topic     string `json:"topic"`
	Message   string `json:"message"`
	Bytes     []byte `json:"bytes"`
	Timestamp string `json:"timestamp"`
}


var(
	restKafkaName	string
	restKafkaAddress string
	restKafkaHealthCheckKey string
)

func init()  {
	restKafkaName = ottoutils.GetEnv("KAFKA_REST_NAME", "otto-api-publisher")
	restKafkaAddress = ottoutils.GetEnv("KAFKA_PUBLISH_REST", "http://13.228.25.85:8922/v1/publish")
	restKafkaHealthCheckKey = ottoutils.GetEnv("KAFKA_HEALTH_CHECK_KEY", "ROSE_BE_GO_HEALTH_CHECK:KAFKA")

}

func SendPublishKafka(request PublishReq) ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	header := make(http.Header)
	header.Add("Content-Type", "application/json")

	datareq, _ := json.Marshal(request)

	fmt.Println("Headerrrrrr======>", header)


	data, err := utilhttp.HTTPPostWithHeader(restKafkaAddress, request, header)
	fmt.Println("xxxx-----------xxxx")
	fmt.Println("urlSvr", restKafkaAddress)
	fmt.Println("msgreq", string(datareq))
	fmt.Println("header", header)
	fmt.Println("err", err)
	fmt.Println("xxxx-----------xxxx")


	return data, err
}

// GetServiceHealthCheck ..
func GetServiceHealthCheck() hcmodels.ServiceHealthCheck {
	redisClient := redis_single.GetRedisConnection()
	return hcutils.GetServiceHealthCheck(&redisClient, &hcmodels.ServiceEnv{
		Name:           restKafkaName,
		Address:        restKafkaAddress,
		HealthCheckKey: restKafkaHealthCheckKey,
	})
}