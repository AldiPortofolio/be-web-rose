package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"net/http"
	"ottodigital.id/library/healthcheck"
	"ottodigital.id/library/utils"
	"rose-be-go/db"
)

var	GitCommit              string
var	ReleaseVersion         string
var	DateTime               string
var	TeamCreated            string
var	NameServices           string
var	VersionType            string // SIT , UAT , PROD
var	GitBranchName          string

func HealthCheck(ctx *gin.Context) {
	fmt.Println(">>> HealthCheck - Controller <<<")

	parent := context.Background()
	defer parent.Done()

	dbEnv:= db.Env{}
	err := envconfig.Process("", &dbEnv)
	if err != nil {
		fmt.Println("Failed to get DB env:", err)
	}

	dbReq := healthcheck.HealthCheckDBReq{
		Host:     dbEnv.DbHost,
		Port:     dbEnv.DbPort,
		User:     dbEnv.DbUser,
		Password: dbEnv.DbPass,
		DBName:   dbEnv.DbName,
	}
	dbHealth := healthcheck.GenerateHealthCheckPostgres(dbReq)



	redisReq := healthcheck.HealthCheckRedisReq{
		Host: utils.GetEnv("ROSE_REDIS_HOST", "13.228.23.160"),
		Port: utils.GetEnv("ROSE_REDIS_PORT", "6077"),
	}

	redisHealth := healthcheck.GenerateHealthCheckRedis(redisReq)

	redisClusterReq := healthcheck.HealthCheckRedisReq{
		HostCluster: []string{utils.GetEnv("REDIS_HOST_CLUSTER1", "13.228.23.160:8079"), utils.GetEnv("REDIS_HOST_CLUSTER2", "13.228.23.160:8078"), utils.GetEnv("REDIS_HOST_CLUSTER3", "13.228.23.160:8077")},
	}

	redisClusterHealth := healthcheck.GenerateHealthCheckRedisCluster(redisClusterReq)

	// Genereate HealthCheck Kafka
	serviceReqKafka := healthcheck.CheckServiceReq{
		ServiceName: "otto-api-publisher",
		Host:        utils.GetEnv("KAFKA_PUBLISH_REST", "http://13.228.25.85:8922/"),
		Endpoint:    "v1/publish",
		Method:      "POST",
	}
	serviceKafkaHealth := healthcheck.GenerateHealthCheckService(serviceReqKafka)



	res := healthcheck.GenerateResponseHealthCheck(<-dbHealth, <-redisHealth, <-redisClusterHealth,<-serviceKafkaHealth)



	ctx.JSON(http.StatusOK, res)


}

func Version (c *gin.Context) {

	c.String(http.StatusOK, "Version Release : [%s] \n "+
		"Branch Name : [%s]\n "+
		"GitCommit : [%s]\n "+
		"DateTime : [%s] \n "+
		"Team By: [%s] \n"+
		"Services :[%s] \n"+
		"VersionType : [%s]\n", ReleaseVersion, GitBranchName, GitCommit, DateTime, TeamCreated, NameServices, VersionType)
}