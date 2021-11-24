package dbmodels

import "time"

type BankList struct {
	ID                  int64     `json:"id"`
	Code                string    `json:"code"`
	FullName            string    `json:"fullName"`
	ShortName           string    `json:"shortName"`
	UrlImage            string    `json:"urlImage"`
	Status              string    `json:"status"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
	CreatedBy           string    `json:"createdBy"`
	UpdatedBy           string    `json:"updatedBy"`
	Seq                 int32     `json:"seq"`
}

func (t *BankList) TableName() string {
	return "public.bank_list"
}