package db

import "rose-be-go/models/dbmodels"

type HistoryAppVersionRepository struct {

}

func InitHistoryAppVersionRepository() *HistoryAppVersionRepository {
	return &HistoryAppVersionRepository{}
}

func (repo *HistoryAppVersionRepository) Save(req dbmodels.HistoryAppVersion) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}