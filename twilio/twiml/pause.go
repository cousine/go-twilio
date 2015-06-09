package twiml

import ()

// Verb
// https://www.twilio.com/docs/api/twiml/pause
type Pause struct {
	Length uint `xml:", attr"`
}
