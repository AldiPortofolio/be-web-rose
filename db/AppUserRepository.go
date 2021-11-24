package db

import (
	"ottodigital.id/library/logger/v2"
	"rose-be-go/models/dbmodels"
	
)

type AppUserRepository struct {
	Ottolog logger.OttologInterface

}

func InitAppUserRepository(logs logger.OttologInterface) *AppUserRepository {
	return &AppUserRepository{
		Ottolog:logs,
	}
}

func (r *AppUserRepository)Save(req *dbmodels.AppUser) error {
	db := GetDbCon()

	if err:= db.Save(&req).Error; err !=nil{
		r.Ottolog.Error("Error save to db " + err.Error())
		return err
	}

	return nil
}

func (r *AppUserRepository)FindByUserName(username string) (dbmodels.AppUser, error) {
	db := GetDbCon()

	var res dbmodels.AppUser


	if err:= db.Where("user_name = ?", username).First(&res).Error; err !=nil{
		r.Ottolog.Error("Error find data " + err.Error())
		return res, err
	}

	return res, nil
}
