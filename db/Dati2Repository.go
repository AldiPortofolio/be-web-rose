package db

import "rose-be-go/models/dbmodels"

type Dati2Repository struct {

}

func InitDati2Repository() *Dati2Repository {
	return &Dati2Repository{}
}

func (repo *Dati2Repository) Get(data dbmodels.Dati2) (dbmodels.Dati2, error) {
	db := GetDbCon()
	var res dbmodels.Dati2

	err := db.Where(&data).First(&res).Error

	return res, err
}
