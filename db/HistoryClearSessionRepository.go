package db

import "rose-be-go/models/dbmodels"

type HistoryClearSessionRepository struct {

}

func InitHistoryClearSessionRepository() *HistoryClearSessionRepository {
	return &HistoryClearSessionRepository{}
}

func (repo *HistoryClearSessionRepository) Save(req dbmodels.HistoryClearSession) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

func (repo *HistoryClearSessionRepository) GetLastUpdated() (dbmodels.HistoryClearSession, error) {
	db := GetDbCon()

	var data dbmodels.HistoryClearSession

	err := db.Order("id desc").First(&data).Error

	return data, err
}