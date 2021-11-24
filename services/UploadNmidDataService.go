package services

import (
	"rose-be-go/db"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models"
)

// GetDataUploadNmid ...
func GetDataUploadNmid(req dbmodels.UploadNmidData) models.Response {
	var res models.Response
	var total int

	list, total, err := db.InitUploadNmidDataRepository().GetDataUploadNmid(req)
	if err != nil {
		res.ErrCode = "05"
		return res
	}

	// sort by selain DB
	//sort.SliceStable(userCreatedDate, func(i, j int) bool {
	//	return userCreatedDate[i].Date > list[j].Date
	//})

	res.ErrCode = "00"
	res.ErrDesc = "Success"
	res.TotalData = total
	res.Contents = list

	return res
}



