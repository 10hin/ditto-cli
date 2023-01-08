package lowmodel

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type DittoEnvelope interface {
	json.Marshaler
	json.Unmarshaler
	Topic() string
	SetTopic(topic string)
	Headers() DittoHeader
	SetHeaders(headers DittoHeader)
	Path() string
	SetPath(path string)
	Fields() *string
	SetFields(fields *string)
	Value() interface{}
	SetValue(value interface{})
	Extra() map[string]interface{}
	SetExtra(extra map[string]interface{})
	Revision() *int
	SetRevision(revision *int)
	Timestamp() *time.Time
	SetTimestamp(timestamp *time.Time)
}

func NewDittoEnvelope(topic string, headers DittoHeader, path string) DittoEnvelope {
	return &dittoEnvelopeImpl{
		topic:   topic,
		headers: headers,
		path:    path,
	}
}

type dittoEnvelopeImpl struct {
	topic     string
	headers   DittoHeader
	path      string
	fields    *string
	value     interface{}
	extra     map[string]interface{}
	revision  *int
	timestamp *time.Time
}

func (d *dittoEnvelopeImpl) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"topic":   d.topic,
		"headers": d.headers,
		"path":    d.path,
	}

	if d.fields != nil {
		m["fields"] = *d.fields
	}

	if d.value != nil {
		m["value"] = d.value
	}

	if d.extra != nil {
		m["extra"] = d.extra
	}

	if d.revision != nil {
		m["revision"] = *d.revision
	}

	if d.timestamp != nil {
		m["timestamp"] = d.timestamp.Format(time.RFC3339)
	}

	return json.Marshal(m)
}

func (d *dittoEnvelopeImpl) UnmarshalJSON(bytes []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return err
	}

	// required fields

	{
		var topicI interface{}
		topicI, ok := m["topic"]
		if !ok {
			return fmt.Errorf("required field \"topic\" not found")
		}
		var topic string
		topic, ok = topicI.(string)
		if !ok {
			return fmt.Errorf("\"topic\" field expected string type, but not: %#v", topicI)
		}
		d.topic = topic
	}

	{
		var headersI interface{}
		headersI, ok := m["headers"]
		if !ok {
			return fmt.Errorf("required field \"headers\" not found")
		}
		var headersBytes []byte
		headersBytes, err = json.Marshal(headersI)
		headers := NewHeader("")
		err = json.Unmarshal(headersBytes, headers)
		if err != nil {
			return fmt.Errorf("\"headers\" field expected DittoHeader type, but not: %s; %w", (string)(headersBytes), err)
		}
		d.headers = headers
	}

	{
		var pathI interface{}
		pathI, ok := m["path"]
		if !ok {
			return fmt.Errorf("required field \"path\" not found")
		}
		var path string
		path, ok = pathI.(string)
		if !ok {
			return fmt.Errorf("\"path\" field expected string type, but not: %#v", pathI)
		}
		d.path = path
	}

	// optional fields

	{
		var fieldsI interface{}
		fieldsI, ok := m["fields"]
		if ok {
			var fields string
			fields, ok = fieldsI.(string)
			if !ok {
				return fmt.Errorf("\"fields\" field expected string type, but not: %#v", fieldsI)
			}
			d.fields = &fields
		}
	}

	{
		value, ok := m["value"]
		if ok {
			d.value = value
		}
	}

	{
		var extraI interface{}
		extraI, ok := m["extra"]
		if ok {
			var extra map[string]interface{}
			extra, ok = extraI.(map[string]interface{})
			if !ok {
				return fmt.Errorf("\"extra\" field expected object type, but not: %#v", extraI)
			}
			d.extra = extra
		}
	}

	{
		var revisionI interface{}
		revisionI, ok := m["revision"]
		if ok {
			var revision int
			revision, ok = revisionI.(int)
			if !ok {
				return fmt.Errorf("\"revision\" field expected int type, but not: %#v", revisionI)
			}
			d.revision = &revision
		}
	}

	{
		var timestampI interface{}
		timestampI, ok := m["timestamp"]
		if ok {
			var timestampStr string
			timestampStr, ok = timestampI.(string)
			if !ok {
				return fmt.Errorf("\"timestamp\" field expected string type, but not: %#v", timestampI)
			}
			var timestamp time.Time
			timestamp, err = time.Parse(time.RFC3339, timestampStr)
			if err != nil {
				return fmt.Errorf("\"timestamp\" field expected formatted as RFC3339 date-time, but not: %s; %w", timestampStr, err)
			}
			d.timestamp = &timestamp
		}
	}

	return nil
}

func (d *dittoEnvelopeImpl) Topic() string {
	return d.topic
}

func (d *dittoEnvelopeImpl) SetTopic(topic string) {
	d.topic = topic
}

func (d *dittoEnvelopeImpl) Headers() DittoHeader {
	return d.headers
}

func (d *dittoEnvelopeImpl) SetHeaders(headers DittoHeader) {
	d.headers = headers
}

func (d *dittoEnvelopeImpl) Path() string {
	return d.path
}

func (d *dittoEnvelopeImpl) SetPath(path string) {
	d.path = path
}

func (d *dittoEnvelopeImpl) Fields() *string {
	return d.fields
}

func (d *dittoEnvelopeImpl) SetFields(fields *string) {
	d.fields = fields
}

func (d *dittoEnvelopeImpl) Value() interface{} {
	return d.value
}

func (d *dittoEnvelopeImpl) SetValue(value interface{}) {
	d.value = value
}

func (d *dittoEnvelopeImpl) Extra() map[string]interface{} {
	return d.extra
}

func (d *dittoEnvelopeImpl) SetExtra(extra map[string]interface{}) {
	d.extra = extra
}

func (d *dittoEnvelopeImpl) Revision() *int {
	return d.revision
}

func (d *dittoEnvelopeImpl) SetRevision(revision *int) {
	d.revision = revision
}

func (d *dittoEnvelopeImpl) Timestamp() *time.Time {
	return d.timestamp
}

func (d *dittoEnvelopeImpl) SetTimestamp(timestamp *time.Time) {
	d.timestamp = timestamp
}

type DittoHeader interface {
	json.Marshaler
	json.Unmarshaler
	ContentType() *string
	SetContentType(contentType *string)
	CorrelationID() string
	SetCorrelationID(correlationID string)
	DittoOriginator() *string
	IfMatch() *string
	SetIfMatch(condition *string)
	IfNoneMatch() *string
	SetIfNoneMatch(condition *string)
	ResponseRequired() bool
	SetResponseRequired(required bool)
	RequestedACKs() *string
	SetRequestedACKs(requestedACKs *string)
	Timeout() *time.Duration
	SetTimeout(timeout *time.Duration)
	Version() *int
	SetVersion(version *int)
	Condition() *string
	SetCondition(rql *string)
}

func NewHeader(correlationID string) DittoHeader {
	return &dittoHeaderImpl{
		correlationID:    correlationID,
		responseRequired: false,
	}
}

type dittoHeaderImpl struct {
	contentType      *string
	correlationID    string
	dittoOriginator  *string
	ifMatch          *string
	ifNoneMatch      *string
	responseRequired bool
	requestedACKs    *string
	timeout          *time.Duration
	version          *int
	condition        *string
}

var (
	headerOptionalStringFields = map[string]string{
		"content-type":     "contentType",
		"ditto-originator": "dittoOriginator",
		"If-Match":         "ifMatch",
		"If-None-Match":    "ifNoneMatch",
		"requested-acks":   "requestedACKs",
		"condition":        "condition",
	}
)

func (d *dittoHeaderImpl) ContentType() *string {
	return d.contentType
}

func (d *dittoHeaderImpl) SetContentType(contentType *string) {
	d.contentType = contentType
}

func (d *dittoHeaderImpl) CorrelationID() string {
	return d.correlationID
}

func (d *dittoHeaderImpl) SetCorrelationID(correlationID string) {
	d.correlationID = correlationID
}

func (d *dittoHeaderImpl) DittoOriginator() *string {
	return d.dittoOriginator
}

func (d *dittoHeaderImpl) IfMatch() *string {
	return d.ifMatch
}

func (d *dittoHeaderImpl) SetIfMatch(condition *string) {
	d.ifMatch = condition
}

func (d *dittoHeaderImpl) IfNoneMatch() *string {
	return d.ifNoneMatch
}

func (d *dittoHeaderImpl) SetIfNoneMatch(condition *string) {
	d.ifNoneMatch = condition
}

func (d *dittoHeaderImpl) ResponseRequired() bool {
	return d.responseRequired
}

func (d *dittoHeaderImpl) SetResponseRequired(required bool) {
	d.responseRequired = required
}

func (d *dittoHeaderImpl) RequestedACKs() *string {
	return d.requestedACKs
}

func (d *dittoHeaderImpl) SetRequestedACKs(requestedACKs *string) {
	d.requestedACKs = requestedACKs
}

func (d *dittoHeaderImpl) Timeout() *time.Duration {
	return d.timeout
}

func (d *dittoHeaderImpl) SetTimeout(timeout *time.Duration) {
	d.timeout = timeout
}

func (d *dittoHeaderImpl) Version() *int {
	return d.version
}

func (d *dittoHeaderImpl) SetVersion(version *int) {
	d.version = version
}

func (d *dittoHeaderImpl) Condition() *string {
	return d.condition
}

func (d *dittoHeaderImpl) SetCondition(rql *string) {
	d.condition = rql
}

func (d *dittoHeaderImpl) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"correlation-id": d.correlationID,
	}

	if d.timeout != nil {
		m["timeout"] = d.timeout.String()
	}

	if d.version != nil {
		m["version"] = *d.version
	}

	for key, fieldName := range headerOptionalStringFields {
		val := reflect.ValueOf(d).Elem().FieldByName(fieldName)
		if !val.IsNil() {
			m[key] = val.Elem().String()
		}
	}

	return json.Marshal(m)
}

func (d *dittoHeaderImpl) UnmarshalJSON(bytes []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return err
	}

	cID, ok := m["correlation-id"]
	if !ok {
		return fmt.Errorf("required field \"correlation-id\" not found")
	}
	var cIDString string
	cIDString, ok = cID.(string)
	if !ok {
		return fmt.Errorf("\"correlation-id\" field not string")
	}
	d.correlationID = cIDString

	var respReqI interface{}
	respReqI, ok = m["response-required"]
	if ok {
		var respReq bool
		respReq, ok = respReqI.(bool)
		if !ok {
			return fmt.Errorf("\"response-required\" field not bool: %s", respReqI)
		}
		d.responseRequired = respReq
	}

	var timeoutI interface{}
	timeoutI, ok = m["timeout"]
	if ok {
		var timeoutStr string
		timeoutStr, ok = timeoutI.(string)
		if !ok {
			return fmt.Errorf("\"timeout\" field not string: %s", timeoutI)
		}

		var timeout time.Duration
		timeout, err = time.ParseDuration(timeoutStr)
		if err != nil {
			return fmt.Errorf("\"timeout\" field is expected contains duration formatted string, but not: %s; %w", timeoutStr, err)
		}
		d.timeout = &timeout
	}

	var verI interface{}
	verI, ok = m["version"]
	if ok {
		var ver int
		ver, ok = verI.(int)
		if !ok {
			return fmt.Errorf("\"version\" field not int: %s", verI)
		}
		d.version = &ver
	}

	for key, fieldName := range headerOptionalStringFields {
		err = d.parseOptionalStringField(m, fieldName, key)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *dittoHeaderImpl) parseOptionalStringField(m map[string]interface{}, fieldName, key string) error {
	valI, ok := m[key]
	if !ok {
		return nil
	}
	var val string
	val, ok = valI.(string)
	if !ok {
		return fmt.Errorf("field \"%s\" expected having typr string, but it's not: %#v", key, valI)
	}
	reflect.ValueOf(d).Elem().FieldByName(fieldName).Set(reflect.ValueOf(&val))
	return nil
}

type DittoResponse interface {
	json.Marshaler
	json.Unmarshaler
	Topic() string
	SetTopic(topic string)
	Path() string
	SetPath(path string)
	Value() interface{}
	SetValue(value interface{})
	Status() int
	SetStatus(status int)
}

func NewDittoResponse() DittoResponse {
	return &dittoResponseImpl{}
}

type dittoResponseImpl struct {
	topic   string
	headers DittoHeader
	path    string
	value   interface{}
	status  int
}

func (d *dittoResponseImpl) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"topic":   d.topic,
		"headers": d.headers,
		"path":    d.path,
		"status":  d.status,
	}

	if d.value != nil {
		m["value"] = d.value
	}

	return json.Marshal(m)
}

func (d *dittoResponseImpl) UnmarshalJSON(bytes []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return err
	}

	// required fields

	{
		var topicI interface{}
		topicI, ok := m["topic"]
		if !ok {
			return fmt.Errorf("required field \"topic\" not found")
		}
		var topic string
		topic, ok = topicI.(string)
		if !ok {
			return fmt.Errorf("\"topic\" field expected string type, but not: %#v", topicI)
		}
		d.topic = topic
	}

	{
		var headersI interface{}
		headersI, ok := m["headers"]
		if !ok {
			return fmt.Errorf("required field \"headers\" not found")
		}
		var headersBytes []byte
		headersBytes, err = json.Marshal(headersI)
		headers := NewHeader("")
		err = json.Unmarshal(headersBytes, headers)
		if err != nil {
			return fmt.Errorf("\"headers\" field expected DittoHeader type, but not: %s; %w", (string)(headersBytes), err)
		}
		d.headers = headers
	}

	{
		var pathI interface{}
		pathI, ok := m["path"]
		if !ok {
			return fmt.Errorf("required field \"path\" not found")
		}
		var path string
		path, ok = pathI.(string)
		if !ok {
			return fmt.Errorf("\"path\" field expected string type, but not: %#v", pathI)
		}
		d.path = path
	}

	// optional fields

	{
		value, ok := m["value"]
		if ok {
			d.value = value
		}
	}

	{
		var statusI interface{}
		statusI, ok := m["status"]
		if !ok {
			return fmt.Errorf("required field \"status\" not found")
		}
		var status int
		status, ok = statusI.(int)
		if !ok {
			return fmt.Errorf("\"status\" field expected int type, but not: %#v", statusI)
		}
		d.status = status
	}

	return nil
}

func (d *dittoResponseImpl) Topic() string {
	return d.topic
}

func (d *dittoResponseImpl) SetTopic(topic string) {
	d.topic = topic
}

func (d *dittoResponseImpl) Path() string {
	return d.path
}

func (d *dittoResponseImpl) SetPath(path string) {
	d.path = path
}

func (d *dittoResponseImpl) Value() interface{} {
	return d.value
}

func (d *dittoResponseImpl) SetValue(value interface{}) {
	d.value = value
}

func (d *dittoResponseImpl) Status() int {
	return d.status
}

func (d *dittoResponseImpl) SetStatus(status int) {
	d.status = status
}
