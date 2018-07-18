package arlo

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

const (
	loginURL  = "https://arlo.netgear.com/hmsweb/login/v2"
	logoutURL = "https://arlo.netgear.com/hmsweb/logout"
)

func (c *Client) Login(ctx context.Context, email, password string) error {
	resp, err := c.doRequest(NewLoginRequest(email, password).WithContext(ctx))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return errors.Wrap(err, "login failed")
}

func (c *Client) Logout(ctx context.Context) error {
	resp, err := c.doRequest(NewLogoutRequest().WithContext(ctx))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return errors.Wrap(err, "logout failed")
}

func NewLoginRequest(email, password string) *http.Request {
	body, err := json.Marshal(struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{Email: email, Password: password})
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", loginURL, bytes.NewReader(body))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	return req
}

func NewLogoutRequest() *http.Request {
	return mustNewRequest("PUT", logoutURL, nil)
}
