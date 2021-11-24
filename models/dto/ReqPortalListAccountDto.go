package dto

// ReqPortalListAccountFilter ...
type ReqPortalListAccountFilter struct {
	Search			string `json:"search"`
	Filter_by		string `json:"filter_by"`
	Keyword			string `json:"keyword"`
	FilterAction 	string `json:"filter_action"`
	FilterMtype 	string `json:"filter_mtype"`
	Page 			int	`json:"page"`
	Limit 			int	`json:"limit"`
}

type ReqFilterOutlet struct {
	OutletName string `json:"outlet_name"`
	GroupName string `json:"group_name"`
	Page int `json:"page"`
	Limit int `json:"limit"`
}
