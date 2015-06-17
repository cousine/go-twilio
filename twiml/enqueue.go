package twiml

import (
	"encoding/xml"
)

// Verb
// https://www.twilio.com/docs/api/twiml/enqueue
type Enqueue struct {
	XMLName xml.Name `xml:"Enqueue"`

	Action        string `xml:"action,attr,omitempty"`
	Method        string `xml:"method,attr,omitempty"`
	WaitUrl       string `xml:"waitUrl,attr,omitempty"`
	WaitUrlMethod string `xml:"waitUrlMethod,attr,omitempty"`
	WorkflowSid   string `xml:"workflowSid,attr,omitempty"`

	Value string `xml:",chardata"`
}

func (Enqueue) isTwiml() bool {
	return true
}
