package twilio

import (
	"encoding/json"
	"fmt"
	"github.com/cousine/go-twilio/helpers"
	"io"
	"strconv"
)

const (
	QUEUE_ENDPOINT = "Queues"
)

type Queue struct {
	Sid             string `json:"sid"`
	FriendlyName    string `json:"friendly_name"`
	CurrentSize     int    `json:"current_size"`
	MaxSize         int    `json:"max_size"`
	AverageWaitTime int    `json:"average_wait_time"`

	AccountSid string

	client *Client
}

func NewQueue(client *Client, accountSid string) *Queue {
	return &Queue{
		AccountSid: accountSid,
		client:     client,
	}
}

func queueUri(accountSid string, queueSid string) string {
	return fmt.Sprintf("%v/%v/%v/%v", ACCOUNT_ENDPOINT, accountSid, QUEUE_ENDPOINT, queueSid)
}

func queuesUri(accountSid string) string {
	return fmt.Sprintf("%v/%v/%v", ACCOUNT_ENDPOINT, accountSid, QUEUE_ENDPOINT)
}

func (queue *Queue) Client() *Client {
	return queue.client
}

func (queue *Queue) SetClient(client *Client) {
	queue.client = client
}

func (queue *Queue) FromJson(rawJson io.ReadCloser) error {
	jsonDecoder := json.NewDecoder(rawJson)
	err := jsonDecoder.Decode(queue)

	return err
}

func (queue *Queue) Update(friendlyName string, maxSize int) (err error) {
	queue.FriendlyName = friendlyName
	queue.MaxSize = maxSize

	return queue.client.updateQueue(queue)
}

func (queue *Queue) Delete() (err error) {
	return queue.client.deleteQueue(queue)
}

func (client *Client) GetQueue(accountSid string, queueSid string) (queue *Queue, err error) {
	queue = NewQueue(client, accountSid)

	uri := generateEndpointUrl(queueUri(accountSid, queueSid))

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

	err = queue.FromJson(resp.Body)
	if err != nil {
		return
	}

	return
}

func (client *Client) CreateQueue(accountSid string, friendlyName string, maxSize int) (queue *Queue, err error) {
	queue = NewQueue(client, accountSid)

	uri := generateEndpointUrl(queuesUri(accountSid))

	req, err := client.generateRequest("POST", uri, map[string]string{
		"friendly_name": friendlyName,
		"max_size":      strconv.Itoa(maxSize),
	})
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

	err = queue.FromJson(resp.Body)
	if err != nil {
		return
	}

	return
}

func (client *Client) updateQueue(queue *Queue) (err error) {
	uri := generateEndpointUrl(queueUri(queue.AccountSid, queue.Sid))

	req, err := client.generateRequest("POST", uri, map[string]string{
		"friendly_name": queue.FriendlyName,
		"max_size":      fmt.Sprintf("%i", queue.MaxSize),
	})
	if err != nil {
		return
	}

	_, err = client.httpClient.Do(req)
	if err != nil {
		return
	}

	return
}

func (client *Client) deleteQueue(queue *Queue) (err error) {
	uri := generateEndpointUrl(queueUri(queue.AccountSid, queue.Sid))

	req, err := client.generateRequest("DELETE", uri, map[string]string{})
	if err != nil {
		return
	}

	_, err = client.httpClient.Do(req)
	if err != nil {
		return
	}

	return
}
