package twilio

import (
	"encoding/json"
	"fmt"
	"github.com/cousine/go-twilio/helpers"
	"io"
	"time"
)

const (
	ACCOUNT_ENDPOINT = "Accounts"
)

type AccountType string

const (
	TYPE_TRIAL AccountType = "Trial"
	TYPE_FULL  AccountType = "Full"
)

type AccountStatus string

const (
	STATUS_ACTIVE    AccountStatus = "active"
	STATUS_SUSPENDED AccountStatus = "suspended"
	STATUS_CLOSED    AccountStatus = "closed"
)

type Account struct {
	Sid             string            `json:"sid"`
	DateCreated     time.Time         `json:"date_created"`
	DateUpdated     time.Time         `json:"date_updated"`
	FriendlyName    string            `json:"friendly_name"`
	Type            AccountType       `json:"type"`
	Status          AccountStatus     `json:"status"`
	AuthToken       string            `json:"auth_token"`
	Uri             string            `json:"uri"`
	SubresourceUris map[string]string `json:"subresource_uris"`
	OwnerAccountSid string            `json:"owner_account_sid"`

	client *Client
}

func NewAccount(client *Client) *Account {
	return &Account{
		client: client,
	}
}

func accountUri(sid string) string {
	return fmt.Sprintf("%v/%v", ACCOUNT_ENDPOINT, sid)
}

func (account *Account) FromJson(rawJson io.ReadCloser) error {
	jsonDecoder := json.NewDecoder(rawJson)
	err := jsonDecoder.Decode(account)

	return err
}

func (account *Account) Update(friendlyName string, status AccountStatus) (err error) {
	account.FriendlyName = friendlyName
	account.Status = status

	return account.client.updateAccount(account)
}

func (client *Client) GetAccount(accountSid string) (account *Account, err error) {
	account = NewAccount(client)

	uri := generateEndpointUrl(accountUri(accountSid))

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

	err = account.FromJson(resp.Body)
	if err != nil {
		return
	}

	return
}

func (client *Client) CreateSubaccount(friendlyName string) (account *Account, err error) {
	account = NewAccount(client)

	uri := generateEndpointUrl(ACCOUNT_ENDPOINT)

	req, err := client.generateRequest("POST", uri, map[string]string{"friendly_name": friendlyName})
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

	err = account.FromJson(resp.Body)
	if err != nil {
		return
	}

	return
}

func (client *Client) updateAccount(account *Account) (err error) {
	uri := generateEndpointUrl(accountUri(account.Sid))

	req, err := client.generateRequest("PUT", uri, map[string]string{
		"friendly_name": account.FriendlyName,
		"status":        string(account.Status),
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
