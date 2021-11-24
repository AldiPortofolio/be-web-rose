package db

import (
	"encoding/json"
	"log"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"testing"
)

func TestInstructionListRepository_FindByID(t *testing.T) {
	res, err := InitInstructionListRepository().FindByID(1)

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println(err)
}

func TestInstructionListRepository_Filter(t *testing.T) {
	req := dto.ReqInstructionListDto{
		Limit:           10,
		Page:            1,
	}
	res, total, err := InitInstructionListRepository().Filter(req)

	resByte,_ := json.Marshal(res)

	log.Println("res -->", string(resByte))
	log.Println(total, err)
}

func TestInstructionListRepository_Save(t *testing.T) {
	req := dbmodels.InstructionList{
		ID:              0,
		Title:           "test",
		Logo:            "http:/test",
		Description:     "description test",
		Sequence:        5,
	}
	err := InitInstructionListRepository().Save(&req)

	log.Println(err)
}