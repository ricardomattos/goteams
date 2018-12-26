package goteams

type ReceivedMessage struct {
	Text           string    `json:"text"`
	TextFormat     string    `json:"textFormat"`
	Type           string    `json:"type"`
	Timestamp      string    `json:"timestamp"`
	LocalTimestamp string    `json:"localTimestamp"`
	ID             string    `json:"id"`
	ChannelID      string    `json:"channelId"`
	ServiceURL     string    `json:"serviceUrl"`
	From           struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		AadObjectID string `json:"aadObjectId"`
	} `json:"from"`
	Conversation struct {
		ConversationType string `json:"conversationType"`
		ID               string `json:"id"`
	} `json:"conversation"`
	Recipient struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"recipient"`
	Entities []struct {
		Locale   string `json:"locale"`
		Country  string `json:"country"`
		Platform string `json:"platform"`
		Type     string `json:"type"`
	} `json:"entities"`
	ChannelData struct {
		Tenant struct {
			ID string `json:"id"`
		} `json:"tenant"`
	} `json:"channelData"`
}

type SendMessage struct {
	Type string `json:"type"`
	From struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"from"`
	Conversation struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"conversation"`
	Recipient struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"recipient"`
	Text      string `json:"text"`
	ReplyToID string `json:"replyToId"`
}

type OAuth2Token struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	ExtExpiresIn int    `json:"ext_expires_in"`
	AccessToken  string `json:"access_token"`
}
