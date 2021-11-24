package dto

type ReqUserCategoryDto struct {
	ID    int64  `json:"id" gorm:"id"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Logo  string `json:"logo"`
	AppID string `json:"appId"`
	Notes string `json:"notes"`
	Seq   int    `json:"seq"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}

type ResUserCategoryDto struct {
	ID      int64  `json:"id" gorm:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	Logo    string `json:"logo"`
	AppID   string `json:"appId"`
	AppName string `json:"appName"`
	Notes   string `json:"notes"`
	Seq     int    `json:"seq"`
}

type ResUserCategoryDropdownLisDto struct {
	ID   int64  `json:"id" gorm:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
