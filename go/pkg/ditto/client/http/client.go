package http

import (
	"golang.10h.in/ditto/cli/pkg/ditto/client/interfaces"
	"golang.10h.in/ditto/cli/pkg/ditto/config"
	"net/http"
	"time"
)

func NewClient(c *config.HTTPConfig) interfaces.Client {
	var customTransport = &customRoundTripper{
		requestCustomizers:  []requestCustomizer{},
		responseCustomizers: []responseCustomizer{},
	}
	if c.Basic != nil {
		customTransport.requestCustomizers = append(customTransport.requestCustomizers, func(req *http.Request) {
			req.SetBasicAuth(c.Basic.Username, c.Basic.Password)
		})
	}
	client := &http.Client{
		Transport: customTransport,
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

type requestCustomizer func(r *http.Request)
type responseCustomizer func(r *http.Response)

type customRoundTripper struct {
	requestCustomizers  []requestCustomizer
	responseCustomizers []responseCustomizer
}

func (crt *customRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	for _, customizer := range crt.requestCustomizers {
		customizer(req)
	}

	resp, err := http.DefaultTransport.RoundTrip(req)

	if resp != nil {
		for _, customizer := range crt.responseCustomizers {
			customizer(resp)
		}
	}

	return resp, err
}
