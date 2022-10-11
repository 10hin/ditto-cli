package interfaces

import (
	"golang.10h.in/ditto/cli/pkg/ditto/lowmodel"
	"golang.10h.in/ditto/cli/pkg/ditto/model"
)

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

type EnvelopeHandler func(lowmodel.DittoEnvelope, Sender) error

type HandlerID interface {
	Topic() string
	Handler() EnvelopeHandler
}

type Sender interface {
	Send(envelope lowmodel.DittoEnvelope) error
	Reply(reply lowmodel.DittoResponse) error
}

type Exchanger interface {
	Exchange(envelope lowmodel.DittoEnvelope) (lowmodel.DittoResponse, error)
}

type HandlerManager interface {
	RegisterHandler(topic string, handler EnvelopeHandler) (handlerID HandlerID, err error)
	UnregisterHandler(handlerID HandlerID) error
}

type UnknownHandlerIDError interface {
	Error() string
}

type LowClient interface {
	Sender
	Exchanger
	HandlerManager
}
