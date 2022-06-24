package entities

type SendReqSt struct {
	To   string `json:"to"`
	Text string `json:"text"`
	Sync bool   `json:"sync"`
}

type SendRepSt struct {
	ID string `json:"id"`
}
