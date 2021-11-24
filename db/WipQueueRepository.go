package db

import (
	"encoding/json"
	"rose-be-go/models/dbmodels"
)

type WipQueueRepository struct {

}

func InitWipQueueRepository() *WipQueueRepository {
	return  &WipQueueRepository{}
}


func (repo *WipQueueRepository) GetWipQueueByUserAndStatus(user string, status string) (dbmodels.WipQueue, error) {
	db := GetDbCon()
	var wipQueue dbmodels.WipQueue
	err := db.Where("user_name = ? and status_registration = ?", user, status).Order("id asc").First(&wipQueue).Error // query

	return wipQueue, err
}

func GetReasonWipQueue(wipid int) (dbmodels.ReportRejected, error) {
	db := GetDbCon()
	var data dbmodels.ReportRejected
	var res dbmodels.ResultRejected
	var wipQueue dbmodels.WipQueue

	//err := db.Table("wip_queue a").Select("a.store_name, d.owner_first_name owner_name, c.transaction_date, c.username, c.reason, a.wip_id id_merchants").Joins("left join merchant_wip b on a.wip_id = b.id").Joins("left join merchant_wip_status_list2 c on a.wip_id = c.merchant_wip_id and a.status_registration = c.registration_status").Joins("left join owner_wip d on b.owner_wip_id = d.id").Where("a.status_registration like '%REJECTED%'").Limit(limit).Offset((page-1)*limit).Limit(-1).Offset(0).Order("a.id desc").Find(&data).Count(&total).Error

	err := db.Where("wip_id = ? AND status_registration IN ('VVIP_REGISTERED', 'VIP_REGISTERED', 'REGISTERED')", wipid).Order("id desc").First(&wipQueue).Error

	json.Unmarshal([]byte(wipQueue.Value), &res)

	data.IdMerchant = wipQueue.WipId
	data.Reason = res.Reason
	data.Username = wipQueue.UserName
	data.TransactionDate = wipQueue.TransactionTime
	data.OwnerName = res.OwnerWip.OwnerFirstName
	data.StoreName = wipQueue.StoreName
	data.Status = res.StatusRegistration

	if err != nil {
		return data, err
	}

	return data, nil
}
