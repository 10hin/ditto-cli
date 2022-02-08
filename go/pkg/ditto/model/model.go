package model

import "time"

type Thing struct {
	ThingID    string
	PolicyID   string
	Definition string
	Attributes Attributes
	Features   Features
	namespace  string
	revision   int
	created    time.Time
	modified   time.Time
	metadata   *Metadata
}

func (t Thing) Namespace() string {
	return t.namespace
}

func (t Thing) Revision() int {
	return t.revision
}

func (t Thing) Created() time.Time {
	return t.created
}

func (t Thing) Modified() time.Time {
	return t.modified
}

func (t Thing) Metadata() *Metadata {
	return t.metadata
}

func (t Thing) HasMetadata() bool {
	return t.metadata != nil
}

type Attributes map[string]interface{}

type Features map[string]Feature
type Feature struct {
	Definition        []string
	Properties        Properties
	DesiredProperties DesiredProperties
}

type Properties map[string]interface{}
type DesiredProperties map[string]interface{}

type Metadata map[string]interface{}
