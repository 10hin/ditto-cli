package interfaces

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"golang.10h.in/ditto/cli/pkg/ditto/lowmodel"
	"golang.10h.in/ditto/cli/pkg/ditto/model"
	"strings"
)

type Client interface {
	Twin() TwinClient
	Live() LiveClient
	Policy() PolicyClient
}

type ThingClient interface {
	Get(thingID string, opt GetThingOption) (*model.Thing, error)
	List(thingIDs []string, opt GetThingOption) ([]model.Thing, error)
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

type Channel string

const (
	TWIN Channel = "twin"
	LIVE Channel = "live"
)

func (c Channel) String() string {
	return (string)(c)
}

type thingClientImpl struct {
	lowClient LowClient
	channel   Channel
}

func newThingClientImpl(lowClient LowClient) ThingClient {
	return &thingClientImpl{
		lowClient: lowClient,
	}
}

type RequestOption func(header lowmodel.DittoHeader)

type GetThingOption struct {
	Fields         []string
	RequestOptions []RequestOption
}

func (t *thingClientImpl) Get(thingID string, opt GetThingOption) (*model.Thing, error) {
	thingIDSlice := strings.Split(thingID, ":")
	namespace := thingIDSlice[0]
	thingName := thingIDSlice[1]
	topic := fmt.Sprintf("%s/%s/things/%s/commands/retrieve", namespace, thingName, t.channel.String())
	correlationID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	headers := lowmodel.NewHeader(correlationID.String())
	for _, reqOpt := range opt.RequestOptions {
		reqOpt(headers)
	}
	path := "/"

	envelope := lowmodel.NewDittoEnvelope(topic, headers, path)
	fields := strings.Join(opt.Fields, ",")
	envelope.SetFields(&fields)

	var resp lowmodel.DittoResponse
	resp, err = t.lowClient.Exchange(envelope)
	if err != nil {
		return nil, err
	}

	var valueBytes []byte
	valueBytes, err = json.Marshal(resp.Value())
	if err != nil {
		return nil, err
	}

	var thing model.Thing
	err = json.Unmarshal(valueBytes, &thing)
	if err != nil {
		return nil, err
	}

	return &thing, nil
}

func (t *thingClientImpl) List(thingIDs []string, opt GetThingOption) ([]model.Thing, error) {
	topic := fmt.Sprintf("/_/things/%s/commands/retrieve", t.channel.String())
	correlationID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	headers := lowmodel.NewHeader(correlationID.String())
	for _, reqOpt := range opt.RequestOptions {
		reqOpt(headers)
	}
	path := "/"

	envelope := lowmodel.NewDittoEnvelope(topic, headers, path)
	fields := strings.Join(opt.Fields, ",")
	envelope.SetFields(&fields)
	envelope.SetValue(thingIDs)

	var resp lowmodel.DittoResponse
	resp, err = t.lowClient.Exchange(envelope)
	if err != nil {
		return nil, err
	}

	var valueBytes []byte
	valueBytes, err = json.Marshal(resp.Value())
	if err != nil {
		return nil, err
	}

	var things []model.Thing
	err = json.Unmarshal(valueBytes, &things)
	if err != nil {
		return nil, err
	}

	return things, nil
}

func (t *thingClientImpl) Create(thingDraft *model.NewThing) (*model.Thing, error) {
	//TODO implement me
	panic("implement me")
}

func (t *thingClientImpl) Put(thingReplacing *model.Thing) (*model.Thing, error) {
	//TODO implement me
	panic("implement me")
}

func (t *thingClientImpl) Patch(thingID string, patchThing *model.PatchThing) (*model.Thing, error) {
	//TODO implement me
	panic("implement me")
}

type thingsSearchClientImpl struct {
	lowClient LowClient
}

func newThingsSearchClientImpl(lowClient LowClient) ThinsSearchClient {
	return &thingsSearchClientImpl{
		lowClient: lowClient,
	}
}

func (t thingsSearchClientImpl) Search() (model.ThingList, error) {
	//TODO implement me
	panic("implement me")
}

func (t thingsSearchClientImpl) Count() (int, error) {
	//TODO implement me
	panic("implement me")
}

type twinClientImpl struct {
	ThingClient
	ThinsSearchClient
}

func newTwinClientImpl(client LowClient) TwinClient {
	return twinClientImpl{
		newThingClientImpl(client),
		newThingsSearchClientImpl(client),
	}
}
