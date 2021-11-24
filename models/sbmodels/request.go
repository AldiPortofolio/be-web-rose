package sbmodels

type ReqGenerateQr struct {
	Tid string `json:"tid"`
	Mid string `json:"mid"`
}

type ResGenerateQr struct {
	Rc string `json:"rc"`
	Msg string `json:"msg"`
	QrData string `json:"qrData"`
}
