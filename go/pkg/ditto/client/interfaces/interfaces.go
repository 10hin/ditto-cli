package interfaces

import "golang.10h.in/ditto/cli/pkg/ditto/model"

type Client interface {
	Thing() ThingClient
}

type ThingClient interface {
	Get(thingID string) (*model.Thing, error)
	List(thingIDs []string) ([]*model.Thing, error)
	Create(tingDraft *model.ThingDraft) (*model.Thing, error)
}
