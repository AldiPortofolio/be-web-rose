package services

import (
	"encoding/json"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"log"
	"math/rand"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/hosts/minio"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"strconv"
	"strings"
)

// ImageManagementService struct
type ImageManagementService struct {
	General models.GeneralModel
	ImageManagementRepository *db.ImageManagementRepository
	MinioHost *minio.MinioHost
}

// InitImageManagementService ...
func InitImageManagementService(gen models.GeneralModel) *ImageManagementService {
	return &ImageManagementService{
		General:gen,
		ImageManagementRepository: db.InitImageManagementRepository(),
		MinioHost: minio.InitMinioHost(),
	}
}

// Filter ...
func (service *ImageManagementService) Filter(req dto.ReqImageManagementDto) models.Response {
	fmt.Println(">>> ImageManagementService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("ImageManagementService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "ImageManagementService: Filter")
	defer span.Finish()

	var res models.Response
	list, total, err := service.ImageManagementRepository.Filter(req)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.TotalData = total
		res.Contents = list
		return res
	}

	log.Println("total -->", total)

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.TotalData = total
	res.Contents = list

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}

// Save ...
func (service *ImageManagementService) Save(req dto.ReqImageManagementDto) models.Response {
	fmt.Println(">>> ImageManagementService - save <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("ImageManagementService: save",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "ImageManagementService: save")
	defer span.Finish()

	var res models.Response

	var data dbmodels.ImageManagement
	var err error

	if req.ID > 0  {
		data, err = service.ImageManagementRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return res
		}
	}

	if req.Images != ""  {
		reqMinio := minio.UploadRequest{
			BucketName: "rose",
			Data:        req.Images,
			NameFile:    strings.ToLower(strings.Replace(req.Name, " ", "-", -1)) + "-" +  strconv.Itoa(rand.Intn(100000000)) + ".jpeg",
			ContentType: "image/jpeg",
		}

		data, err := service.MinioHost.Send(reqMinio, constants.MinioUpload)
		if err != nil {
			log.Println("err -> ", err)
			res.ErrCode = constants.EC_FAIL_SEND_TO_HOST
			res.ErrDesc = constants.EC_FAIL_SEND_TO_HOST_DESC
			return res
		}

		resMinio := minio.UploadResponse{}
		json.Unmarshal(data, &resMinio)

		req.URL = resMinio.Url
	}

	data = dbmodels.ImageManagement{
		ID:    req.ID,
		Name:  req.Name,
		URL:   req.URL,
		Notes: req.Notes,
	}

	if err:=service.ImageManagementRepository.Save(&data); err!=nil{
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}
