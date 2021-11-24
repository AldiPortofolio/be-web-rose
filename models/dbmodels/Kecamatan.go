package dbmodels

import "time"

type Kecamatan struct {
	Id  		int64 		`json:"id" gorm:"column:id"`
	Dati2Id     int64 		`json:"dati2Id" gorm:"column:dati2_id"`
	Name        string 		`json:"name" gorm:"column:kecamatan_id"`
	AreaId      int64 		`json:"areaId" gorm:"column:area_id"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt   time.Time 	`json:"created_at" gorm:"column:created_at"`
}


func (q *Kecamatan) TableName() string {
	return "public.kecamatan"
}