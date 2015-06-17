package twiml

import (
	"encoding/xml"
)

// Verb
// https://www.twilio.com/docs/api/twiml/leave
type Leave struct {
	XMLName xml.Name `xml:"Leave"`
}

func (Leave) isTwiml() bool {
	return true
}
