package model

import "time"

type Thing struct {
	ThingID    string      `json:"thingId,omitempty"`
	PolicyID   string      `json:"policyId,omitempty"`
	Definition string      `json:"definition,omitempty"`
	Attributes *Attributes `json:"attributes,omitempty"`
	Features   *Features   `json:"features,omitempty"`
	Namespace  string      `json:"_namespace,omitempty"`
	Revision   int         `json:"_revision,omitempty"`
	Created    *time.Time  `json:"_created,omitempty"`
	Modified   *time.Time  `json:"_modified,omitempty"`
	Metadata   *Metadata   `json:"_metadata,omitempty"`
}

type ThingDraft struct {
	PolicyID   string      `json:"policyId,omitempty"`
	Definition string      `json:"definition,omitempty"`
	Attributes *Attributes `json:"attributes,omitempty"`
	Features   *Features   `json:"features,omitempty"`
	Metadata   *Metadata   `json:"_metadata,omitempty"`
}

type Attributes map[string]interface{}

type Features map[string]Feature
type Feature struct {
	Definition        []string           `json:"definition,omitempty""`
	Properties        *Properties        `json:"properties,omitempty"`
	DesiredProperties *DesiredProperties `json:"desiredProperties,omitempty"`
}

type Properties map[string]interface{}
type DesiredProperties map[string]interface{}

type Metadata map[string]interface{}
