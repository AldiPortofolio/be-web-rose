package redis

type LimitTransactionDeposit struct {
	MaxLimit			int64 `json:"maxLimit"`
	MinLimit			int64 `json:"minLimit"`
	Category 			string `json:"category"`
	MemberType 			string `json:"memberType"`
}

