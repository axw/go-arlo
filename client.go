package arlo

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Client *http.Client
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	for _, cookie := range c.Client.Jar.Cookies(req.URL) {
		if cookie.Name == "token" {
			req.Header.Set("Authorization", cookie.Value)
			break
		}
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		return resp, nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return nil, &Error{
		StatusCode: resp.StatusCode,
		Body:       string(body),
	}
}

// Error represents an error response from the Arlo API.
type Error struct {
	StatusCode int
	// TODO(axw) unpack JSON? {"data":{"error":"NNNN","message:"foo","reason":"bar"},"success":false}
	Body string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d (%s): %s", e.StatusCode, http.StatusText(e.StatusCode), e.Body)
}
