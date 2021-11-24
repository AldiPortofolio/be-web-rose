package services

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"rose-be-go/models"
)

type TestService struct {
	General 		models.GeneralModel
}

func InitTestService(gen models.GeneralModel)  *TestService {
	return &TestService{
		General:gen,
	}

}

func (service *TestService)GetData(data string) models.Response {

	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("TestService: GetData",
		zap.String("data", data))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "TestService: Generate")
	defer span.Finish()

	fmt.Println(">>> TestService - Service <<<")

	var res models.Response

	res.Data = data
	return res
}