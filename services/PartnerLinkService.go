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

	"ottodigital.id/library/logger/v2"
)

// PartnerLinkService struct
type PartnerLinkService struct {
	Ottolog logger.OttologInterface
	PartnerLinkRepository *db.PartnerLinkRepository
}

// InitPartnerLinkService ...
func InitPartnerLinkService(logs logger.OttologInterface)*PartnerLinkService  {
	return &PartnerLinkService{
		Ottolog:logs,
		PartnerLinkRepository: db.InitPartnerLinkRepository(logs),
	}
}

// Save ...
func (svc *PartnerLinkService) Save(req dto.ReqPartnerLinkDto, res *models.Response)  {
	svc.Ottolog.Info("PartnerLinkService - Save")

	var data dbmodels.PartnerLink
	
	data.Code = req.Code
	data.MerchantId = req.MerchantId
	data.PartnerId = req.PartnerId
	data.CreatedBy = auth.UserLogin.Name
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	

	if err := svc.PartnerLinkRepository.Save(&data); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return
	}
	
	res.Contents = data
	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

}

// Filter ...
func (svc *PartnerLinkService) Filter(req dto.ReqPartnerLinkDto, res *models.Response)  {
	svc.Ottolog.Info("PartnerLinkService - Filter")


	data, total, err := svc.PartnerLinkRepository.Filter(req)
	if err != nil {

		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND

		svc.Ottolog.Error(fmt.Sprintf("Failed to get data from database: %s", fmt.Sprintf("ERR:%s", err.Error())))
		return
	}

	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.ErrCode = constants.ERR_SUCCESS
	res.Contents = data
	res.TotalData = total

}

func (svc *PartnerLinkService) Delete(id int, res *models.Response)  {
	svc.Ottolog.Info("PartnerLinkService - Delete")


	if err := svc.PartnerLinkRepository.Delete(id); err != nil {
		res.ErrCode = constants.EC_FAIL_DELETE
		res.ErrDesc = constants.EC_FAIL_DELETE_DESC
		return
	}
	
	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

}