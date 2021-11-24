package redis_cluster

import (
	"log"
	"testing"
)

func TestGetRedisKey(t *testing.T) {
	res, err :=GetRedisKey("OTTOMART:ANDROID-VERSION")
	log.Println(err)
	log.Println(res)
}