package http

import (
	"golang.10h.in/ditto/cli/pkg/ditto/client/interfaces"
	"golang.10h.in/ditto/cli/pkg/ditto/config"
	"net/http"
	"time"
)

func NewClient(c *config.HTTPConfig) interfaces.Client {
	client := &http.Client{
		// fixed timeout; user can't set this
		Timeout: 120 * time.Second,
	}
	thingCli := newThingClient(client, c)
	return &httpClient{
		client:      client,
		thingClient: thingCli,
	}
}

type httpClient struct {
	client      *http.Client
	thingClient interfaces.ThingClient
}

func (c *httpClient) Thing() interfaces.ThingClient {
	return c.thingClient
}
