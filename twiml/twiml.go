package twiml

import (
	"encoding/xml"
	"io"
)

type TwimlElement interface {
	isTwiml() (ok bool)
}

// TwiML basic structure
// https://www.twilio.com/docs/api/twiml
type TwiML struct {
	XMLName xml.Name `xml:"Response"`

	SayList  []Say  `xml:"Say"`
	PlayList []Play `xml:"Play"`
	Dial     *Dial
	Record   *Record
	Gather   *Gather
	Sms      *Sms
	Hangup   *Hangup `xml:",omitempty"` // https://www.twilio.com/docs/api/twiml/hangup
	Enqueue  *Enqueue
	Leave    *Leave
	Queue    *Queue
	Redirect *Redirect
	Pause    *Pause
	Reject   bool `xml:",omitempty"` // https://www.twilio.com/docs/api/twiml/reject
	Message  *Message
}

type OrderedTwiML struct {
	XMLName xml.Name `xml:"Response"`

	ElementsList []TwimlElement
}

func (TwiML) isTwiml() bool {
	return true
}

func (OrderedTwiML) isTwiml() bool {
	return true
}

func NewTwiml() TwiML {
	return TwiML{
		Dial:    nil,
		Record:  nil,
		Gather:  nil,
		Sms:     nil,
		Queue:   nil,
		Pause:   nil,
		Message: nil,
	}
}

func WriteTwiml(writer io.Writer, twiml TwimlElement) error {
	ml, err := xml.Marshal(twiml)
	if err != nil {
		return err
	}

	writer.Write(ml)

	return nil
}
