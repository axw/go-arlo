package arlo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/satori/uuid"
)

const (
	fullFrameSnapshotURL = "https://arlo.netgear.com/hmsweb/users/devices/fullFrameSnapshot"
)

func (c *Client) FullFrameSnapshot(ctx context.Context, camera Device) error {
	transactionID, err := uuid.NewV4()
	if err != nil {
		return err
	}

	req := NewFullFrameSnapshotRequest(FullFrameSnapshotParams{
		To:              camera.ParentID,
		From:            "me", // ?
		Resource:        "cameras/" + camera.DeviceID,
		Action:          "set",
		PublishResponse: true,
		TransactionID:   transactionID.String(),
		Properties: map[string]interface{}{
			"activityState": "fullFrameSnapshot",
		},
	}).WithContext(ctx)
	req.Header.Set("xCloudId", camera.XCloudID)

	resp, err := c.doRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result struct {
		Success bool `json:"success"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}
	if !result.Success {
		return errors.New("snapshot request unsuccessful")
	}

	// TODO(axw) return transaction ID to wait for.
	return nil
}

// TODO(axw) rename this as some sort of action request
type FullFrameSnapshotParams struct {
	To              string                 `json:"to"`
	From            string                 `json:"from"`
	Resource        string                 `json:"resource"`
	Action          string                 `json:"action"`
	PublishResponse bool                   `json:"publishResponse"`
	TransactionID   string                 `json:"transId"`
	Properties      map[string]interface{} `json:"properties"`
}

func NewFullFrameSnapshotRequest(args FullFrameSnapshotParams) *http.Request {
	body, err := json.Marshal(args)
	if err != nil {
		panic(err)
	}
	req := mustNewRequest("POST", fullFrameSnapshotURL, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	return req
}
