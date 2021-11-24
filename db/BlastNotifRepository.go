package db

import "rose-be-go/models/dbmodels"

type BlastNotifRepository struct {

}

func InitBlastNotifRepository() *BlastNotifRepository {
	return &BlastNotifRepository{}
}

func (r *BlastNotifRepository)Save(data *dbmodels.BlastNotif)  error{

	db:= GetDbCon()

	err :=db.Save(&data).Error
	return err
}
