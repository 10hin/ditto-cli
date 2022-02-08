package http

import (
	"golang.10h.in/ditto/cli/pkg/ditto/client/interfaces"
	"golang.10h.in/ditto/cli/pkg/ditto/model"
	"net/http"
)

func newThingClient(client *http.Client) interfaces.ThingClient {
	return &thingClient{
		client: client,
	}
}

type thingClient struct {
	client *http.Client
}

func (c *thingClient) GetThing(thingID string) model.Thing {

}
