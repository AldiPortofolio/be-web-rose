package dto

type ReqMerchantAggregatorDto struct {
	ID 		int64 	`json:"id"`
	MidAggregator		  string	  `json:"midAggregator"`
	MidMerchant 		  []string 	  `json:"midMerchant"`
}

type ReqMerchantAggregatorUpload struct {
	FilePath string `json:"file_path"`
	User string `json:"user"`
	MidAggregator string `json:"mid_aggregator"`
}