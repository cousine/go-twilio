package twiml

import ()

// Verb
// https://www.twilio.com/docs/api/twiml/queue
type Queue struct {
	Url    string `xml:",attr"`
	Method string `xml:",attr"`

	Value string `xml:",chardata"`
}

func (Queue) isTwiml() bool {
	return true
}
