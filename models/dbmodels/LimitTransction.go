package dbmodels

import "time"

type LimitTransaction struct {
	ID             int64     `json:"id"`
	UserCategory   string    `json:"userCategoryId" gorm:"column:user_category"`
	LevelMerchant  string    `json:"levelMerchant" gorm:"column:level_merchant"`
	LimitFreq      int64     `json:"limitFreq"`
	MinLimitAmount int64     `json:"minLimitAmount" gorm:"column:min_limit_amount"`
	LimitAmount    int64     `json:"limitAmount"`
	TimeFrame      string    `json:"timeFrame"`
	FeatureProduct string    `json:"featureProduct"`
	CreatedAt      time.Time `json:"createdAt" gorm:"column:created_at;"`
	CreatedBy      string    `json:"createdBy" gorm:"column:created_by;"`
	UpdatedAt      time.Time `json:"updatedAt" gorm:"column:updated_at;"`
	UpdatedBy      string    `json:"updatedBy" gorm:"column:updated_by;"`
}

func (q *LimitTransaction) TableName() string {
	return "public.limit_transaction"
}
