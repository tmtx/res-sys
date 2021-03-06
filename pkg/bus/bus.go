package bus

import (
	"bytes"
	"encoding/gob"

	"github.com/tmtx/res-sys/pkg/validator"
	"go.mongodb.org/mongo-driver/bson"
)

type MessageKey string
type MessageType int
type Status int

type MessageParams bson.M

type Callback func(MessageParams) (validator.Messages, error)

type Message struct {
	Key    MessageKey    `bson:"key"`
	Params MessageParams `bson:"params"`
	Type   MessageType   `bson:"type"`
}

const (
	StatusSuccess Status = iota
	StatusError
)

const (
	CommandMessage MessageType = iota
	EventMessage
)

type ErrorHandler interface {
	Handle(err error, messages validator.Messages)
}

type MessageBus interface {
	Dispatch(m Message)
	DispatchSync(m Message) (validator.Messages, error)
	Subscribe(key MessageKey, cb Callback)
	Listen()
}

func NewCommand(key MessageKey, params MessageParams) Message {
	return Message{
		Key:    key,
		Params: params,
		Type:   CommandMessage,
	}
}

func (m *Message) MarshalBinary() (data []byte, err error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	err = encoder.Encode(m.Key)
	err = encoder.Encode(m.Type)

	gob.Register(m.Params)
	err = encoder.Encode(&m.Params)

	if err != nil {
		return data, err
	}

	return buf.Bytes(), err
}

func (m *Message) UnmarshalBinary(data []byte) (err error) {
	var result Message

	reader := bytes.NewReader(data)
	decoder := gob.NewDecoder(reader)

	err = decoder.Decode(&result.Key)
	err = decoder.Decode(&result.Type)
	err = decoder.Decode(&result.Params)
	if err != nil {
		return err
	}

	m.Key = result.Key
	m.Params = result.Params
	m.Type = result.Type

	return nil
}
