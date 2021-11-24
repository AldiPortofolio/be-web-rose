package db

import (
	"encoding/json"
	"log"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"testing"
)

func TestCategoryLevelFeatureRepository_FindByID(t *testing.T) {
	res, err := InitCategoryLevelFeatureRepository().FindByID(1)

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println(err)
}

func TestCategoryLevelFeatureRepository_Filter(t *testing.T) {
	req := dto.ReqCategoryLevelFeatureDto{
		ID:              1,
		UserCategoryId:  1,
		LevelMerchantId: 2,
		FiturProductId:  0,
		Status:          "",
		Limit:           10,
		Page:            1,
	}
	res, total, err := InitCategoryLevelFeatureRepository().Filter(req)

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println(total, err)
}

func TestCategoryLevelFeatureRepository_Save(t *testing.T) {
	req := dbmodels.CategoryLevelFitur{
		ID:              1,
		UserCategoryId:  1,
		LevelMerchantId: 2,
		FiturProductId:  2,
		Status:          "N",
	}
	err := InitCategoryLevelFeatureRepository().Save(&req)

	log.Println(err)
}