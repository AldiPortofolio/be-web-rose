package dto

type ReqBannerDto struct {
	ID              int64  `json:"id" gorm:"id"`
	UserCategoryId  int64  `json:"userCategoryId"`
	LevelMerchantId int64  `json:"levelMerchantId"`
	Name            string `json:"name"`
	AdsImage        string `json:"adsImage"`
	AdsLink         string `json:"adsLink"`
	Seq             string `json:"seq"`
	Status          string `json:"status"`
	Limit           int    `json:"limit"`
	Page            int    `json:"page"`
	BannerName      string `json:"bannerName"`
	DetailBanner    string `json:"bannerDetail"`
}

type ResBannerDto struct {
	ID                int64  `json:"id" gorm:"id"`
	UserCategoryId    int64  `json:"userCategoryId" gorm:"user_category_id"`
	UserCategoryName  string `json:"userCategoryName" gorm:"user_category_name"`
	LevelMerchantId   int64  `json:"levelMerchantId" gorm:"level_merchant_id"`
	LevelMerchantName string `json:"levelMerchantName" gorm:"level_merchant_name"`
	Name              string `json:"name"`
	AdsImage          string `json:"adsImage"`
	AdsLink           string `json:"adsLink"`
	Seq               string `json:"seq"`
	Status            string `json:"status"`
	BannerName        string `json:"bannerName" gorm:"banner_name"`
	DetailBanner      string `json:"bannerDetail" gorm:"banner_detail"`
}
