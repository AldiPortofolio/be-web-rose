package db

import (
	"encoding/json"
	"log"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"testing"
)

func TestUserCategoryRepository_FindByID(t *testing.T) {
	res, err := InitUserCategoryRepository().FindByID(1)

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println(err)
}

func TestUserCategoryRepository_Filter(t *testing.T) {
	req := dto.ReqUserCategoryDto{
		ID:    1,
		Code:  "",
		Name:  "",
		Notes: "",
		Seq:   0,
		Limit: 10,
		Page:  1,
	}
	res, total, err := InitUserCategoryRepository().Filter(req)

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println(total, err)
}

func TestUserCategoryRepository_Save(t *testing.T) {
	req := dbmodels.UserCategory{
		ID:    1,
		Code:  "op",
		Name:  "Ottopay",
		Notes: "Merchant Ottopay",
		Seq:   1,
	}
	err := InitUserCategoryRepository().Save(&req)

	log.Println(err)
}