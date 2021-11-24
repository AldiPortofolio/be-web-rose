package dbmodels

import "time"

type Dati2 struct {
	Id  		int64 		`json:"id" gorm:"column:id"`
	ProvinsiId  int64 		`json:"provinsiId" gorm:"column:provinsi_id"`
	Name        string 		`json:"name" gorm:"column:kecamatan_id"`
	BranchId      int64 		`json:"branchId" gorm:"column:branch_id"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt   time.Time 	`json:"created_at" gorm:"column:created_at"`
}

func (q *Dati2) TableName() string {
	return "public.dati2"
}
