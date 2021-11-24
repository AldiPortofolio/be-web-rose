package db

import "rose-be-go/models/dbmodels"

type FeeCicilanRepository struct {

}

func InitFeeCicilanRepository() *FeeCicilanRepository {
	return &FeeCicilanRepository{}
}

// Save ..
func (repo *FeeCicilanRepository) Save(req *dbmodels.FeeCicilanSetting) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

func (repo *FeeCicilanRepository) Find() (dbmodels.FeeCicilanSetting, error)  {
	db := GetDbCon()
	var data dbmodels.FeeCicilanSetting
	err := db.Last(&data).Error

	return data, err
}