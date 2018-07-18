package arlo

import (
	"io"
	"net/http"
	"time"
)

func newTime(ms int64) time.Time {
	return time.Unix(ms/1000, (ms%1000)*1000000)
}

func mustNewRequest(method, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}
	return req
}
