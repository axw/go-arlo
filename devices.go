package arlo

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

const (
	devicesURL = "https://arlo.netgear.com/hmsweb/users/devices"
)

type Device struct {
	UniqueID                      string
	DeviceID                      string
	ParentID                      string
	UserID                        string
	DeviceType                    string
	DeviceName                    string
	LastModified                  time.Time
	FirmwareVersion               string
	XCloudID                      string
	UserRole                      string
	DisplayOrder                  int
	MediaObjectCount              int
	State                         string
	ModelID                       string
	DateCreated                   time.Time
	ArloMobilePlan                bool
	InterfaceVersion              string
	InterfaceSchemaVer            string
	AutomationRevision            int
	PresignedLastImageURL         string
	PresignedSnapshotURL          string
	PresignedFullFrameSnapshotURL string
	Owner                         struct {
		FirstName string
		LastName  string
		OwnerID   string
	}
	Connectivity struct {
		Type      string
		Connected bool
		MEPStatus string
	}
	Properties map[string]interface{}
}

func (c *Client) Devices(ctx context.Context) ([]Device, error) {
	resp, err := c.doRequest(NewDevicesRequest().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	type InnerDevice struct {
		Device
		LastModified int64
		DateCreated  int64
	}

	var result struct {
		Data    []InnerDevice
		Success bool
	}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New("devices request unsuccessful")
	}

	devices := make([]Device, len(result.Data))
	for i, d := range result.Data {
		d.Device.LastModified = newTime(d.LastModified)
		d.Device.DateCreated = newTime(d.DateCreated)
		devices[i] = d.Device
	}
	return devices, nil
}

func NewDevicesRequest() *http.Request {
	return mustNewRequest("GET", devicesURL, nil)
}
