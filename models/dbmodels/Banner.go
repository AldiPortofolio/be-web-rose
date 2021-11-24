package dbmodels

type Banner struct {
	ID              int64  `json:"id" gorm:"id"`
	UserCategoryId  int64  `json:"userCategoryId"`
	LevelMerchantId int64  `json:"levelMerchantId"`
	Name            string `json:"name"`
	AdsImage        string `json:"adsImage"`
	AdsLink         string `json:"adsLink"`
	Seq             string `json:"seq"`
	Status          string `json:"status"`
	BannerName      string `json:"banner_name"`
	DetailBanner    string `json:"banner_detail"`
}

func (q *Banner) TableName() string {
	return "public.banner"
}
