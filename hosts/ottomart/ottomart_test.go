package ottomart

import (
	"encoding/json"
	"log"
	"rose-be-go/constants"
	"rose-be-go/models/ottomartmodels"
	"testing"
)


func TestOttomartHost_Send(t *testing.T) {

	res := ottomartmodels.ClearSessionRes{}

	data, err := InitOttomartHost().Send("", constants.OttomartClearSession)

	json.Unmarshal(data, &res)

	log.Println(res)
	log.Println(string(data))
	log.Println(err)
}