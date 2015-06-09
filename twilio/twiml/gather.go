package twiml

import ()

// Verb
// https://www.twilio.com/docs/api/twiml/gather
type Gather struct {
	Action      string `xml:", attr"`
	Method      string `xml:", attr"`
	Timeout     uint   `xml:", attr"`
	FinishOnKey string `xml:", attr"`
	NumDigits   uint   `xml:", attr"`

	Say
	Play
	Pause
}
