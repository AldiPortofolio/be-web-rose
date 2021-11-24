package db

import (
	"encoding/json"
	"log"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"testing"
)

func TestProfileThemeRepository_FindByID(t *testing.T) {
	res, err := InitProfileThemeRepository().FindByID(1)

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println(err)
}

func TestProfileThemeRepository_Filter(t *testing.T) {
	req := dto.ReqProfileThemeDto{
		ID:                     1,
		UserCategoryId:         0,
		LevelMerchantId:        0,
		DashboardTopBackground: "",
		ThemeColor:             "",
		DashboardLogo:          "",
		DashboardText:          "",
		ProfileBackgroundImage: "",
		Status:                 "",
		Limit:                  10,
		Page:                   1,
	}
	res, total, err := InitProfileThemeRepository().Filter(req)

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println(total, err)
}

func TestProfileThemeRepository_Save(t *testing.T) {
	req := dbmodels.ProfileTheme{
		ID:                     1,
		UserCategoryId:         0,
		LevelMerchantId:        0,
		DashboardTopBackground: "",
		ThemeColor:             "",
		DashboardLogo:          "",
		DashboardText:          "",
		ProfileBackgroundImage: "",
		Status:                 "",
	}
	err := InitProfileThemeRepository().Save(&req)

	log.Println(err)
}