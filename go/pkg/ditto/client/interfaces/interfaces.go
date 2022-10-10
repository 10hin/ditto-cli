package interfaces

import "golang.10h.in/ditto/cli/pkg/ditto/model"

type Client interface {
	Twin() TwinClient
	Live() LiveClient
	Policy() PolicyClient
}

type ThingClient interface {
	Get(thingID string) (*model.Thing, error)
	List(thingIDs []string) ([]model.Thing, error)
	Create(thingDraft *model.NewThing) (*model.Thing, error)
	Put(thingReplacing *model.Thing) (*model.Thing, error)
	Patch(thingID string, patchThing *model.PatchThing) (*model.Thing, error)
}

type ThinsSearchClient interface {
	Search() (model.ThingList, error)
	Count() (int, error)
}

type PolicyClient interface {
	Get(policyID string) (*model.Policy, error)
	Create(policyDraft *model.NewPolicy) (*model.Policy, error)
	Put(policyReplacing *model.Policy) (*model.Policy, error)
}

type MessageClient interface {
	Claim(thingID string) error
	SendToDevice(thingID, subject string, body []byte) error
	SendFromDevice(thingID, subject string, body []byte) error
	SendToFeature(thingID, featureID, subject string, body []byte) error
	SendFromFeature(thingID, featureID, subject string, body []byte) error
}

type TwinClient interface {
	ThingClient
	ThinsSearchClient
}

type LiveClient interface {
	ThingClient
	MessageClient
}
