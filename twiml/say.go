package twiml

import (
	"encoding/xml"
)

// Verb
// https://www.twilio.com/docs/api/twiml/say
type Say struct {
	XMLName xml.Name `xml:"Say"`

	Voice    string `xml:"voice,attr,omitempty"`
	Language string `xml:"language,attr,omitempty"`
	Loop     int    `xml:"loop,attr,omitempty"`

	Value string `xml:",chardata"`
}

func (Say) isTwiml() bool {
	return true
}
