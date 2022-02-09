package config

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Parse(raw map[string]interface{}) (*DittoConfig, error) {
	var err error

	jsonLiteral := &bytes.Buffer{}
	err = json.NewEncoder(jsonLiteral).Encode(raw)
	if err != nil {
		return nil, err
	}

	parsed := &DittoConfig{}
	err = json.NewDecoder(jsonLiteral).Decode(parsed)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

type DittoConfig struct {
	Server *ServerConfig `json:"server"`
}

type ServerConfig struct {
	// TODO: implement
	HTTP      *HTTPConfig      `json:"http,omitempty"`
	Websocket *WebsocketConfig `json:"websocket,omitempty"`
	MQTT      *MQTTConfig      `json:"mqtt,omitempty"`
	AMQP      *AMQPConfig      `json:"amqp,omitempty"`
	Kafka     *KafkaConfig     `json:"kafka,omitempty"`
}

type HTTPConfig struct {
	// TODO: implement
	URLPrefix string           `json:"url_prefix"`
	Basic     *BasicAuthConfig `json:"basic,omitempty"`
	TLS       *TLSConfig       `json:"tls,omitempty"`
}

func (c *HTTPConfig) Client() *http.Client {

}

type WebsocketConfig struct {
	// TODO: implement
	URL   string           `json:"url"`
	Basic *BasicAuthConfig `json:"basic,omitempty"`
	TLS   *TLSConfig       `json:"tls,omitempty"`
}

type MQTTConfig struct {
	// TODO: implement
	Host            string           `json:"host"`
	Port            uint16           `json:"port"`
	Basic           *BasicAuthConfig `json:"basic,omitempty"`
	EnableWebsocket bool             `json:"enable_websocket,omitempty"`
	TLS             *TLSConfig       `json:"tls,omitempty"`
}

type AMQPConfig struct {
	// TODO: implement
	Host  string           `json:"host"`
	Port  uint16           `json:"port"`
	Basic *BasicAuthConfig `json:"basic,omitempty"`
	TLS   *TLSConfig       `json:"tls,omitempty"`
}

type KafkaConfig struct {
	// TODO: implement
	Host string     `json:"host"`
	Port uint16     `json:"port"`
	TLS  *TLSConfig `json:"tls,omitempty"`
}

type BasicAuthConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TLSConfig struct {
	CA     *CAConfig         `json:"ca,omitempty"`
	Client *ClientCertConfig `json:"client,omitempty"`
}

type CAConfig struct {
	CertPath string `json:"cert_path,omitempty"`
	Cert     string `json:"cert,omitempty"`
}

type ClientCertConfig struct {
	CA   *CAConfig       `json:"ca,omitempty"`
	Path *ClientCertPath `json:"path,omitempty"`
	Data *ClientCertData `json:"data,omitempty"`
}

type ClientCertPath struct {
	Cert string `json:"cert,omitempty"`
	Key  string `json:"key,omitempty"`
}

type ClientCertData struct {
	Cert string `json:"cert,omitempty"`
	Key  string `json:"key,omitempty"`
}
