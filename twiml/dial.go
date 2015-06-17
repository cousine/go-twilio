package twiml

import (
	"encoding/xml"
)

// Nouns
// https://www.twilio.com/docs/api/twiml/number
type Number struct {
	XMLName xml.Name `xml:"Number"`

	SendDigits           string `xml:"sendDigits,attr,omitempty"`
	Url                  string `xml:"url,attr,omitempty"`
	Method               string `xml:"method,attr,omitempty"`
	StatusCallbackEvent  string `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallback       string `xml:"statusCallback,attr,omitempty"`
	StatusCallbackMethod string `xml:"statusCallbackMethod,attr,omitempty"`

	Value string `xml:",chardata"`
}

// https://www.twilio.com/docs/api/twiml/client
type Client struct {
	XMLName xml.Name `xml:"Client"`

	Url                  string `xml:"url,attr,omitempty"`
	Method               string `xml:"method,attr,omitempty"`
	StatusCallbackEvent  string `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallback       string `xml:"statusCallback,attr,omitempty"`
	StatusCallbackMethod string `xml:"statusCallbackMethod,attr,omitempty"`

	Value string `xml:",chardata"`
}

// https://www.twilio.com/docs/api/twiml/Sip
type Sip struct {
	XMLName xml.Name `xml:"Sip"`

	Username             string `xml:"username,attr,omitempty"`
	Password             string `xml:"password,attr,omitempty"`
	Url                  string `xml:"url,attr,omitempty"`
	Method               string `xml:"method,attr,omitempty"`
	StatusCallbackEvent  string `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallback       string `xml:"statusCallback,attr,omitempty"`
	StatusCallbackMethod string `xml:"statusCallbackMethod,attr,omitempty"`

	Value string `xml:",chardata"`
}

// https://www.twilio.com/docs/api/twiml/conference
type Conference struct {
	XMLName xml.Name `xml:"Conference"`

	Muted                  bool   `xml:"muted,attr,omitempty"`
	Beep                   string `xml:"beep,attr,omitempty"`
	StartConferenceOnEnter bool   `xml:"startConferenceOnEnter,attr,omitempty"`
	EndConferenceOnExit    bool   `xml:"endConferenceOnExit,attr,omitempty"`
	WaitUrl                string `xml:"waitUrl,attr,omitempty"`
	WaitMethod             string `xml:"waitMethod,attr,omitempty"`
	MaxParticipants        uint   `xml:"maxParticipants,attr,omitempty"`
	Record                 string `xml:"record,attr,omitempty"`
	Trim                   string `xml:"trim,attr,omitempty"`
	EventCallbackUrl       string `xml:"eventCallbackUrl,attr,omitempty"`

	Value string `xml:",chardata"`
}

// Verb
// https://www.twilio.com/docs/api/twiml/dial
type Dial struct {
	XMLName xml.Name `xml:"Dial"`

	Action       string `xml:"action,attr,omitempty"`
	Method       string `xml:"method,attr,omitempty"`
	Timeout      uint   `xml:"timeout,attr,omitempty"`
	HangupOnStar bool   `xml:"hangupOnStar,attr,omitempty"`
	TimeLimit    uint   `xml:"timeLimit,attr,omitempty"`
	CallerId     string `xml:"callerId,attr,omitempty"`
	Record       string `xml:"record,attr,omitempty"`
	Trim         string `xml:"trim,attr,omitempty"`

	NumberList []Number
	ClientList []Client `xml:"Client"`
	Sip        *Sip
	Conference *Conference
	Queue      *Queue

	Value string `xml:",chardata"`
}

func (Dial) isTwiml() bool {
	return true
}
