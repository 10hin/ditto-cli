package http

import (
	"bytes"
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
	apiV2Path             = "/api/2"
	thingsAPIPath         = "/things"
	listingQueryKeyForIDs = "ids"
	contentType           = "application/json"
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
		apiV2Path,
		thingsAPIPath,
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

func (c *thingClient) List(thingIDs []string) ([]*model.Thing, error) {
	var err error

	reqURL := fmt.Sprintf(
		"%s%s%s/?%s=%s&fields=%s&timeout=%dms",
		c.urlPrefix,
		apiV2Path,
		thingsAPIPath,
		listingQueryKeyForIDs,
		strings.Join(thingIDs, ","),
		strings.Join([]string{
			"thingId",
			"policyId",
			"attributes",
			"features",
		}, ","),
		c.timeout.Milliseconds(),
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

	var l []*model.Thing
	err = json.NewDecoder(resp.Body).Decode(&l)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response as JSON: %w", err)
	}

	return l, nil
}

func (c *thingClient) Create(thingDraft *model.ThingDraft) (*model.Thing, error) {
	var err error
	url := fmt.Sprintf("%s%s%s/", c.urlPrefix, apiV2Path, thingsAPIPath)

	var body bytes.Buffer
	err = json.NewEncoder(&body).Encode(thingDraft)
	if err != nil {
		return nil, err
	}

	var resp *http.Response
	resp, err = c.client.Post(url, contentType, &body)
	if err != nil {
		return nil, err
	}

	if 400 <= resp.StatusCode && resp.StatusCode < 500 {
		var respBody []byte
		respBody, err = ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("server respond client error: %s: %s", resp.Status, (string)(respBody))
	}

	return nil, fmt.Errorf("not implemented yet")
}
