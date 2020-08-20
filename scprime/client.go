package scprime

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var (
	client = &http.Client{
		Timeout: 30 * time.Second,
	}
)

type (
	//HTTPMethod an http method
	HTTPMethod string
	//APIClient a client used to access the Sia Central API
	APIClient struct {
		BaseAddress string
		AccessKey   string
		AuthToken   string
	}

	//APIResponse APIResponse
	APIResponse struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	}
)

func drainAndClose(rc io.ReadCloser) {
	io.Copy(ioutil.Discard, rc)
	rc.Close()
}

func (a *APIClient) makeAPIRequest(method string, url string, body interface{}, value interface{}) (statusCode int, err error) {
	var req *http.Request

	if !strings.HasPrefix(url, "http") {
		url = a.BaseAddress + url
	}

	if method == http.MethodGet {
		req, err = http.NewRequest(string(method), url, nil)
	} else {
		var buf []byte

		if body != nil {
			buf, err = json.Marshal(body)

			if err != nil {
				return
			}
		}

		req, err = http.NewRequest(string(method), url, bytes.NewBuffer(buf))
	}

	if err != nil {
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		return
	}

	defer drainAndClose(resp.Body)

	dec := json.NewDecoder(resp.Body)
	statusCode = resp.StatusCode
	err = dec.Decode(value)

	return
}

//NewClient creates a new API client
func NewClient() *APIClient {
	return &APIClient{
		BaseAddress: "https://api.siacentral.com/v2/scprime",
	}
}
