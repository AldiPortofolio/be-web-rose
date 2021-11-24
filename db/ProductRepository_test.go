package db

import (
	"encoding/json"
	"log"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"testing"
)

func TestProductRepository_FindByID(t *testing.T) {
	res, err := InitProductRepository().FindByID(1)

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println(err)
}

func TestProductRepository_Filter(t *testing.T) {
	req := dto.ReqProductDto{
		ID:    1,
		Code:  "",
		Name:  "",
		Title: "",
		Desc:  "",
		Notes: "",
		Seq:   0,
		Limit: 10,
		Page:  1,
	}
	res, total, err := InitProductRepository().Filter(req)

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println(total, err)
}

func TestProductRepository_Save(t *testing.T) {
	req := dbmodels.Product{
		ID:    1,
		Code:  "tgr",
		Name:  "Toko Grosir",
		Title: "Toko Grosir",
		Desc:  "Belanja stok",
		Notes: "notes toko grosir",
		Seq:   1,
	}
	err := InitProductRepository().Save(&req)

	log.Println(err)
}