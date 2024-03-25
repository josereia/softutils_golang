package messages

import "errors"

// import "github.com/pkg/errors"

var messageLvl Level
var messageLanguage Language

func SetMessageLvl(lvl Level) {
	messageLvl = lvl
}

func SetMessageLanguage(lang Language) {
	messageLanguage = lang
}

type MessageFunc func(*Message)

type Message struct {
	Description string `json:"description,omitempty"`
	Stack       string `json:"stack,omitempty"`
	Level       Level  `json:"level,omitempty"`
}

func WithCustomDescription(description string) MessageFunc {
	return func(msg *Message) {
		msg.Description = description
	}
}

func WithRootError(err error) MessageFunc {
	return func(msg *Message) {
		// msg.Stack = errors.Cause(err).Error()
		msg.Stack = errors.Unwrap(err).Error()
	}
}

func (msg *Message) Edit(configs ...MessageFunc) {
	for _, fn := range configs {
		fn(msg)
	}
}
