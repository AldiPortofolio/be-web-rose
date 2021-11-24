package db

import (
	"rose-be-go/models/dbmodels"
)

type KategoriBisnisRepository struct {

}

func InitKategoriBisnisRepository() *KategoriBisnisRepository {
	return &KategoriBisnisRepository{}
}

func (r *KategoriBisnisRepository)GetAll() ([]dbmodels.KategoriBisnis, error) {
	db := GetDbCon()

	var res []dbmodels.KategoriBisnis

	err := db.Find(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil

}