package twiml

import ()

// Nouns
// https://www.twilio.com/docs/api/twiml/number
type Number struct {
	SendDigits           string `xml:", attr"`
	Url                  string `xml:", attr"`
	Method               string `xml:", attr"`
	StatusCallbackEvent  string `xml:", attr"`
	StatusCallback       string `xml:", attr"`
	StatusCallbackMethod string `xml:", attr"`

	Value string `xml:", innerxml, omitempty"`
}

// https://www.twilio.com/docs/api/twiml/client
type Client struct {
	Url                  string `xml:", attr"`
	Method               string `xml:", attr"`
	StatusCallbackEvent  string `xml:", attr"`
	StatusCallback       string `xml:", attr"`
	StatusCallbackMethod string `xml:", attr"`

	Value string `xml:", innerxml, omitempty"`
}

// https://www.twilio.com/docs/api/twiml/Sip
type Sip struct {
	Username             string `xml:", attr"`
	Password             string `xml:", attr"`
	Url                  string `xml:", attr"`
	Method               string `xml:", attr"`
	StatusCallbackEvent  string `xml:", attr"`
	StatusCallback       string `xml:", attr"`
	StatusCallbackMethod string `xml:", attr"`

	Value string `xml:", innerxml, omitempty"`
}

// https://www.twilio.com/docs/api/twiml/conference
type Conference struct {
	Muted                  bool   `xml:", attr"`
	Beep                   string `xml:", attr"`
	StartConferenceOnEnter bool   `xml:", attr"`
	EndConferenceOnExit    bool   `xml:", attr"`
	WaitUrl                string `xml:", attr"`
	WaitMethod             string `xml:", attr"`
	MaxParticipants        uint   `xml:", attr"`
	Record                 string `xml:", attr"`
	Trim                   string `xml:", attr"`
	EventCallbackUrl       string `xml:", attr"`

	Value string `xml:", innerxml, omitempty"`
}

// Verb
// https://www.twilio.com/docs/api/twiml/dial
type Dial struct {
	Action       string `xml:", attr"`
	Method       string `xml:", attr"`
	Timeout      uint   `xml:", attr"`
	HangupOnStar bool   `xml:", attr"`
	TimeLimit    uint   `xml:", attr"`
	CallerId     string `xml:", attr"`
	Record       string `xml:", attr"`
	Trim         string `xml:", attr"`

	Number
	Client
	Sip
	Conference
	Queue

	Value string `xml:", innerxml, omitempty"`
}
