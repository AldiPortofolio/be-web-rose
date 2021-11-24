package dbmodels

import "time"

type AppUser struct {
	ID                  int64     `json:"id"`
	// ApprovalStatus      int       `json:"approvalStatus" gorm:"column:approval_status"`
	// Email               string    `json:"email" gorm:"column:email"`
	// FullName            string    `json:"fullName" gorm:"column:full_name"`
	// IsLock              int       `json:"isLock" gorm:"column:is_lock"`
	// LastLoginDate       time.Time `json:"lastLoginDate" gorm:"column:last_login_date"`
	// LastLoginFrom       string    `json:"lastLoginFrom" gorm:"column:last_login_from"`
	// LastLogoutDate      time.Time `json:"lastLogoutDate" gorm:"column:last_logout_date"`
	// MobileNo            string    `json:"mobileNo" gorm:"column:mobile_no"`
	// MustChangePassword  int       `json:"mustChangePassword" gorm:"column:must_change_password"`
	UserName            string    `json:"username" gorm:"column:user_name"`
	Password            string    `json:"password" gorm:"column:password"`
	PwdExpiredDate      time.Time `json:"pwdExpiredDate" gorm:"column:pwd_expired_date"`
	// Status              int64     `json:"status" gorm:"column:status"`
	// Type                int       `json:"type" gorm:"column:type"`
	// Version             int       `json:"version" gorm:"column:version"`
	// AreaId              int64     `json:"areaId" gorm:"column:area_id"`
	// RoleId              int64     `json:"roleId" gorm:"column:role_id"`
	// NumberOfFailedLogin int       `json:"numberOfFailedLogin" gorm:"column:number_of_failed_login"`
	// Logged              int       `json:"logged" gorm:"column:logged"`
}

func (t *AppUser) TableName() string {
	return "public.app_user"
}
