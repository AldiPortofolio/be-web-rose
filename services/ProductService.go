package services

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"log"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type ProductService struct {
	General models.GeneralModel
	ProductRepository *db.ProductRepository
}

func InitProductService(gen models.GeneralModel) *ProductService {
	return &ProductService{
		General:gen,
		ProductRepository: db.InitProductRepository(),
	}
}

func (service *ProductService) Filter(req dto.ReqProductDto) models.Response {
	fmt.Println(">>> ProductService - Filter <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("ProductService: Filter",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "ProductService: Filter")
	defer span.Finish()

	var res models.Response
	list, total, err := service.ProductRepository.Filter(req)
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

func (service *ProductService) Save(req dto.ReqProductDto) models.Response {
	fmt.Println(">>> ProductService - save <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("ProductService: save",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "ProductService: save")
	defer span.Finish()

	var res models.Response

	var data dbmodels.Product
	var err error

	if req.ID > 0  {
		data, err = service.ProductRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return res
		}
	}

	data = dbmodels.Product{
		ID:    req.ID,
		Code:  req.Code,
		Name:  req.Name,
		Title: req.Title,
		Desc:  req.Desc,
		Notes: req.Notes,
		Seq:   req.Seq,
	}

	if err:=service.ProductRepository.Save(&data); err!=nil{
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}
