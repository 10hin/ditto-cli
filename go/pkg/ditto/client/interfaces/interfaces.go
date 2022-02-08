package interfaces

import "golang.10h.in/ditto/cli/pkg/ditto/model"

type Client interface {
	Thing() ThingClient
}

type ThingClient interface {
	GetThing(thingID string) model.Thing
}
