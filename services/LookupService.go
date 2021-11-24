package services

import (
	"fmt"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"time"
)

// LookupService struct
type LookupService struct {
	LookupRepository *db.LookupRepository
}

// InitLookupService ...
func InitLookupService() *LookupService {
	return &LookupService{
		LookupRepository: db.InitLookupRepository(),
	}
}

// Save ...
func (service *LookupService) Save(req dto.ReqLookupDto, res *models.Response)  {
	fmt.Println("<< LookupService - Save >>")

	var data dbmodels.Lookup
	var err error

	data.CreatedAt = time.Now()
	data.CreatedBy = auth.UserLogin.Name

	if req.ID > 0 {
		data, err = service.LookupRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return
		}
	}
	data.UpdatedAt = time.Now()
	data.UpdatedBy = auth.UserLogin.Name
	data.LookupGroup = req.LookupGroup
	data.Name = req.Name
	data.Code = req.Code
	data.Descr = req.Descr
	data.OrderNo = req.OrderNo

	if err:=service.LookupRepository.Save(&data); err!=nil{
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG


}

// Filter ...
func (service *LookupService) Filter(req dto.ReqLookupDto, res *models.Response)  {
	fmt.Println("<< LookupService - Filter >>")


	data, total, err := service.LookupRepository.Filter(req)
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