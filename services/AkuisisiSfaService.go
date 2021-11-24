package services

import (
	"fmt"
	"log"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dto"
)

// AkuisisiSfaService struct
type AkuisisiSfaService struct {
	AkuisisiSfaRepository *db.AkuisisiSfaRepository
}

// InitAkuisisiSfaService ...
func InitAkuisisiSfaService() *AkuisisiSfaService {
	return &AkuisisiSfaService{
		AkuisisiSfaRepository: db.InitAkuisisiSfaRepository(),
	}
}

// Filter ...
func (service AkuisisiSfaService) Filter(req dto.ReqAkuisisiSfa) models.Response {
	fmt.Println(">>> AkuisisiSfaService - Filter <<<")

	var res models.Response
	list, total, err := service.AkuisisiSfaRepository.Filter(req)
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