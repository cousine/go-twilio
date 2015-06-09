package twiml

import ()

// Verb
// https://www.twilio.com/docs/api/twiml/say
type Say struct {
	Voice    string `xml:", attr"`
	Language string `xml:", attr"`
	Loop     int    `xml:", attr"`

	Value string `xml:", innerxml, omitempty"`
}
