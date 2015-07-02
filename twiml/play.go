package twiml

import (
	"encoding/xml"
)

// Verb
// https://www.twilio.com/docs/api/twiml/play
type Play struct {
	XMLName xml.Name `xml:"Play"`

	Loop   int    `xml:"loop,attr"`
	Digits string `xml:"digits,attr,omitempty"`

	Value string `xml:",chardata"`
}

func (Play) isTwiml() bool {
	return true
}
