package redis_single

import (
	"log"
	"testing"
)

func TestGetKeysByPattern(t *testing.T) {
	res,_:= GetKeysByPattern("*")

	i := 0

	for i=0; i< len(res); i++  {
		log.Println(res[i])
		a,_:=GetRedisKey(res[i])
		log.Println(a)
	}
	log.Println(i)
}

func TestGetRedisKey(t *testing.T) {
	data, err:= GetRedisKey("QRIS:SFA:REGISTER_BY_SFA:087770366577")

	log.Println(err)
	log.Println(data)
}

func TestDelRedisKey(t *testing.T) {
	err := DelRedisKey("QRIS:SFA:REGISTER_BY_SFA:087770366577")
	log.Println(err)

}