package entities

type DevinoSendReqSt struct {
	Messages []DevinoSendReqMessageSt `json:"messages"`
}

type DevinoSendReqMessageSt struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Text     string `json:"text"`
	Validity int64  `json:"validity"`
	Priority int    `json:"priority"` // 0,1,2,3. 3 is high
}

type DevinoSendRepSt struct {
	Result []DevinoSendRepResultSt `json:"result"`
}

type DevinoSendRepResultSt struct {
	Code        string                 `json:"code"`
	MessageId   string                 `json:"messageId"`
	Description string                 `json:"description"`
	Reasons     []DevinoRejectReasonSt `json:"reasons"`
}

type DevinoRejectReasonSt struct {
	Key string `json:"key"`
	Ref string `json:"ref"`
}
