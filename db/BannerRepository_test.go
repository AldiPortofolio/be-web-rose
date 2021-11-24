package db

import (
	"encoding/json"
	"log"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"testing"
)

func TestBannerRepository_FindByID(t *testing.T) {
	res, err := InitBannerRepository().FindByID(1)

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println(err)
}

func TestBannerRepository_Filter(t *testing.T) {
	req := dto.ReqBannerDto{
		ID:              1,
		UserCategoryId:  1,
		LevelMerchantId: 2,
		AdsImage:        "ads",
		AdsLink:         "ads",
		Seq:             "2",
		Status:          "Y",
		Limit:           10,
		Page:            1,
	}
	res, total, err := InitBannerRepository().Filter(req)

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println(total, err)
}

func TestBannerRepository_Save(t *testing.T) {
	req := dbmodels.Banner{
		ID:              4,
		UserCategoryId:  1,
		LevelMerchantId: 2,
		AdsImage:        "image edit",
		AdsLink:         "link",
		Seq:             "5",
		Status:          "Y",
	}
	err := InitBannerRepository().Save(&req)

	log.Println(err)
}