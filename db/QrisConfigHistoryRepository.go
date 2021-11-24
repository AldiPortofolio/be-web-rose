package db

import (
	"rose-be-go/models/dbmodels"
)

type QrisConfigHistoryRepository struct {

}

func InitQrisConfigHistoryRepository()  *QrisConfigHistoryRepository{
	return &QrisConfigHistoryRepository{}
}

func (repo *QrisConfigHistoryRepository) Save(req dbmodels.QrisConfigHistory) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}