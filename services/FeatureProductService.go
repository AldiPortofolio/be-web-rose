package services

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"log"
	"rose-be-go/constants"
	"rose-be-go/db"
	_"rose-be-go/hosts/minio"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

// FeatureProductService struct
type FeatureProductService struct {
	General models.GeneralModel
	FeatureProductRepository *db.FeatureProductRepository
	//MinioHost *minio.MinioHost
}

// InitFeatureProductService ...
func InitFeatureProductService(gen models.GeneralModel) *FeatureProductService {
	return &FeatureProductService{
		General:gen,
		FeatureProductRepository: db.InitFeatureProductRepository(),
		//MinioHost: minio.InitMinioHost(),
	}
}

// Filter ...
func (service *FeatureProductService) Filter(req dto.ReqFiturProductDto) models.Response {
	fmt.Println(">>> FeatureProductService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("FeatureProductService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "FeatureProductService: Filter")
	defer span.Finish()

	var res models.Response
	list, total, err := service.FeatureProductRepository.Filter(req)
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
func (service *FeatureProductService) Save(req dto.ReqFiturProductDto) models.Response {
	fmt.Println(">>> FeatureProductService - save <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("FeatureProductService: save",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "FeatureProductService: save")
	defer span.Finish()

	var res models.Response

	var data dbmodels.FiturProduct
	var err error

	if req.ID > 0  {
		data, err = service.FeatureProductRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return res
		}
	}

	//if strings.Contains(req.Icon, "base64")  {
	//	reqMinio := minio.UploadRequest{
	//		BucketName: "rose",
	//		Data:        req.Icon,
	//		NameFile:    "icon-" + strings.ToLower(strings.Replace(req.Name, " ", "-", -1)) + "-" +  strconv.Itoa(rand.Intn(100000000)) + ".jpeg",
	//		ContentType: "image/jpeg",
	//	}
	//
	//	data, err := service.MinioHost.Send(reqMinio, constants.MinioUpload)
	//	if err != nil {
	//		log.Println("err -> ", err)
	//		res.ErrCode = constants.EC_FAIL_SEND_TO_HOST
	//		res.ErrDesc = constants.EC_FAIL_SEND_TO_HOST_DESC
	//		return res
	//	}
	//
	//	resMinio := minio.UploadResponse{}
	//	json.Unmarshal(data, &resMinio)
	//
	//	req.Icon = resMinio.Url
	//}

	data = dbmodels.FiturProduct{
		ID:        req.ID,
		ProductID: req.ProductID,
		Code:      req.Code,
		Icon:      req.Icon,
		Name:      req.Name,
		Notes:     req.Notes,
		Seq:       req.Seq,
		Url: 	   req.Url,
	}

	if err:=service.FeatureProductRepository.Save(&data); err!=nil{
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}