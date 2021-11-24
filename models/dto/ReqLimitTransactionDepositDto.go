package dto

type ReqLimitTransactionDepositDto struct {
	MaxLimit			int64 `json:"maxLimit"`
	MinLimit			int64 `json:"minLimit"`
	Category 			string `json:"category"`
	MemberType 			string `json:"memberType"`
	Limit       		  int `json:"limit"`
	Page        		  int `json:"page"`
}