package dto

type ReqMerchantDto struct {
	MerchantGroupID 	int64 `json:"merchantGroupId"`
	StoreName 		string `json:"storeName"`
	Limit       		  int `json:"limit"`
	Page        		  int `json:"page"`
}

type ReqMerchantGroupDto struct {
	MerchantGroupID 		int64 	`json:"merchantGroupId"`
	StoreName 				string 	`json:"storeName"`
	PortalStatus    		string 	`json:"portalStatus"`
	Limit       		  	int 	`json:"limit"`
	Page        		  	int 	`json:"page"`
}

type ReqDashboardMerchantDto struct {
	Name         string `json:"name"`
	PortalStatus string `json:"portalStatus"`
	TipeMerchant string `json:"tipeMerchant"`
	Limit        int    `json:"limit"`
	Page         int    `json:"page"`
}

type ResEmailDto struct {
	Email 	string `json:"email"`
}
