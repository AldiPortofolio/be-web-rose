package dto

type ReqLookupGroupDto struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}
