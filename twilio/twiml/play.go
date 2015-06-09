package twiml

import ()

// Verb
// https://www.twilio.com/docs/api/twiml/play
type Play struct {
	Loop   int    `xml:", attr"`
	Digits string `xml:", attr"`

	Value string `xml:", innerxml, omitempty"`
}
