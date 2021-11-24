package dbmodels

import "time"

type LogUpgradeFds struct {
	ID int64 `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
	Req string `json:"req"`
	Res string `json:"res"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedBy string `json:"createdBy"`
	UpdatedBy string `json:"updatedBy"`
	RetryAt time.Time `json:"retryAt"`
	RetryBy string `json:"retryBy"`
}

func (t *LogUpgradeFds) TableName() string {
	return "public.log_upgrade_fds"
}