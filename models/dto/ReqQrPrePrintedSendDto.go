package dto


type ReqQrPrePrintedSendDto struct {
	Topic           string `json:"topic"`
	TotalReq 		int32 `json:"totalReq"`
	City 		   	string `json:"city"`
	PostalCode		string `json:"postalCode"`
	Mcc 			string `json:"mcc"`
	User 			string `json:"user"`
	Key 			string `json:"key"`
}
