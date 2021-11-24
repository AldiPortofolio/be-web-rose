package services

import (
	"fmt"
	"log"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dto"
)

// AkuisisiSfaFailedService struct
type AkuisisiSfaFailedService struct {
	AkuisisiSfaFailedRepository *db.AkuisisiSfaFailedRepository
}

// InitAkuisisiSfaFailedService ...
func InitAkuisisiSfaFailedService() *AkuisisiSfaFailedService {
	return &AkuisisiSfaFailedService{
		AkuisisiSfaFailedRepository: db.InitAkuisisiSfaFailedRepository(),
	}
}

// Filter ...
func (service AkuisisiSfaFailedService) Filter(req dto.ReqAkuisisiSfaFailed) models.Response {
	fmt.Println(">>> AkuisisiSfaFailedService - Filter <<<")

	var res models.Response
	list, total, err := service.AkuisisiSfaFailedRepository.Filter(req)
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


	return res
}