package models


// Response ...
//type Response struct {
//	ErrCode string `json:"errCode"`
//	ErrDesc string `json:"errDesc"`
//}


// ResponseData ...
type Response struct {
	ErrCode string `json:"errCode"`
	ErrDesc string `json:"errDesc"`
	Data 	interface{} `json:"data,omitempty"`
	TotalData 	int `json:"totalData"`
	Contents  	interface{} `json:"contents"`

}

// ResReportFinished ...
//type ResReportFinished struct {
//	ErrCode   string      `json:"errCode"`
//	ErrDesc   string      `json:"errDesc"`
//	TotalData int         `json:"totalData"`
//	Contents  interface{} `json:"contents"`
//}

// ResponseReportIssuer ...
//type ResUploadNmidData struct {
//	ErrCode 	string 		`json:"errCode"`
//	ErrDesc 	string 		`json:"errDesc"`
//	TotalData 	int         `json:"totalData"`
//	Contents  	interface{} `json:"contents"`
//}