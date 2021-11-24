package db

import "rose-be-go/models/dbmodels"

type KecamatanRepository struct {

}

func InitKecamatanRepository() *KecamatanRepository {
	return &KecamatanRepository{}
}

func (repo *KecamatanRepository) Get(data dbmodels.Kecamatan) (dbmodels.Kecamatan, error) {
	db := GetDbCon()
	var res dbmodels.Kecamatan

	err := db.Where(&data).First(&res).Error

	return res, err
}
