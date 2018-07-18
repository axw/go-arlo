package arlo

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

const (
	profileURL = "https://arlo.netgear.com/hmsweb/users/profile"
)

type Profile struct {
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Language       string `json:"language"`
	Country        string `json:"country"`
	AcceptedPolicy int    `json:"acceptedPolicy"`
	CurrentPolicy  int    `json:"acceptedPolicy"`
	ValidEmail     bool   `json:"bool"`
}

func (c *Client) Profile(ctx context.Context) (Profile, error) {
	resp, err := c.doRequest(NewProfileRequest().WithContext(ctx))
	if err != nil {
		return Profile{}, err
	}
	defer resp.Body.Close()

	var result struct {
		Data    Profile `json:"data"`
		Success bool    `json:"success"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Profile{}, err
	}
	if !result.Success {
		return Profile{}, errors.New("profile request unsuccessful")
	}
	return result.Data, nil
}

func NewProfileRequest() *http.Request {
	return mustNewRequest("GET", profileURL, nil)
}
