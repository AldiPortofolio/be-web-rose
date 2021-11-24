package auth

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ottodigital.id/library/utils"
	"rose-be-go/models"
	"rose-be-go/redis/redis_single"
	"strings"
)

var (
	redisPrefixToken string
	UserLogin	models.User
)

func init() {
	redisPrefixToken = utils.GetEnv("ROSE_REDIS_PREFIX_TOKEN","ROSE:TOKEN:")
}


func CekToken(c *gin.Context)  {

	res := models.Response{}
	tokenString := c.Request.Header.Get("Authorization")

	if !strings.HasPrefix(tokenString, "Bearer ") {
		//res.ErrCode = kons.ERR_CODE_54
		//res.ErrDesc = kons.ERR_CODE_54_MSG

		res.ErrCode = "54"
		res.ErrDesc = "Invalid Authorization !"

		c.JSON(http.StatusUnauthorized, res)
		c.Abort()
		return
	}


	log.Println("tokenString ", tokenString)
	log.Println(redisPrefixToken+tokenString)

	token := redisPrefixToken+tokenString
	log.Println("token--> ", token)

	userRedis,err := redis_single.GetRedisKey(token)
	log.Println("err -> ", err)

	if err != nil {
		res.ErrCode = "54"
		res.ErrDesc = "Invalid Authorization !"

		c.JSON(http.StatusUnauthorized, res)
		c.Abort()
		return
	}


	byteUser :=  []byte(userRedis)
	if err := json.Unmarshal(byteUser, &UserLogin); err != nil {
		fmt.Println("failed unmarshal user login", err)
		return
	}
	//log.Println(string(byteUser))

}
