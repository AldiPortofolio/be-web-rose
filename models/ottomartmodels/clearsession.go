package ottomartmodels

// ClearSessionRes ..
type ClearSessionRes struct {
	Meta ClearSessionResMeta `json:"meta"`
}

// ClearSessionResMeta ..
type ClearSessionResMeta struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}