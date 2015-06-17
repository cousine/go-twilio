package twilio

import (
	"github.com/cousine/go-twilio/helpers"
	"io"
	"strconv"
)

type Paginator interface {
	NextPage(int) (Paginator, error)
	PreviousPage(int) (Paginator, error)
	FromJson(io.ReadCloser) error
}

type PagingInformation struct {
	Uri             string `json:"uri"`
	FirstPageUri    string `json:"first_page_uri"`
	NextPageUri     string `json:"next_page_uri"`
	Page            int    `json:"page"`
	PageSize        int    `json:"page_size"`
	PreviousPageUri string `json:"previous_page_uri"`
}

func getPage(paginator Paginator, pageUri string, client *Client, pageSize int) (err error) {
	uri := generateEndpointUrl(pageUri)

	req, err := client.generateRequest("GET", uri, map[string]string{"PageSize": strconv.Itoa(pageSize)})
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

	err = paginator.FromJson(resp.Body)
	if err != nil {
		return
	}

	return
}

func (pagination PagingInformation) Next(paginator Paginator, client *Client, pageSize int) (err error) {
	return getPage(paginator, pagination.NextPageUri, client, pageSize)
}

func (pagination PagingInformation) Previous(paginator Paginator, client *Client, pageSize int) (err error) {
	return getPage(paginator, pagination.PreviousPageUri, client, pageSize)
}
