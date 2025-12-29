package dto

type InitCall struct {
	Endpoint       string `json:"endpoint" binding:"required"`
	Extension      string `json:"extension" binding:"required"`
	Context        string `json:"context"`
	Priority       int    `json:"priority"`
	Label          string `json:"label"`
	App            string `json:"app"`
	AppArgs        string `json:"appArgs"`
	CallerId       string `json:"callerId"`
	Timeout        int    `json:"timeout"`
	Variables      string `json:"variables"`
	ChannelId      string `json:"channelId"`
	OtherChannelId string `json:"otherChannelId"`
	Originator     string `json:"originator"`
	Formats        string `json:"formats"`
}
