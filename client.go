package apisdkgo

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const apiBaseAddress = "https://api.siacentral.com/v2"

/*
 */
const (
	HTTPPost   HTTPMethod = "POST"
	HTTPGet    HTTPMethod = "GET"
	HTTPPut    HTTPMethod = "PUT"
	HTTPDelete HTTPMethod = "DELETE"
	HTTPPatch  HTTPMethod = "PATCH"
)

var (
	client = http.Client{
		Timeout: 30 * time.Second,
	}
)

type (
	//HTTPMethod an http method
	HTTPMethod string
	//APIClient a client used to access the Sia Central API
	APIClient struct {
		AccessKey string
		AuthToken string
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

func makeAPIRequest(method HTTPMethod, url string, body interface{}, value interface{}) (statusCode int, err error) {
	var req *http.Request

	if !strings.HasPrefix(url, "http") {
		url = apiBaseAddress + url
	}

	if method == HTTPGet {
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
	return &APIClient{}
}
