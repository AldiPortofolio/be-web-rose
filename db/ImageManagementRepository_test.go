package db

import (
	"encoding/json"
	"log"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"testing"
)

func TestImageManagementRepository_FindByID(t *testing.T) {
	res, err := InitImageManagementRepository().FindByID(1)

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println(err)
}

func TestImageManagementRepository_Filter(t *testing.T) {
	req := dto.ReqImageManagementDto{
		ID:     1,
		Name:   "",
		URL:    "",
		Notes:  "",
		Images: "",
		Limit:  10,
		Page:   1,
	}
	res, total, err := InitImageManagementRepository().Filter(req)

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println(total, err)
}

func TestImageManagementRepository_Save(t *testing.T) {
	req := dbmodels.ImageManagement{
		ID:    1,
		Name:  "test upload 1",
		URL:   "ttp://13.228.25.85:9000/rose/test-upload-1-27131847.jpeg",
		Notes: "test notes upload 1",
	}
	err := InitImageManagementRepository().Save(&req)

	log.Println(err)
}