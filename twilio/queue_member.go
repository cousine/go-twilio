package twilio

import (
	"encoding/json"
	"fmt"
	"github.com/cousine/go-twilio/helpers"
	"io"
	"time"
)

const (
	QUEUE_MEMBER_ENDPOINT = "Members"
)

func queueMemberUri(accountSid string, queueSid string, callSid string) string {
	return fmt.Sprintf(
		"%v/%v/%v/%v/%v/%v",
		ACCOUNT_ENDPOINT,
		accountSid,
		QUEUE_ENDPOINT,
		queueSid,
		QUEUE_MEMBER_ENDPOINT,
		callSid,
	)
}

type QueueMember struct {
	CallSid      string    `json:"call_sid"`
	DateEnqueued time.Time `json:"date_enqueued"`
	WaitTime     int       `json:"wait_time"`
	Position     int       `json:"position"`

	AccountSid string
	QueueSid   string

	queue *Queue
}

func NewQueueMember(queue *Queue, accountSid string, queueSid string) *QueueMember {
	return &QueueMember{
		AccountSid: accountSid,
		QueueSid:   queueSid,
		queue:      queue,
	}
}

func (queueMember *QueueMember) FromJson(rawJson io.ReadCloser) error {
	jsonDecoder := json.NewDecoder(rawJson)
	err := jsonDecoder.Decode(queueMember)

	return err
}

func (queueMember *QueueMember) Dequeue(twimlUrl string) error {
	return queueMember.queue.dequeueMember(queueMember.CallSid, twimlUrl)
}

func (q *Queue) GetMember(callSid string) (qMember *QueueMember, err error) {
	qMember = NewQueueMember(q, q.AccountSid, q.Sid)

	uri := generateEndpointUrl(queueMemberUri(q.AccountSid, q.Sid, callSid))

	req, err := q.Client().generateRequest("GET", uri, map[string]string{})
	if err != nil {
		return
	}

	resp, err := q.Client().httpClient.Do(req)
	if err != nil {
		return
	}

	if helpers.ResponseHasError(resp) {
		return nil, NewTwilioError(resp.Body)
	}

	err = qMember.FromJson(resp.Body)
	if err != nil {
		return
	}

	return
}

func (q *Queue) GetFrontMember() (qMember *QueueMember, err error) {
	return q.GetMember("Front")
}

func (q *Queue) DequeueFront(twimlUrl string) error {
	return q.dequeueMember("Front", twimlUrl)
}

func (q *Queue) dequeueMember(callSid string, twimlUrl string) (err error) {
	uri := generateEndpointUrl(queueMemberUri(q.AccountSid, q.Sid, callSid))

	req, err := q.Client().generateRequest("POST", uri, map[string]string{"Url": twimlUrl})
	if err != nil {
		return
	}

	resp, err := q.Client().httpClient.Do(req)
	if err != nil {
		return
	}

	if helpers.ResponseHasError(resp) {
		return NewTwilioError(resp.Body)
	}

	return
}
