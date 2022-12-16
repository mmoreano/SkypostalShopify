package models

type SkypostalUserInfo struct {
	UserCode int
	AppKey   string
	UserKey  string
}

type SkypostalAdditionalInfo struct {
	Server   serverInfo         `json:"server"`
	Internal serverInternalInfo `json:"internal"`
}

type serverInfo struct {
	ServerId   string `json:"server_id"`
	ServerTime int    `json:"server_time"`
}

type serverInternalInfo struct {
	ServerId   string
	ProcessId  string
	ServerTime int //???
}

type SkypostalError struct {
	ErrorCode        string `json:"error_code"`
	ErrorDescription string `json:"error_description"`
	ErrorLocation    string `json:"error_location"`
	SystemError      bool   `json:"system_error"`
}
