package twiml

import ()

// Verb
// https://www.twilio.com/docs/api/twiml/reject
type Reject struct {
	Reason string `xml:",attr"`
}

func (Reject) isTwiml() bool {
	return true
}
