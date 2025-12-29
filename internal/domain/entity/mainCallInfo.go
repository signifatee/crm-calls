package entity

type MainCallInfo struct {
	CallId         string `json:"callId"`
	Trunk          string `json:"trunk"`
	FinishBy       int    `json:"finishBy"`
	Status         string `json:"status"`
	FinishReason   string `json:"finishReason"`
	RecordUrl      string `json:"recordUrl"`
	CallDuration   int    `json:"duration"`
	AnswerDuration int    `json:"answerDuration"`
}

type Extra struct {
	HangupCause  int    `json:"hangupcause"`
	HangupSource string `json:"hangupsource"`
	Dialstatus   string `json:"dialstatus"`
}
