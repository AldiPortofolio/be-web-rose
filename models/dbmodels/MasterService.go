package dbmodels

import "time"

// MasterService ..
type MasterService struct {
	ID        int64     `json:"id" gorm:"column:id;"`
	Name      string    `json:"name" gorm:"column:name;"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;"`
	UpdatedBy string    `json:"updatedBy" gorm:"column:updated_by;"`
}

// TableName ..
func (q *MasterService) TableName() string {
	return "public.master_service"
}
