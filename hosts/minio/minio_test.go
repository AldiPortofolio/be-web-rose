package minio

import (
	"encoding/json"
	"log"
	"rose-be-go/constants"
	"rose-be-go/models/ottomartmodels"
	"testing"
)


func TestMiniHost_Send(t *testing.T) {

	res := ottomartmodels.ClearSessionRes{}

	data, err := InitMinioHost().Send("", constants.MinioUpload)

	json.Unmarshal(data, &res)

	log.Println(res)
	log.Println(string(data))
	log.Println(err)
}
