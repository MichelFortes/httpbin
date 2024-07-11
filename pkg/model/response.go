package model

type ResponseBody struct {
	ServiceId   string              `json:"serviceId"`
	RemoteAddr  string              `json:"remoteAddr"`
	Method      string              `json:"method"`
	Path        string              `json:"path"`
	QueryParams map[string][]string `json:"queryParams"`
	Headers     map[string][]string `json:"headers"`
}
