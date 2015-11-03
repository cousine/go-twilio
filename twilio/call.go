package twilio

import (
	"encoding/json"
	"fmt"
	"github.com/cousine/go-twilio/helpers"
	"io"
)

const (
	CALL_ENDPOINT = "Calls"
)

type Call struct {
	Sid            string `json:"sid"`
	ParentCallSid  string `json:"parent_call_sid"`
	DateCreated    string `json:"date_created"`
	DateUpdated    string `json:"date_updated"`
	AccountSid     string `json:"account_sid"`
	To             string `json:"to"`
	From           string `json:"from"`
	PhoneNumberSid string `json:"phone_number_sid"`
	Status         string `json:"status"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
	Duration       string `json:"duration"`
	Price          string `json:"price"`
	PriceUnit      string `json:"price_unit"`
	Direction      string `json:"direction"`
	AnsweredBy     string `json:"answered_by"`
	ForwardedFrom  string `json:"forwarded_from"`
	CallerName     string `json:"caller_name"`
	Uri            string `json:"uri"`

	client *Client
}

func NewCall(client *Client, accountSid string) *Call {
	return &Call{
		AccountSid: accountSid,
		client:     client,
	}
}

func callUri(accountSid string, callSid string) string {
	return fmt.Sprintf("%v/%v/%v/%v", ACCOUNT_ENDPOINT, accountSid, CALL_ENDPOINT, callSid)
}

func callsUri(accountSid string) string {
	return fmt.Sprintf("%v/%v/%v", ACCOUNT_ENDPOINT, accountSid, CALL_ENDPOINT)
}

func (call *Call) FromJson(rawJson io.ReadCloser) error {
	jsonDecoder := json.NewDecoder(rawJson)
	err := jsonDecoder.Decode(call)

	return err
}

func (call *Call) Update(url string, method string, status string) (err error) {
	return call.client.updateCall(call, url, method, status)
}

func (call *Call) Delete() (err error) {
	return call.client.deleteCall(call)
}

func (client *Client) GetCall(accountSid string, callSid string) (call *Call, err error) {
	call = NewCall(client, accountSid)

	uri := generateEndpointUrl(callUri(accountSid, callSid))

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

	err = call.FromJson(resp.Body)
	if err != nil {
		return
	}

	return
}

func (client *Client) updateCall(call *Call, url string, method string, status string) (err error) {
	uri := generateEndpointUrl(callUri(call.AccountSid, call.Sid))

	formParams := make(map[string]string)
	formParams["Url"] = url
	formParams["Method"] = method

	if status != "" {
		formParams["Status"] = status
	}

	req, err := client.generateRequest("POST", uri, formParams)
	if err != nil {
		return
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return
	}

	if helpers.ResponseHasError(resp) {
		return NewTwilioError(resp.Body)
	}

	err = call.FromJson(resp.Body)
	if err != nil {
		return
	}

	return
}

func (client *Client) deleteCall(call *Call) (err error) {
	uri := generateEndpointUrl(callUri(call.AccountSid, call.Sid))

	req, err := client.generateRequest("DELETE", uri, map[string]string{})
	if err != nil {
		return
	}

	resp, err := client.httpClient.Do(req)
	if helpers.ResponseHasError(resp) {
		return NewTwilioError(resp.Body)
	}

	return
}

func (client *Client) Call(accountSid string, from string, to string, url string, appSid string, optionalParams ...string) (call *Call, err error) {
	call = NewCall(client, accountSid)
	uri := generateEndpointUrl(callsUri(accountSid))

	formParams := make(map[string]string)
	formParams["From"] = from
	formParams["To"] = to
	formParams["Url"] = url

	if appSid != "" {
		formParams["ApplicationSid"] = appSid
	}

	if len(optionalParams)%2 != 0 {
		return nil, NewLocalTwilioError("Parameters list must be pairs of keys and values, and must have an even length")
	}

	for i := 0; i < len(optionalParams); i += 2 {
		key := optionalParams[i]
		val := optionalParams[i+1]

		formParams[key] = val
	}

	req, err := client.generateRequest("POST", uri, formParams)
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

	err = call.FromJson(resp.Body)
	if err != nil {
		return
	}

	return
}
