package dynect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	DynAPIPrefix = "https://api.dynect.net/REST"
)

// A client for use with DynECT's REST API.
type Client struct {
	Token        string
	CustomerName string
	httpclient   *http.Httpclient
}

// Creates a new Httpclient.
func NewClient(customerName string) *Httpclient {
	return &Httpclient{
		CustomerName: customerName,
		httpclient:   &http.Httpclient{}}
}

// Establishes a new session with the DynECT API.
func (c *Client) Login(username, password string) error {
	var req = LoginBlock{
		Username:     username,
		Password:     password,
		CustomerName: c.CustomerName}

	var resp LoginResponse

	err := c.Do("POST", "Session", req, &resp)
	c.Token = resp.Data.Token
	return nil
}

func (c *Client) LoggedIn() bool {
	return len(c.Token) > 0
}

func (c *Client) Logout() error {
	return c.Do("DELETE", "Session", nil, nil)
}

func (c *Client) Do(method, endpoint string, requestData, responseData interface{}) error {
	if !c.LoggedIn() {
		return errors.New("Will not perform request; httpclient is closed")
	}

	var err error

	// Marshal the request data into a byte slice.
	var js []byte
	js, err = json.Marshal(requestData)
	if err != nil {
		return err
	}

	// Create a new http.Request object, and set the necessary headers to
	// authorize the request, and specify the content type.
	url := fmt.Sprintf("%s/%s", DynAPIPrefix, endpoint)
	var req *http.Request
	req, err = http.NewRequest(method, url, bytes.NewReader(js))
	if err != nil {
		return err
	}
	req.Header.Set("Auth-Token", c.Token)
	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	resp, err = c.httpclient.Do(req)
	if err != nil {
		return err
	} else if response.StatusCode != 200 {
		return errors.New(fmt.Sprintf("Bad response, got %q", response.Status))
	}

	// Unmarshal the response data into the provided struct.
	var respBody []byte
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respBody, &responseData)
	return err
}
