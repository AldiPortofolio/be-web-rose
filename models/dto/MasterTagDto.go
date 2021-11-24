package dto

type ReqMasterTagDto struct {
	ID          int64  `json:"id"`
	Code        string `json:"code" example:"00001"`
	Name        string `json:"name" example:"sampo"`
	Description string `json:"description" example:"nama nama ikan"`
	Status      string `json:"status" example:"active"`
	Limit       int    `json:"limit" example:"10"`
	Page        int    `json:"page" example:"1"`
}
