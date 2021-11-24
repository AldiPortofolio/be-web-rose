package dto

import "time"

type KategoriBisnisDto struct {
	ID					int64		`json:"id"`
	ActionType			int			`json:"actionType"`
	ApprovalStatus		int			`json:"approvalStatus"`
	LatestApproval 		time.Time	`json:"latestApproval"`
	LatestApprover		string		`json:"latestApprover"`
	LatestSuggestion	time.Time	`json:"latestSuggestion"`
	LatestSuggestor		string		`json:"latestSuggestor"`
	Status 				int64		`json:"status"`
	Version 			int			`json:"version"`
	Code 				string		`json:"code"`
	Name				string		`json:"name"`
	JenisUsahaId 		int64		`json:"jenisUsahaId"`
	JenisUsahaCode 		string		`json:"jenisUsahaCode"`
}
