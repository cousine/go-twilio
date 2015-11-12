package twiml

import (
	"encoding/xml"
)

// Verb
// https://www.twilio.com/docs/api/twiml/redirect
type Redirect struct {
	XMLName xml.Name `xml:"Redirect"`

	Method string `xml:"method,attr"`

	Value string `xml:",chardata"`
}

func (Redirect) isTwiml() bool {
	return true
}
