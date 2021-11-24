package dto

type ReqCategoryLevelFeatureDto struct {
	ID 					int64  `json:"id" gorm:"id"`
	UserCategoryId 		int64  `json:"userCategoryId"`
	LevelMerchantId 	int64  `json:"levelMerchantId"`
	FiturProductId 		int64  `json:"fiturProductId"`
	Status 				string  `json:"status"`
	Limit       		  int `json:"limit"`
	Page        		  int `json:"page"`
}

type ResCategoryLevelFeatureDto struct {
	ID 					int64  `json:"id" gorm:"id"`
	UserCategoryId 		int64  `json:"userCategoryId"`
	UserCategoryName 	string  `json:"userCategoryName" gorm:"user_category_name"`
	LevelMerchantId 	int64  `json:"levelMerchantId"`
	LevelMerchantName 	string  `json:"levelMerchantName" gorm:"level_merchant_name"`
	FiturProductId 		int64  `json:"fiturProductId"`
	FiturProductName 	string  `json:"fiturProductName" gorm:"fitur_product_name"`
	Status 				string  `json:"status"`
}
