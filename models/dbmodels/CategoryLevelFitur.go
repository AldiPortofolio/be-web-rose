package dbmodels

type CategoryLevelFitur struct {
	ID 					int64  `json:"id" gorm:"id"`
	UserCategoryId 		int64  `json:"userCategoryId"`
	LevelMerchantId 	int64  `json:"levelMerchantId"`
	FiturProductId 		int64  `json:"fiturProductId"`
	Status 				string  `json:"status"`
}

func (q *CategoryLevelFitur) TableName() string {
	return "public.category_level_fitur"
}