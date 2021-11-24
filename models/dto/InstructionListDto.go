package dto

type ReqInstructionListDto struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	Sequence    int    `json:"sequence"`
	Limit       int    `json:"limit"`
	Page        int    `json:"page"`
}

type ResInstructionListDto struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	Sequence    int    `json:"sequence"`
}
