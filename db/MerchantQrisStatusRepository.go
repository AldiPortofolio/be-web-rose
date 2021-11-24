package db

import (
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"rose-be-go/utils"
	"time"
)

type MerchantQrisStatusRepository struct {

}

func InitMerchantQrisStatusRepository() *MerchantQrisStatusRepository {
	return &MerchantQrisStatusRepository{}
}

func (repo *MerchantQrisStatusRepository) Filter(req dto.ReqFilterDto) ([]dbmodels.MerchantQrisStatus, int, error) {
	db := GetDbCon()

	var data []dbmodels.MerchantQrisStatus
	var total int

	if req.Mid != "" {
		db = db.Where("merchant_outlet_id like ?", "%"+req.Mid+"%")

	}
	if req.Status != "" {
		db = db.Where("qris_status = ?", req.Status)
	} else {
		db = db.Where("qris_status > 0")

	}
	if req.AgentID != "" {
		db = db.Where("agent_id like ?", "%"+req.AgentID+"%")
	}
	if req.RequestDate != "" {
		requestDateStart := utils.ConverDateStringToTime(req.RequestDate)
		requestDateEnd := requestDateStart.Add(time.Hour*24)
		db = db.Where("qris_request_date >= ? AND qris_request_date < ?", requestDateStart, requestDateEnd)
	}
	if req.InstallDate != "" {
		requestDateStart := utils.ConverDateStringToTime(req.InstallDate)
		requestDateEnd := requestDateStart.Add(time.Hour*24)
		db = db.Where("qris_install_date >= ? AND qris_request_date < ?", requestDateStart, requestDateEnd)
	}


	err := db.Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Order("id asc").Find(&data).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		return data, 0, err
	}

	return data, total, nil
}