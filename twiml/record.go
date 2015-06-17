package twiml

import (
	"encoding/xml"
)

// Verb
// https://www.twilio.com/docs/api/twiml/record
type Record struct {
	XMLName xml.Name `xml:"Record"`

	Action      string `xml:"action,attr,omitempty"`
	Method      string `xml:"method,attr,omitempty"`
	Timeout     uint   `xml:"timeout,attr,omitempty"`
	FinishOnKey string `xml:"finishOnKey,attr,omitempty"`
	MaxLength   uint   `xml:"maxLength,attr,omitempty"`
	Transcribe  bool   `xml:"transcribe,attr,omitempty"`
	PlayBeep    bool   `xml:"playBeep,attr,omitempty"`
	Trim        string `xml:"trip,attr,omitempty"`

	Value string `xml:",chardata"`
}

func (Record) isTwiml() bool {
	return true
}
