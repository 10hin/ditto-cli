package interfaces

import "golang.10h.in/ditto/cli/pkg/ditto/lowmodel"

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
