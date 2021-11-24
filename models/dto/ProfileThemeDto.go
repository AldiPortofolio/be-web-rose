package dto

type ReqProfileThemeDto struct {
	ID 						int64  `json:"id" gorm:"id"`
	UserCategoryId 			int64  `json:"userCategoryId"`
	LevelMerchantId 		int64  `json:"levelMerchantId"`
	DashboardTopBackground 	string  `json:"dashboardTopBackground"`
	ThemeColor 				string  `json:"themeColor"`
	DashboardLogo       	string 	`json:"dashboardLogo"`
	DashboardText       	string 	`json:"dashboardText"`
	ProfileBackgroundImage  string  `json:"profileBackgroundImage"`
	Status       			string  `json:"status"`
	Url						string `json:"url"`
	Limit       		  	int `json:"limit"`
	Page        		  	int `json:"page"`
	FontColor				string `json:"fontColor"`

}

type ResProfileThemeDto struct {
	ID 						int64  `json:"id" gorm:"id"`
	UserCategoryId 			int64  `json:"userCategoryId"`
	UserCategoryName 		string  `json:"userCategoryName" gorm:"user_category_name"`
	LevelMerchantId 		int64  `json:"levelMerchantId"`
	LevelMerchantName 		string  `json:"levelMerchantName" gorm:"level_merchant_name"`
	DashboardTopBackground 	string  `json:"dashboardTopBackground"`
	ThemeColor 				string  `json:"themeColor"`
	DashboardLogo       	string 	`json:"dashboardLogo"`
	DashboardText       	string 	`json:"dashboardText"`
	ProfileBackgroundImage  string  `json:"profileBackgroundImage"`
	Status       			string  `json:"status"`
	Url						string `json:"url"`
	FontColor				string `json:"fontColor"`


}
