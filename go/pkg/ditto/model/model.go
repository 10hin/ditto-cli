package model

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Thing struct {
	ThingID    string     `json:"thingId"`
	PolicyID   string     `json:"policyId"`
	Definition string     `json:"definition,omitempty"`
	Attributes Attributes `json:"attributes,omitempty"`
	Features   Features   `json:"features,omitempty"`
	Namespace  string     `json:"_namespace,omitempty"`
	Revision   int        `json:"_revision,omitempty"`
	Created    *time.Time `json:"_created,omitempty"`
	Modified   *time.Time `json:"_modified,omitempty"`
	Metadata   Metadata   `json:"_metadata,omitempty"`
}

type NewThing struct {
	PolicyID   string     `json:"policyId"`
	Policy     *Policy    `json:"_policy,omitempty"`
	Definition string     `json:"definition,omitempty"`
	Attributes Attributes `json:"attributes,omitempty"`
	Features   Features   `json:"features,omitempty"`
	Metadata   Metadata   `json:"_metadata,omitempty"`
}

type PatchThing struct {
	ThingID    string     `json:"thingId"`
	PolicyID   string     `json:"policyId"`
	Definition string     `json:"definition,omitempty"`
	Attributes Attributes `json:"attributes,omitempty"`
	Features   Features   `json:"features,omitempty"`
}

type ThingList struct {
	Items  []Thing `json:"items,omitempty"`
	Cursor string  `json:"cursor,omitempty"`
}

type Attributes map[string]interface{}

type Features map[string]*Feature
type Feature struct {
	Definition        []string    `json:"definition,omitempty""`
	Properties        *Properties `json:"properties,omitempty"`
	DesiredProperties *Properties `json:"desiredProperties,omitempty"`
}

type Properties map[string]interface{}

type Metadata map[string]interface{}

type Policy struct {
	PolicyID string        `json:"policyId"`
	Entries  PolicyEntries `json:"entries"`
}

type NewPolicy struct {
	Entries PolicyEntries `json:"entries"`
}

type PolicyEntries map[string]*PolicyEntry

type PolicyEntry struct {
	Subjects  Subjects  `json:"subjects"`
	Resources Resources `json:"resources"`
}

type Subjects map[string]*SubjectEntry

type SubjectEntry struct {
	Type                string               `json:"type"`
	Expiry              *time.Time           `json:"expiry,omitempty"`
	SubjectAnnouncement *SubjectAnnouncement `json:"announcement,omitempty"`
}

type SubjectAnnouncement struct {
	BeforeExpiry          *StringDuration `json:"beforeExpiry,omitempty"`
	WhenDeleted           bool            `json:"whenDeleted,omitempty"`
	RequestedACKs         *RequestedACKs  `json:"requestedAcks,omitempty"`
	RandomizationInterval *StringDuration `json:"randomizationInterval,omitempty"`
}

type StringDuration time.Duration

func (d *StringDuration) MarshalJSON() ([]byte, error) {
	return ([]byte)(fmt.Sprintf("\"%ds\"", ((*time.Duration)(d)).Nanoseconds()/1_000_000_000)), nil
}

func (d *StringDuration) UnmarshalJSON(content []byte) error {
	trimed := strings.Trim((string)(content), "\"")
	suffix := trimed[len(trimed)-1 : len(trimed)]
	count, err := strconv.Atoi(trimed[:len(trimed)-1])
	if err != nil {
		return err
	}

	var parsed time.Duration
	if suffix == "h" {
		parsed = time.Hour * (time.Duration)(count)
	} else if suffix == "m" {
		parsed = time.Minute * (time.Duration)(count)
	} else if suffix == "s" {
		parsed = time.Second * (time.Duration)(count)
	} else {
		return &json.UnsupportedValueError{
			Value: reflect.ValueOf(trimed),
			Str:   trimed,
		}
	}

	*d = (StringDuration)(parsed)
	return nil
}

type RequestedACKs struct {
	Labels  []string        `json:"labels,omitempty"`
	Timeout *StringDuration `json:"timeout,omitempty"`
}

type Resources map[string]*ResourceEntry

type ResourceEntry struct {
	Grant  *Permission `json:"grant,omitempty"`
	Revoke *Permission `json:"revoke,omitempty"`
}

type Permission uint8

func (p *Permission) MarshalJSON() ([]byte, error) {
	switch *p {
	case 0:
		return ([]byte)("[]"), nil
	case 1:
		return ([]byte)("[\"READ\"]"), nil
	case 2:
		return ([]byte)("[\"WRITE\"]"), nil
	case 3:
		return ([]byte)("[\"READ\", \"WRITE\"]"), nil
	default:
		return nil, &json.MarshalerError{
			Type: reflect.TypeOf(p),
			Err:  fmt.Errorf("%d is not expected value as Permission", *p),
		}
	}
}

func (p *Permission) GrantRead() {
	*p = *p | 1
}

func (p *Permission) RevokeRead() {
	*p = *p & 2
}

func (p *Permission) GrantWrite() {
	*p = *p | 2
}

func (p *Permission) RevokeWrite() {
	*p = *p & 1
}
