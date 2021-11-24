package dto

type ResVersionAppDto struct {
	Ottomart 	string `json:"ottomart"`
	Sfa 		string `json:"sfa"`
	Nfc 		string `json:"nfc"`
	Indomarco 	string `json:"indomarco"`
}

type ReqUpdateVersionAppDto struct {
	AppName 	string `json:"appName"`
	Version		string `json:"version"`
}