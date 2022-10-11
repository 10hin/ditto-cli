package model

import (
	"time"
)

type DittoEnvelope struct {
	Topic     string
	Headers   map[string]string
	Path      string
	Fields    string
	Value     interface{}
	Extra     map[string]interface{}
	Revision  int
	Timestamp *time.Time
}

type DittoResponse struct {
	Topic   string
	Headers map[string]string
	Path    string
	Value   interface{}
	Status  int
}
