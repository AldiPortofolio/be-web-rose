package db

import (
	"encoding/json"
	"log"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"testing"
)

func TestFeatureProductRepository_FindByID(t *testing.T) {
	res, err := InitFeatureProductRepository().FindByID(1)

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println(err)
}

func TestFeatureProductRepository_Filter(t *testing.T) {
	req := dto.ReqFiturProductDto{
		ID:        1,
		ProductID: 0,
		Code:      "",
		Icon:      "",
		Name:      "",
		Notes:     "",
		Seq:       0,
		Limit:     10,
		Page:      1,
	}
	res, total, err := InitFeatureProductRepository().Filter(req)

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println(total, err)
}

func TestFeatureProductRepository_Save(t *testing.T) {
	req := dbmodels.FiturProduct{
		ID:        1,
		ProductID: 1,
		Code:      "pesan bar",
		Icon:      "img1",
		Name:      "Pesan Barang",
		Notes:     "notes",
		Seq:       2,
	}
	err := InitFeatureProductRepository().Save(&req)

	log.Println(err)
}