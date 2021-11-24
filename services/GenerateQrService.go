package services

import (
	"encoding/json"
	"fmt"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/constants"
	"rose-be-go/host/stringbuilder"
	"rose-be-go/models"
	"rose-be-go/models/sbmodels"
)

// GenerateQrService struct
type GenerateQrService struct {
	Ottolog logger.OttologInterface
	SbHost *stringbuilder.SbHost
}

// InitGenerateQrService ...
func InitGenerateQrService(logs logger.OttologInterface) *GenerateQrService {
	return &GenerateQrService{
		Ottolog:logs,
		SbHost: stringbuilder.InitSbHost(logs),
	}
}

// GenerateQr ...
func (svc *GenerateQrService)GenerateQr(req sbmodels.ReqGenerateQr, res *models.Response)  {
	svc.Ottolog.Info("GenerateQrService - GenerateQr")

	qr, err := svc.SbHost.GenerateQr(req)
	if err != nil {
		fmt.Println("failed generate qr")
		res.ErrCode = constants.EC_FAILED_GENERATE_QR
		res.ErrDesc = constants.EC_FAILED_GENERATE_QR_DESC
		return

	}

	var qrData sbmodels.ResGenerateQr
	json.Unmarshal(qr, &qrData)
	if qrData.Rc != constants.ERR_SUCCESS {
		fmt.Println("failed generate qr")
		res.ErrCode = constants.EC_FAILED_GENERATE_QR
		res.ErrDesc = constants.EC_FAILED_GENERATE_QR_DESC + " " +qrData.Msg
		return
	}

	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.ErrCode = constants.ERR_SUCCESS
	res.Contents = qrData.QrData
}