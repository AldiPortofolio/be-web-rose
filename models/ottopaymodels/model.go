package ottopaymodels

type ReqPushNotif struct {
	CustAccount string `json:"custAccount"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Target string `json:"target"`
}
