package services

import (
	"fmt"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dto"
)

type MerchantMasterTagService struct {
	MerchantMasterTagRepository *db.MerchantMasterTagRepository
}

func InitMerchantMasterTagService()  *MerchantMasterTagService{
	return &MerchantMasterTagService{
		MerchantMasterTagRepository: db.InitMerchantMasterTagRepository(),
	}
}

func (svc *MerchantMasterTagService) Filter(req dto.ReqMerchantMasterTagDto, res *models.Response)  {
	fmt.Println("MerchantMasterTagService - Filter")


	data, total, err := svc.MerchantMasterTagRepository.Filter(req)
	if err != nil {

		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND

		fmt.Println(fmt.Printf("Failed to get data from database: %s", fmt.Sprintf("ERR:%s", err.Error())))
		return
	}

	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.ErrCode = constants.ERR_SUCCESS
	res.Contents = data
	res.TotalData = total

}

func (svc *MerchantMasterTagService) FindByMid(req dto.ReqMerchantMasterTagByMidDto, res *models.Response)  {
	fmt.Println("MerchantMasterTagService - FindByMid")


	data, total, err := svc.MerchantMasterTagRepository.FindByMid(req)
	if err != nil {

		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND

		fmt.Println(fmt.Printf("Failed to get data from database: %s", fmt.Sprintf("ERR:%s", err.Error())))
		return
	}

	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.ErrCode = constants.ERR_SUCCESS
	res.Contents = data
	res.TotalData = total

}


