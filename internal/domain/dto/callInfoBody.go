package dto

type CallInfoBody struct {
	CallId string `json:"callId,omitempty"`
	Dst    string `json:"dst,omitempty"`
}
