package api

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func GetClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
