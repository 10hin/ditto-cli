package http

import (
	"encoding/json"
	"fmt"
	"golang.10h.in/ditto/cli/pkg/ditto/client/interfaces"
	"golang.10h.in/ditto/cli/pkg/ditto/config"
	"golang.10h.in/ditto/cli/pkg/ditto/model"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	pathAPIV2  = "/api/2"
	pathThings = "/things"
)

func newThingClient(client *http.Client, cfg *config.HTTPConfig) interfaces.ThingClient {
	return &thingClient{
		urlPrefix: strings.TrimSuffix(cfg.URLPrefix, "/"),
		client:    client,
		// currently user can't set timeout
		timeout: 60 * time.Second,
	}
}

type thingClient struct {
	urlPrefix string
	client    *http.Client
	// use fixed timeout to any request
	timeout time.Duration
}

func (c *thingClient) Get(thingID string) (*model.Thing, error) {
	var err error

	reqURL := fmt.Sprintf(
		"%s%s%s/%s?fields=%s&timeout=%dms&channel=%s",
		c.urlPrefix,
		pathAPIV2,
		pathThings,
		thingID,
		strings.Join([]string{
			"thingId",
			"policyId",
			"attributes",
			"features",
		}, ","),
		c.timeout.Milliseconds(),
		"twin",
	)

	var resp *http.Response
	resp, err = c.client.Get(reqURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		var msg []byte
		msg, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}
		return nil, fmt.Errorf("unexpected response status: %d; response body: %s", resp.StatusCode, (string)(msg))
	}

	var t model.Thing
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response as JSON: %w", err)
	}

	return &t, nil
}
