package dbmodels

import "time"

type MdrTenor struct {
	ID 			int64 `json:"id"`
	TenorCode	string `json:"tenorCode"`
	TenorName	string `json:"tenorName"`
	DokuTenorCode string `json:"dokuTenorCode"`
	Status 		string `json:"status"`
	Seq 		int32 `json:"seq"`
	UpdatedAt 	time.Time `json:"updatedAt"`
	UpdatedBy 	string `json:"updatedBy"`

}

func (q *MdrTenor) TableName() string {
	return "public.mdr_tenor"
}