package services

import (
	"encoding/json"
	"fmt"
	"log"
	"rose-be-go/constants"
	"rose-be-go/hosts/minio"
	"rose-be-go/models"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// UploadImageService struct
type UploadImageService struct {
	General                   models.GeneralModel
	
}

// InitUploadImageService ...
func InitUploadImageService(gen models.GeneralModel) *UploadImageService {
	return &UploadImageService{
		General:          gen,
		
	}
}

// Upload ...
func (service *UploadImageService) Upload(req minio.UploadRequest) models.Response {
	fmt.Println(">>> UploadImage - Upload <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("UploadImage: Upload",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "UploadImage: upload")
	defer span.Finish()

	var res models.Response

	data, err := minio.InitMinioHost().Send(req, constants.MinioUpload)
	if err != nil {
		log.Println("err -> ", err)
		res.ErrCode = constants.EC_FAIL_SEND_TO_HOST
		res.ErrDesc = constants.EC_FAIL_SEND_TO_HOST_DESC
		return res
	}

	resMinio := minio.UploadResponse{}
	json.Unmarshal(data, &resMinio)
	fmt.Println(string(data))

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}


