package viro

type Viro struct {
	ApiKey   string
	SenderID string
}

func NewViro(senderID, apiKey string) *Viro {
	return &Viro{
		ApiKey:   apiKey,
		SenderID: senderID,
	}
}
