package twilio

import (
	"encoding/json"
	"github.com/cousine/go-twilio/helpers"
	"io"
)

type ConferenceList struct {
	PagingInformation

	AccountSid string
	client     *Client

	Conferences []*Conference `json:"conferences"`
}

func NewConferenceList(client *Client, accountSid string) *ConferenceList {
	return &ConferenceList{
		AccountSid: accountSid,
		client:     client,
	}
}

func (conferenceList *ConferenceList) FromJson(rawJson io.ReadCloser) error {
	jsonDecoder := json.NewDecoder(rawJson)
	err := jsonDecoder.Decode(conferenceList)
	if err != nil {
		return err
	}

	for _, conference := range conferenceList.Conferences {
		conference.AccountSid = conferenceList.AccountSid
		conference.SetClient(conferenceList.client)
	}

	return err
}

func (conferenceList *ConferenceList) NextPage(pageSize int) (Paginator, error) {
	queueL := &ConferenceList{}
	err := conferenceList.PagingInformation.Next(queueL, conferenceList.client, pageSize)

	return queueL, err
}

func (conferenceList *ConferenceList) PreviousPage(pageSize int) (Paginator, error) {
	queueL := &ConferenceList{}
	err := conferenceList.PagingInformation.Previous(queueL, conferenceList.client, pageSize)

	return queueL, err
}

func (client *Client) ListConferences(accountSid string) (conferenceList *ConferenceList, err error) {
	conferenceList = NewConferenceList(client, accountSid)

	uri := generateEndpointUrl(queuesUri(accountSid))

	req, err := client.generateRequest("GET", uri, map[string]string{})
	if err != nil {
		return
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return
	}

	if helpers.ResponseHasError(resp) {
		return nil, NewTwilioError(resp.Body)
	}

	err = conferenceList.FromJson(resp.Body)
	if err != nil {
		return
	}

	return
}
