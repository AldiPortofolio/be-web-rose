package dto

type ReqFilterDto struct {
	Name 				string `json:"name"`
	MidAggregator		string	  `json:"midAggregator"`
	MidMerchant 		[]string 	  `json:"midMerchant"`
	MidFilter	        string   `json:"midFilter"`
	MpanMerchant 		string		`json:"mpanMerchant"`
	NmidMerchant        string  `json:"nmidMerchant"`
	Action              int `json:"action"`
	Mid 				string `json:"mid"`
	Status 				string `json:"status"`
	AgentID 			string `json:"agentId"`
	RequestDate 		string `json:"requestDate"`
	InstallDate 		string `json:"installDate"`
	Page 				int `json:"page"`
	Limit 				int `json:"limit"`
	Id 					int64 `json:"id"`
} 
