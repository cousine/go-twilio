package twiml

import ()

// Verb
// https://www.twilio.com/docs/api/twiml/pause
type Pause struct {
	Length uint `xml:"length,attr"`
}

func (Pause) isTwiml() bool {
	return true
}
