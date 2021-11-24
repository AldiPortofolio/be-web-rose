package dbmodels

import "time"

type Kelurahan struct {
	Id  		int64 		`json:"id" gorm:"column:id"`
	KecamatanId int64 		`json:"kecamatanId" gorm:"column:kecamatan_id"`
	Name        string 		`json:"name" gorm:"column:kecamatan_id"`
	SubAreaId   int64 		`json:"subAreaId" gorm:"column:sub_area_id"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt   time.Time 	`json:"created_at" gorm:"column:created_at"`
}


func (q *Kelurahan) TableName() string {
	return "public.kelurahan"
}