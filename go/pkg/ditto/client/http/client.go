package http

import (
	"golang.10h.in/ditto/cli/pkg/ditto/client/interfaces"
	"golang.10h.in/ditto/cli/pkg/ditto/config"
	"net/http"
)

func NewClient(c *config.HTTPConfig) interfaces.Client {

}

type httpClient struct {
	client http.Client
}

func (c *httpClient) Thing() interfaces.ThingClient {

}
