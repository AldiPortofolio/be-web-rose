package dto

type ReqMdrAggragtorDto struct {
	MdrAggregatorID		  int64		  `json:"mdrAggregatorId"`
	ActionType            int         `json:"actionType"`
	Status                int64       `json:"status"`
	GroupPartner		  string	  `json:"groupPartner"`
	MerchantCategory	  string 	  `json:"merchantCategory"`
	TransactionType 	  string 	  `json:"transactionType"`
	MdrType				  string 	  `json:"mdr_type"`
	Mdr 			   	  float64	  `json:"mdr"`
	Notes 				  string 	  `json:"notes"`
	MidPartner 			  string 	  `json:"midPartner"`
	MidMerchant 		  string	  `json:"midMerchant"`
	Limit       		  int `json:"limit"`
	Page        		  int `json:"page"`
}

