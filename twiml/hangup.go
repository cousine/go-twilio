package twiml

import (
	"encoding/xml"
)

// Verb
// https://www.twilio.com/docs/api/twiml/Hangup
type Hangup struct {
	XMLName xml.Name `xml:"Hangup"`
}

func (Hangup) isTwiml() bool {
	return true
}
