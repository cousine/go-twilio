package twiml

import ()

// Verb
// https://www.twilio.com/docs/api/twiml/Message
type Message struct {
	To             string `xml:", attr"`
	From           string `xml:", attr"`
	Action         string `xml:", attr"`
	Method         string `xml:", attr"`
	StatusCallback string `xml:", attr"`

	Value string `xml:", innerxml, omitempty"`
	Body  string `xml:", omitempty"`
	Media string `xml:", omitempty"`
}
