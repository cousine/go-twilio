package twiml

import ()

// Verb
// https://www.twilio.com/docs/api/twiml/record
type Record struct {
	Action      string `xml:", attr"`
	Method      string `xml:", attr"`
	Timeout     uint   `xml:", attr"`
	FinishOnKey string `xml:", attr"`
	MaxLength   uint   `xml:", attr"`
	Transcribe  bool   `xml:", attr"`
	PlayBeep    bool   `xml:", attr"`
	Trim        string `xml:", attr"`

	Value string `xml:", innerxml, omitempty"`
}
