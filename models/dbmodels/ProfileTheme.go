package dbmodels

type ProfileTheme struct {
	ID 						int64  `json:"id" gorm:"id"`
	UserCategoryId 			int64  `json:"userCategoryId"`
	LevelMerchantId 		int64  `json:"levelMerchantId"`
	DashboardTopBackground 	string  `json:"dashboardTopBackground"`
	ThemeColor 				string  `json:"themeColor"`
	DashboardLogo       	string 	`json:"dashboardLogo"`
	DashboardText       	string 	`json:"dashboardText"`
	ProfileBackgroundImage  string  `json:"profileBackgroundImage"`
	Status       			string  `json:"status"`
	Url 					string `json:"url"`
	FontColor				string `json:"fontColor"`
}

func (q *ProfileTheme) TableName() string {
	return "public.profile_theme"
}