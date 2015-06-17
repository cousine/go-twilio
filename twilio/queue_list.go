package twilio

import (
	"encoding/json"
	"github.com/cousine/go-twilio/helpers"
	"io"
)

type QueueList struct {
	PagingInformation

	AccountSid string
	client     *Client

	Queues []*Queue `json:"queues"`
}

func NewQueueList(client *Client, accountSid string) *QueueList {
	return &QueueList{
		AccountSid: accountSid,
		client:     client,
	}
}

func (queueList *QueueList) FromJson(rawJson io.ReadCloser) error {
	jsonDecoder := json.NewDecoder(rawJson)
	err := jsonDecoder.Decode(queueList)
	if err != nil {
		return err
	}

	for _, queue := range queueList.Queues {
		queue.AccountSid = queueList.AccountSid
		queue.SetClient(queueList.client)
	}

	return err
}

func (queueList *QueueList) NextPage(pageSize int) (Paginator, error) {
	queueL := &QueueList{}
	err := queueList.PagingInformation.Next(queueL, queueList.client, pageSize)

	return queueL, err
}

func (queueList *QueueList) PreviousPage(pageSize int) (Paginator, error) {
	queueL := &QueueList{}
	err := queueList.PagingInformation.Previous(queueL, queueList.client, pageSize)

	return queueL, err
}

func (client *Client) ListQueues(accountSid string) (queueList *QueueList, err error) {
	queueList = NewQueueList(client, accountSid)

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

	err = queueList.FromJson(resp.Body)
	if err != nil {
		return
	}

	return
}
