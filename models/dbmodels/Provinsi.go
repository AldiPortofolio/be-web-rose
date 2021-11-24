package dbmodels

import "time"

type Provinsi struct {
	Id  		int64 		`json:"id" gorm:"column:id"`
	Name        string 		`json:"name" gorm:"column:kecamatan_id"`
	RegionId      int64 	`json:"regionId" gorm:"column:region_id"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt   time.Time 	`json:"created_at" gorm:"column:created_at"`
}

// TableName ..
func (q *Provinsi) TableName() string {
	return "public.provinsi"
}
