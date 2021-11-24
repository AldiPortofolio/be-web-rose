package dto

type ReqLookupDto struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Descr       string `json:"descr"`
	LookupGroup string `json:"lookupGroup"`
	OrderNo     int64  `json:"orderNo"`
	Limit       int    `json:"limit"`
	Page        int    `json:"page"`
}