package twiml

import ()

// Verb
// https://www.twilio.com/docs/api/twiml/sms
type Sms struct {
	To             string `xml:", attr"`
	From           string `xml:", attr"`
	Action         string `xml:", attr"`
	Method         string `xml:", attr"`
	StatusCallback string `xml:", attr"`

	Value string `xml:", innerxml, omitempty"`
}
