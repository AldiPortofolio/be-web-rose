package models

import "time"

type TestRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ReqKafkaUploadNmid struct {
	FilePath 	string `json:"file_path"`
	Requestor 	string `json:"requestor"`
	DateTime 		time.Time `json:"date_time"`
}

type ReqUploadNmidData struct {
	Id int `json:"id" gorm:"PRIMARY KEY; column:id"`
	Date time.Time
}

type ReqDowloadFile struct {
	FilePath string `json:"filePath"`
}

//Pagination ...
type Pagination struct {
	Limit 	int `json:"limit"`
	Page 	int `json:"page"`
}

type ReqPhoneNumber struct {
	StorePhoneNumber string `json:"storePhoneNumber"`
} 

type ReqMerchantDetail struct {
	ID int64 `json:"id"`
} 


