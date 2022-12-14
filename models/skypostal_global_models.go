package models

type SkypostalUserInfo struct {
	UserCode int
	AppKey   string
	UserKey  string
}

type SkypostalAdditionalInfo struct {
	Server   serverInfo
	Internal serverInternalInfo
}

type serverInfo struct {
	ServerId   string
	ServerTime int
}

type serverInternalInfo struct {
	ServerId   string
	ProcessId  string
	ServerTime int //???
}

type SkypostalError struct {
	ErrorCode        string
	ErrorDescription string
	ErrorLocation    string
	SystemError      bool
}
