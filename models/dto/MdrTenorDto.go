package dto

import "time"

type ReqMdrTenorDto struct {
	ID 			int64 `json:"id"`
	TenorCode	string `json:"tenorCode"`
	TenorName	string `json:"tenorName"`
	DokuTenorCode 	string `json:"dokuTenorCode"`
	Status 		string `json:"status"`
	Seq 		int32 `json:"seq"`
	UpdatedAt 	time.Time `json:"updatedAt"`
	UpdatedBy 	string `json:"updatedBy"`
	Limit       		  int `json:"limit"`
	Page        		  int `json:"page"`
}