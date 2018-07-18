package arlo

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

const (
	locationsURL = "https://arlo.netgear.com/hmsweb/users/locations"
)

type Location struct {
	ID                string
	Name              string
	OwnerID           string
	Address           string
	Latitude          float64
	Longitude         float64
	HomeMode          string
	AwayMode          string
	GeoEnabled        bool
	GeoRadius         float64
	UniqueIDs         []string
	SmartDevices      []string
	PushNotifyDevices []string
}

func (c *Client) Locations(ctx context.Context) ([]Location, error) {
	resp, err := c.doRequest(NewLocationsRequest().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Data    []Location `json:"data"`
		Success bool       `json:"success"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New("locations request unsuccessful")
	}
	return result.Data, nil
}

func NewLocationsRequest() *http.Request {
	return mustNewRequest("GET", locationsURL, nil)
}
