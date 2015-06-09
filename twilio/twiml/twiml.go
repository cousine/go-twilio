package twiml

import (
	"encoding/xml"
)

// TwiML basic structure
// https://www.twilio.com/docs/api/twiml
type TwiML struct {
	XMLName xml.Name `xml:"Response"`

	Say
	Play
	Dial
	Record
	Gather
	Sms
	Hangup bool `xml:", omitempty"` // https://www.twilio.com/docs/api/twiml/hangup
	Queue
	Redirect
	Pause
	Reject `xml:", omitempty"`
	Message
}
