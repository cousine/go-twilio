package twilio

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	TWILIO_ENDPOINT = "https://api.twilio.com/2010-04-01"
)

type Client struct {
	AccountSid string
	AuthToken  string
	httpClient *http.Client
}

func NewClient(accountSid string, authToken string) (twilioClient *Client) {
	twilioClient = &Client{
		AccountSid: accountSid,
		AuthToken:  authToken,
		httpClient: &http.Client{},
	}

	return
}

func generateEndpointUrl(endpoint string) string {
	return fmt.Sprintf("%v/%v.json", TWILIO_ENDPOINT, endpoint)
}

func (client *Client) generateRequest(method string, uri string, data map[string]string) (request *http.Request, err error) {
	form := url.Values{}

	for key, val := range data {
		form.Add(key, val)
	}

	request, err = http.NewRequest(method, uri, strings.NewReader(form.Encode()))
	if err != nil {
		return
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.SetBasicAuth(client.AccountSid, client.AuthToken)

	return
}
