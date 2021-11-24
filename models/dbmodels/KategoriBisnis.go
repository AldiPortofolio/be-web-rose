package dbmodels

import "time"


type KategoriBisnis struct {
	ID					int64		`json:"id" gorm:"id"`
	ActionType			int			`json:"actionType" gorm:"action_type"`
	ApprovalStatus		int			`json:"approvalStatus" gorm:"approval_status"`
	LatestApproval 		time.Time	`json:"latestApproval" gorm:"latest_approval"`
	LatestApprover		string		`json:"latestApprover" gorm:"latest_approver"`
	LatestSuggestion	time.Time	`json:"latestSuggestion" gorm:"latest_suggestion"`
	LatestSuggestor		string		`json:"latestSuggestor" gorm:"latest_suggestor"`
	Status 				int64		`json:"status" gorm:"status"`
	Version 			int			`json:"version" gorm:"version"`
	Code 				string		`json:"code" gorm:"code"`
	Name				string		`json:"name" gorm:"name"`
	JenisUsahaId 		int64		`json:"jenisUsahaId" gorm:"jenis_usaha_id"`
	JenisUsahaCode 		string		`json:"jenisUsahaCode" gorm:"jenis_usaha_code"`
}

func (q *KategoriBisnis) TableName() string {
	return "public.kategori_bisnis"
}