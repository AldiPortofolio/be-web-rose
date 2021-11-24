package db

import "rose-be-go/models/dbmodels"

type KelurahanRepository struct {

}

func InitKelurahanRepository() *KelurahanRepository {
	return &KelurahanRepository{}
}

func (repo *KelurahanRepository) Get(data dbmodels.Kelurahan) (dbmodels.Kelurahan, error) {
	db := GetDbCon()
	var res dbmodels.Kelurahan

	err := db.Where(&data).First(&res).Error

	return res, err
}
