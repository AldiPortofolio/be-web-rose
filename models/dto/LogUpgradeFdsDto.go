package dto

type ReqLogUpgradeFdsDto struct {
	PhoneNumber string `json:"phoneNumber"`
	Status 	string `json:"status"`
	Limit       		  int `json:"limit"`
	Page        		  int `json:"page"`
}
