package models

type User struct {
	ID                  int64    `json:"id"`
	Name                string `json:"name"`
	Email               string `json:"email"`
	Status              int    `json:"status"`
	MobileNo            string `json:"mobileNo"`
	ApprovalStatus      int    `json:"approvalStatus"`
	LastLoginDate       int64  `json:"lastLoginDate"`
	LastLogoutDate      int64  `json:"lastLogoutDate"`
	PasswordExpiredDate int64  `json:"passwordExpiredDate"`
	IsLock              int    `json:"isLock"`
	MustChangePassword  int    `json:"mustChangePassword"`
	Role                struct {
		ID       int    `json:"id"`
		Code     string `json:"code"`
		Name     string `json:"name"`
		Type     int    `json:"type"`
		IsSuper  int    `json:"isSuper"`
		Statuses int    `json:"statuses"`
	} `json:"role"`
	Version           int    `json:"version"`
	Type              string `json:"type"`
	NumbOfFailedLogin int    `json:"numbOfFailedLogin"`
	Islogged          int    `json:"islogged"`
}