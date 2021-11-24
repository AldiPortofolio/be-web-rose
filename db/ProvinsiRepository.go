package db

import "rose-be-go/models/dbmodels"

type ProvinsiRepository struct {

}

func InitProvinsiRepository() *ProvinsiRepository {
	return &ProvinsiRepository{}
}

func (repo *ProvinsiRepository) Get(data dbmodels.Provinsi) (dbmodels.Provinsi, error) {
	db := GetDbCon()
	var res dbmodels.Provinsi

	err := db.Where(&data).First(&res).Error

	return res, err
}
