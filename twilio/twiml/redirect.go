package twiml

import ()

// Verb
// https://www.twilio.com/docs/api/twiml/redirect
type Redirect struct {
	Method string `xml:", attr"`

	Value string `xml:", innerxml, omitempty"`
}
