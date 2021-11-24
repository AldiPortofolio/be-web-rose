package blastnotifmodels

type ReqNotifAll struct {
	Desc   string `json:"desc"`
	Target string `json:"target"`
	Tilte  string `json:"tilte"`
}

type ResNotifAll struct {
	Rc  string `json:"rc"`
	Msg string `json:"msg"`
}