package dto

type ReqAppUserDto struct {
	Username	string `json:"username"`
	OldPass		string `json:"oldPass"`
	NewPass		string `json:"newPass"`
	
}

